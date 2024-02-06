package bito

import (
	"strings"
	"sync"

	"github.com/Newt6611/tradevago/pkg/api"
	"github.com/Newt6611/tradevago/tri/bito/cycles"
)

func getAllCurrencyToCheck() map[string] api.Side {
    return map[string] api.Side {
        cycles.USDT: api.SELL,
        cycles.BTC: api.SELL,
        cycles.ETH: api.SELL,
        cycles.SOL: api.SELL,
        cycles.DOGE: api.SELL,
        cycles.USDC: api.SELL,
        cycles.ADA: api.SELL,
        cycles.XRP: api.SELL,
        cycles.APE: api.SELL,
        cycles.MATIC: api.SELL,
        cycles.TRX: api.SELL,
        cycles.MV: api.SELL,
        cycles.LTC: api.SELL,
        cycles.EOS: api.SELL,
        cycles.SHIB: api.SELL,
        cycles.BNB: api.SELL,
        cycles.BCH: api.SELL,
        cycles.TON: api.SELL,
        cycles.YFI: api.SELL,
    }
}

func getTwdQuotePair(currency string) string {
    // currency = USDT
    // return usdt_twd
    c := strings.ToLower(currency)
    return c + "_twd"
}

func setUserOrderData(userOrders *[]api.WsUserOrder, mapStore *sync.Map) {
    for _, userOrder := range (*userOrders) {
        // In bitopro, sometimes "StatusDone" might come faster than "StatusWait"
        o, ok := mapStore.Load(userOrder.ID)
        preOrder, convert_ok := o.(api.WsUserOrder)
        if ok && convert_ok && preOrder.Status == api.OrderStatusDone {
            continue
        }
        mapStore.Swap(userOrder.ID, userOrder)
    }
}

func setupDepthData(data *api.WsDepth, mapStore *sync.Map) {
    // bito ws: BTC_TWD, normal pair: btc_twd
    data.Pair = strings.ToLower(data.Pair)
    (*mapStore).Store(data.Pair, *data)
}

func setBalanceData(balances *[]api.WsUserAccountBalance, mapStore *sync.Map) {
    for _, balance := range (*balances) {
        (*mapStore).Store(balance.Currency, balance)
    }
}

func convertPairName(name string) string {
    return name
}
