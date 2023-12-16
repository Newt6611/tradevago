package max

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"github.com/Newt6611/tradevago/internal"
	"github.com/Newt6611/tradevago/pkg/api"
	"github.com/gorilla/websocket"
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
    userOrderDataChan := make(chan api.WsUserOrderDatas, 500)
    nonce := generateNonce()
    subscriptions := authenticationRequest {
        Action: "auth",
        APIKey: this.apiKey,
        Nonce: nonce,
        Signature: generateSignature(this.apiSecret, fmt.Sprintf("%d", nonce)),
        ID: "clientOrderId",
        Filters: []string { "order" },
    }

    close := internal.RunWsClient(ctx, MAX_WS_ENDPOINT, nil, subscriptions, func(ws *websocket.Conn, t int, b []byte, err error) {
        if err != nil {
            userOrderDataChan <- api.WsUserOrderDatas { Err: err }
            return
        }

        var errEvent errorEvent
        err = json.Unmarshal(b, &errEvent)
        if err == nil && errEvent.Event == "error"{
            userOrderDataChan <- api.WsUserOrderDatas { Err: errors.New(errEvent.Errors[0]) }
            return
        }

        var orderEvent userOrderEvent
        err = json.Unmarshal(b, &orderEvent)
        if err == nil {
            orders := handleUserOrderSnapShot(orderEvent)
            userOrderDataChan <- api.WsUserOrderDatas { Datas: orders, Err: nil }
            return
        }
    })
    return userOrderDataChan, close
}

func handleUserOrderSnapShot(userOrder userOrderEvent) []api.WsUserOrder {
    orders := []api.WsUserOrder{}
    for _, order := range userOrder.Orders {
        orders = append(orders, api.WsUserOrder {
            ID: strconv.Itoa(order.ID),
            Pair: order.Market,
            Status: getOrderStatus(order.State),
        })
    }
    return orders
}
