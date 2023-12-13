package max

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/Newt6611/tradevago/internal"
	"github.com/Newt6611/tradevago/pkg/api"
)

type userOrder struct {
	ID              int     `json:"i"`
	Side            string  `json:"sd"`
	OrderType       string  `json:"ot"`
	Price           string  `json:"p"`
	StopPrice       string  `json:"sp"`
	AveragePrice    string  `json:"ap"`
	Volume          string  `json:"v"`
	RemainVolume    string  `json:"rv"`
	ExecutedVolume  string  `json:"ev"`
	State           string  `json:"S"`
	Market          string  `json:"M"`
	TradeCount      int     `json:"tc"`
	OrderCreatedTime int     `json:"T"`
	OrderUpdatedTime int     `json:"TU"`
	GroupOrderID    int     `json:"gi"`
	ClientOrderID   string  `json:"ci"`
}

// OrderEvent represents the entire order event
type userOrderEvent struct {
	Channel   string `json:"c"`
	Event     string `json:"e"`
	Orders    []userOrder `json:"o"`
}

func (this *MaxWs) RunUserOrderConsumer(ctx context.Context) (chan api.WsUserOrderDatas, chan struct{}) {
    userOrderDataChan := make(chan api.WsUserOrderDatas, 1000)
    nonce := generateNonce()
    subscriptions := authenticationRequest {
        Action: "auth",
        APIKey: this.apiKey,
        Nonce: nonce,
        Signature: generateSignature(this.apiSecret, fmt.Sprintf("%d", nonce)),
        ID: "clientOrderId",
        Filters: []string { "order" },
    }

    close := internal.RunWsClient(ctx, MAX_WS_ENDPOINT, nil, subscriptions, func(t int, b []byte, err error) {
        if err != nil {
            userOrderDataChan <- api.WsUserOrderDatas { Err: err }
            return
        }

        var orderEvent userOrderEvent
        err = json.Unmarshal(b, &orderEvent)
        if err == nil && orderEvent.Event == "order_snapshot" {
            orders := handleUserOrderSnapShot(orderEvent)
            this.userOrderCache = orders
            userOrderDataChan <- api.WsUserOrderDatas { Datas: orders, Err: nil }
            return
        }

        if err == nil && orderEvent.Event == "order_update" {
            handleUserOrderUpdate(&orderEvent, &this.userOrderCache)
            userOrderDataChan <- api.WsUserOrderDatas { Datas: this.userOrderCache, Err: nil }
            return
        }

    })
    return userOrderDataChan, close
}

func handleUserOrderSnapShot(userOrder userOrderEvent) map[string]api.WsUserOrder {
    orders := map[string]api.WsUserOrder{}
    for _, order := range userOrder.Orders {
        o := api.WsUserOrder {
            ID: strconv.Itoa(order.ID),
            Pair: order.Market,
            Status: getOrderStatus(order.State),
        }
        orders[o.ID] = o
    }
    return orders
}

func handleUserOrderUpdate(userOrders *userOrderEvent, userOrderCache *map[string]api.WsUserOrder) {
    for _, order := range userOrders.Orders {
        orderId := strconv.Itoa(order.ID)

        if o, ok := (*userOrderCache)[orderId]; ok {
            o.Status = getOrderStatus(order.State)

            if o.Status == api.OrderStatusCancel {
                delete((*userOrderCache), orderId)
            } else {
                (*userOrderCache)[orderId] = o
            }
        } else {
            o := api.WsUserOrder {
                ID: strconv.Itoa(order.ID),
                Pair: order.Market,
                Status: getOrderStatus(order.State),
            }
            (*userOrderCache)[o.ID] = o
        }
    }
}
