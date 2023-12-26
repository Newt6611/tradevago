package binance

import (
	"context"
	"strconv"

	"github.com/Newt6611/tradevago/pkg/api"
	b "github.com/adshao/go-binance/v2"
)

func (this *BinanceWS) RunUserOrderConsumer(ctx context.Context) (chan api.WsUserOrderDatas, chan struct{}) {
    userOrderDataChan := make(chan api.WsUserOrderDatas, 500)

    listenKey, err := this.binanceClient.NewStartUserStreamService().Do(ctx)
    if err != nil {
        userOrderDataChan <- api.WsUserOrderDatas{ Err: err }
    }

    _, stopC, err := b.WsUserDataServe(listenKey, userOrderDataHandler(userOrderDataChan), userOrderDataErrorHandler(userOrderDataChan))
    if err != nil {
        userOrderDataChan <- api.WsUserOrderDatas{ Err: err }
    }

    return userOrderDataChan, stopC
}

func userOrderDataHandler(userOrderDataChan chan<- api.WsUserOrderDatas) func(*b.WsUserDataEvent) {
    return func(userDataEvent *b.WsUserDataEvent) {
        orders := []api.WsUserOrder {
            {
                ID: strconv.FormatInt(userDataEvent.OrderUpdate.Id, 10),
                Pair: userDataEvent.OrderUpdate.Symbol,
                Status: getOrderStatus(userDataEvent.OrderUpdate.Status),
            },
        }
        userOrderDataChan <- api.WsUserOrderDatas { Datas: orders, Err: nil }
    }
}

func userOrderDataErrorHandler(userOrderDataChan chan<- api.WsUserOrderDatas) func(error) {
    return func(err error) {
        userOrderDataChan <- api.WsUserOrderDatas{ Err: err }
    }
}
