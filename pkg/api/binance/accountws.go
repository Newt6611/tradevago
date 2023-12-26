package binance

import (
	"context"
	"strconv"

	"github.com/Newt6611/tradevago/pkg/api"
	b "github.com/adshao/go-binance/v2"
)

func (this *BinanceWS) RunAccountConsumer(ctx context.Context) (chan api.WsUserAccountDatas, chan struct{}){
    userAccountChan := make(chan api.WsUserAccountDatas, 500)

    listenKey, err := this.binanceClient.NewStartUserStreamService().Do(ctx)
    if err != nil {
        userAccountChan <- api.WsUserAccountDatas{ Err: err }
    }

    userAccountBalances, err := snapShotAccountBalance(ctx, this.binanceClient)
    userAccountChan <- api.WsUserAccountDatas{ Datas: userAccountBalances, Err: nil }

    _, stopC, err := b.WsUserDataServe(listenKey, userAccountDataHandler(userAccountChan), userAccountDataErrorHandler(userAccountChan))
    if err != nil {
        userAccountChan <- api.WsUserAccountDatas{ Err: err }
    }

    return userAccountChan, stopC
}

func snapShotAccountBalance(ctx context.Context, client *b.Client) ([]api.WsUserAccountBalance, error) {
    datas, err := client.NewGetAccountService().Do(ctx)
    if err != nil {
        return []api.WsUserAccountBalance{}, err
    }

    userAccountBalance := []api.WsUserAccountBalance{}
    for _, balance := range datas.Balances {
        floatBalance, _ := strconv.ParseFloat(balance.Free, 64)
        if floatBalance == 0 {
            continue
        }
        userAccountBalance = append(userAccountBalance, api.WsUserAccountBalance{
            Currency: balance.Asset,
            Balance: floatBalance,
        })
    }

    return userAccountBalance, nil
}

func userAccountDataHandler(userAccountChan chan<- api.WsUserAccountDatas) func(*b.WsUserDataEvent) {
    return func(userDataEvent *b.WsUserDataEvent) {
        userAccountBalance := []api.WsUserAccountBalance{}
        for _, balance := range userDataEvent.AccountUpdate.WsAccountUpdates {
            floatBalance, _ := strconv.ParseFloat(balance.Free, 64)
            userAccountBalance = append(userAccountBalance, api.WsUserAccountBalance{
                Currency: balance.Asset,
                Balance: floatBalance,
            })
        }
        userAccountChan <- api.WsUserAccountDatas { Datas: userAccountBalance, Err: nil }
    }
}

func userAccountDataErrorHandler(userAccountChan chan<- api.WsUserAccountDatas) func(error) {
    return func(err error) {
        userAccountChan <- api.WsUserAccountDatas{ Err: err }
    }
}
