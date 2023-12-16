package binance

import (
	"context"
	"fmt"

	"github.com/Newt6611/tradevago/pkg/api"
)

func (this *BinanceWS) RunUserOrderConsumer(ctx context.Context) (chan api.WsUserOrderDatas, chan struct{}) {
    userOrderDataChan := make(chan api.WsUserOrderDatas, 500)
    c := make(chan struct{})
    fmt.Println("Binance user order not implment yet")
    return userOrderDataChan, c
}
