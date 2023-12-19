package binance

import (
	"context"
	"fmt"

	"github.com/Newt6611/tradevago/pkg/api"
	b "github.com/adshao/go-binance/v2"
)

func (this *BinanceWS) RunUserOrderConsumer(ctx context.Context) (chan api.WsUserOrderDatas, chan struct{}) {
    userOrderDataChan := make(chan api.WsUserOrderDatas, 500)

    listenKey, err := this.binanceClient.NewStartUserStreamService().Do(ctx)
    userOrderDataChan <- api.WsUserOrderDatas{ Err: err }

    _, stopC, err := b.WsUserDataServe(listenKey, userDataHandler(userOrderDataChan), userDataErrorHandler(userOrderDataChan))
    if err != nil {
        userOrderDataChan <- api.WsUserOrderDatas{ Err: err }
    }

    return userOrderDataChan, stopC
}

func userDataHandler(userOrderDataChan chan<- api.WsUserOrderDatas) func(*b.WsUserDataEvent) {
    return func(userDataEvent *b.WsUserDataEvent) {
        fmt.Println(userDataEvent.OrderUpdate)
        // userOrderDatas := api.WsUserOrderDatas {}

        // for _, userData := range userDataEvent.OrderUpdate {
        //     userDataEvent.OrderUpdate
        // }
    }
}

func userDataErrorHandler(userOrderDataChan chan<- api.WsUserOrderDatas) func(error) {
    return func(err error) {
        userOrderDataChan <- api.WsUserOrderDatas{ Err: err }
    }
}
