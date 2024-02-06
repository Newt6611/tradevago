package bito

import (
	"context"
	"encoding/json"

	"github.com/Newt6611/tradevago/internal"
	"github.com/Newt6611/tradevago/pkg/api"
	"github.com/bitoex/bitopro-api-go/pkg/ws"
	"github.com/gorilla/websocket"
)

// History Order is for monitoring "CANCEL" or "DONE"
func (this *BitoWs) RunUserOrderConsumer(ctx context.Context) (chan api.WsUserOrderDatas, chan struct{}) {
    userOrderDataChan := make(chan api.WsUserOrderDatas, 20)
    closeC := make(chan struct{})

    go func() {
        ordersChan, bitoCloseC := this.apiws.RunOrdersWsConsumer(ctx)
        for {
            select {
            case <-closeC:
                close(bitoCloseC)
                goto done
            case orders := <-ordersChan:
                if orders.Err != nil {
                    userOrderDataChan <- api.WsUserOrderDatas{ Err: orders.Err }
                }
                o := mapOrders(&orders)
                userOrderDataChan <- api.WsUserOrderDatas{ Datas: o, Err: nil }
            }
        }
        done:
    }()

    // history orders
    go func() {
        bitoCloseC := this.runOrderHistoryConsumer(ctx, userOrderDataChan)
        for {
            select {
            case <-closeC:
                close(bitoCloseC)
                goto done
            }
        }
        done:
    }()

    return userOrderDataChan, closeC
}

func (this *BitoWs) runOrderHistoryConsumer(ctx context.Context, userOrderDataChan chan api.WsUserOrderDatas) (chan struct{}) {
    path := "ws/v1/pub/auth/orders/histories"
    header, err := newBitoAuthHeader(this.apiws.Email, this.apiws.ApiKey, this.apiws.ApiSecret, path, nil)
    if err != nil {
        userOrderDataChan <- api.WsUserOrderDatas { Err: err }
    }

    close := internal.RunWsClient(ctx, this.apiws.Endpoint + "/" + path, header, nil, func(ws *websocket.Conn, i int, b []byte, err error) {
        if err != nil {
            userOrderDataChan <- api.WsUserOrderDatas { Err: err }
            return
        }

        var historyOrder bitoHistoryEventData
        err = json.Unmarshal(b, &historyOrder)
        if err != nil {
            userOrderDataChan <- api.WsUserOrderDatas{ Err: err }
        } else {
            o := mapHistoryOrders(&historyOrder)
            userOrderDataChan <- api.WsUserOrderDatas{ Datas: o, Err: nil }
        }
    })

    return close
}

func mapOrders(bitoOrders *ws.OrdersData) []api.WsUserOrder {
    orders := []api.WsUserOrder{}

    for pair, _ := range bitoOrders.Data {
        for _, bitoOrder := range bitoOrders.Data[pair] {
            orders = append(orders, api.WsUserOrder {
                ID: bitoOrder.ID,
                Pair: pair,
                Status: getOrderStatus(bitoOrder.Status),
            })
        }
    }
    return orders
}

func mapHistoryOrders(historyOrders *bitoHistoryEventData) []api.WsUserOrder {
    orders := []api.WsUserOrder{}

    for pair, _ := range historyOrders.Data {
        for _, bitoOrder := range historyOrders.Data[pair] {
            status := getOrderStatus(bitoOrder.Status)
            if status != api.OrderStatusCancel && status != api.OrderStatusDone {
                continue
            }
            orders = append(orders, api.WsUserOrder {
                ID: bitoOrder.ID,
                Pair: pair,
                Status: status,
            })
        }
    }
    return orders
}
