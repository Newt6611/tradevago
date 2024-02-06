package bito

import (
	"context"
	"strconv"

	"github.com/Newt6611/tradevago/pkg/api"
	"github.com/bitoex/bitopro-api-go/pkg/ws"
)


func (this *BitoWs) RunAccountConsumer(ctx context.Context) (chan api.WsUserAccountDatas, chan struct{}){
    userAccountChan := make(chan api.WsUserAccountDatas, 50)
    closeC := make(chan struct{})

    go func() {
        dataChan, bitoCloseC := this.apiws.RunAccountBalancesWsConsumer(ctx)
        for {
            select {
            case <-closeC:
                close(bitoCloseC)
                goto done
            case datas := <-dataChan:
                if datas.Err != nil {
                    userAccountChan <- api.WsUserAccountDatas{ Err: datas.Err }
                }
                b := mapUserAccountData(datas)
                userAccountChan <- api.WsUserAccountDatas{ Datas: b, Err: nil }
            }
        }
        done:
    }()

    return userAccountChan, closeC
}

func mapUserAccountData(bitoBalances ws.AccountBalanceData) []api.WsUserAccountBalance {
    userBalances := []api.WsUserAccountBalance{}
    for _, account := range bitoBalances.Data {
        balances, _ := strconv.ParseFloat(account.Amount, 64)
        userBalances = append(userBalances, api.WsUserAccountBalance{
            Currency: account.Currency,
            Balance: balances,
        })
        // if (account.Currency == "USDT" || account.Currency == "TWD") {
        //     fmt.Printf("%s %f\n", account.Currency, balances)
        // }
    }
    return userBalances
}
