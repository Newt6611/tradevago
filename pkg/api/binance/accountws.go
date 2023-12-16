package binance

import (
	"context"
	"fmt"

	"github.com/Newt6611/tradevago/pkg/api"
)

func (this *BinanceWS) RunAccountConsumer(ctx context.Context) (chan api.WsUserAccountDatas, chan struct{}){
    userAccountChan := make(chan api.WsUserAccountDatas, 500)
    c := make(chan struct{})
    fmt.Println("binance ws account not implment yet ")
    return userAccountChan, c
}
