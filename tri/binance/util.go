package binance

import (
	"sync"

	"github.com/Newt6611/tradevago/pkg/api"
	"github.com/Newt6611/tradevago/tri/binance/cycles"
)

func setupDepthData(data *api.WsDepth, mapStore *sync.Map) {
    switch data.Pair {
    case cycles.BTCUSDT:
        (*mapStore).Store(cycles.BTCUSDT, *data)

    case cycles.ADABTC:
        (*mapStore).Store(cycles.ADABTC, *data)
    case cycles.ADAUSDT:
        (*mapStore).Store(cycles.ADAUSDT, *data)
    }
}

func setBalanceData(balances *[]api.WsUserAccountBalance, mapStore *sync.Map) {
    // switch balance.Currency {
        // for _, _ := range (*balances) {
        //
        // }
    // }
}

func setUserOrderData(userOrders *[]api.WsUserOrder, mapStore *sync.Map) {
    for _, userOrder := range (*userOrders) {
        mapStore.Swap(userOrder.ID, userOrder)
    }
}

func convertPairName(name string) string {
    return ""
}
