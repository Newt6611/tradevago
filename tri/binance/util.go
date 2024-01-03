package binance

import (
	"sync"

	"github.com/Newt6611/tradevago/pkg/api"
	"github.com/Newt6611/tradevago/tri/binance/cycles"
)

func setupDepthData(data *api.WsDepth, mapStore *sync.Map) {
    (*mapStore).Store(data.Pair, *data)
}

func setBalanceData(balances *[]api.WsUserAccountBalance, mapStore *sync.Map) {
    for _, balance := range (*balances) {
        switch balance.Currency {
        case "BTC":
            (*mapStore).Store(cycles.BTC, balance)
        case "BNB":
            (*mapStore).Store(cycles.BNB, balance)
        case "USDT":
            (*mapStore).Store(cycles.USDT, balance)
        case "ADA":
            (*mapStore).Store(cycles.ADA, balance)
        case "ACA":
            (*mapStore).Store(cycles.ACA, balance)
        case "AAVE":
            (*mapStore).Store(cycles.AAVE, balance)
        case "ACH":
            (*mapStore).Store(cycles.ACH, balance)
        case "OP":
            (*mapStore).Store(cycles.OP, balance)
        case "EOS":
            (*mapStore).Store(cycles.EOS, balance)
        case "FIL":
            (*mapStore).Store(cycles.FIL, balance)
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
        // cycles.BTC, BNB
        cycles.USDT: api.BUY, // BTCUSDT buy back to BTC
        cycles.ADA: api.SELL, // ADABTC sell back to ADA
        cycles.ACA: api.SELL,
        cycles.AAVE: api.SELL,
        cycles.ACH: api.SELL,
        cycles.OP: api.SELL,
        cycles.EOS: api.SELL,
        cycles.FIL: api.SELL,
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
