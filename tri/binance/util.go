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
    for _, balance := range (*balances) {
        switch balance.Currency {
        case "BTC":
            (*mapStore).Store(cycles.BTC, balance)
        case "USDT":
            (*mapStore).Store(cycles.USDT, balance)
        case "ADA":
            (*mapStore).Store(cycles.ADA, balance)
        }
    }
}

func setUserOrderData(userOrders *[]api.WsUserOrder, mapStore *sync.Map) {
    for _, userOrder := range (*userOrders) {
        mapStore.Swap(userOrder.ID, userOrder)
    }
}
func getAllCurrencyToCheck() map[string] api.Side {
    return map[string] api.Side {
        // cycles.BTC,
        cycles.USDT: api.BUY, // BTCUSDT buy back to BTC
        cycles.ADA: api.SELL, // ADABTC sell back to ADA
    }
}

func getBtcQuotePair(currency string) string {
    if currency == cycles.USDT {
        return "BTCUSDT"
    }
    // currency = ADA
    // return ADABTC
    return currency + "BTC"
}

func convertPairName(name string) string {
    return name
}
