package max

import (
	"strings"
	"sync"

	"github.com/Newt6611/tradevago/pkg/api"
	"github.com/Newt6611/tradevago/tri/max/cycles"
)

func getAllCurrencyToCheck() []string {
    return []string {
        cycles.USDT,
        cycles.BTC,
        cycles.ETH,
        cycles.LTC,
        cycles.BCH,
        cycles.XRP,
        cycles.BCNT,
        cycles.USDC,
        cycles.LINK,
        cycles.COMP,
        cycles.DOGE,
        cycles.ADA,
        cycles.DOT,
        cycles.MATIC,
        cycles.SOL,
        cycles.SHIB,
        cycles.SAND,
        cycles.RLY,
        cycles.LOOT,
        cycles.APE,
        cycles.BNB,
        cycles.ETC,
        cycles.ARB,
    }
}

func getTwdQuotePair(currency string) string {
    // currency = MAX_USDT
    // return usdttwd
    c := strings.Split(currency, "_")[1]
    c = strings.ToLower(c)
    return c + "twd"
}

func setupDepthData(data *api.WsDepth, mapStore *sync.Map) {
    switch data.Pair {
    case cycles.USDTTWD:
        (*mapStore).Store(cycles.USDTTWD, *data)

    case cycles.BTCTWD:
        (*mapStore).Store(cycles.BTCTWD, *data)
    case cycles.BTCUSDT:
        (*mapStore).Store(cycles.BTCUSDT, *data)

    case cycles.ETHTWD:
        (*mapStore).Store(cycles.ETHTWD, *data)
    case cycles.ETHUSDT:
        (*mapStore).Store(cycles.ETHUSDT, *data)

    case cycles.LTCTWD:
        (*mapStore).Store(cycles.LTCTWD, *data)
    case cycles.LTCUSDT:
        (*mapStore).Store(cycles.LTCUSDT, *data)

    case cycles.BCHTWD:
        (*mapStore).Store(cycles.BCHTWD, *data)
    case cycles.BCHUSDT:
        (*mapStore).Store(cycles.BCHUSDT, *data)

    case cycles.XRPTWD:
        (*mapStore).Store(cycles.XRPTWD, *data)
    case cycles.XRPUSDT:
        (*mapStore).Store(cycles.XRPUSDT, *data)

    case cycles.BCNTTWD:
        (*mapStore).Store(cycles.BCNTTWD, *data)
    case cycles.BCNTUSDT:
        (*mapStore).Store(cycles.BCNTUSDT, *data)

    case cycles.USDCTWD:
        (*mapStore).Store(cycles.USDCTWD, *data)
    case cycles.USDCUSDT:
        (*mapStore).Store(cycles.USDCUSDT, *data)

    case cycles.LINKTWD:
        (*mapStore).Store(cycles.LINKTWD, *data)
    case cycles.LINKUSDT:
        (*mapStore).Store(cycles.LINKUSDT, *data)

    case cycles.COMPTWD:
        (*mapStore).Store(cycles.COMPTWD, *data)
    case cycles.COMPUSDT:
        (*mapStore).Store(cycles.COMPUSDT, *data)

    case cycles.DOGETWD:
        (*mapStore).Store(cycles.DOGETWD, *data)
    case cycles.DOGEUSDT:
        (*mapStore).Store(cycles.DOGEUSDT, *data)

    case cycles.ADATWD:
        (*mapStore).Store(cycles.ADATWD, *data)
    case cycles.ADAUSDT:
        (*mapStore).Store(cycles.ADAUSDT, *data)

    case cycles.DOTTWD:
        (*mapStore).Store(cycles.DOTTWD, *data)
    case cycles.DOTUSDT:
        (*mapStore).Store(cycles.DOTUSDT, *data)

    case cycles.MATICTWD:
        (*mapStore).Store(cycles.MATICTWD, *data)
    case cycles.MATICUSDT:
        (*mapStore).Store(cycles.MATICUSDT, *data)

    case cycles.SOLTWD:
        (*mapStore).Store(cycles.SOLTWD, *data)
    case cycles.SOLUSDT:
        (*mapStore).Store(cycles.SOLUSDT, *data)

    case cycles.SHIBTWD:
        (*mapStore).Store(cycles.SHIBTWD, *data)
    case cycles.SHIBUSDT:
        (*mapStore).Store(cycles.SHIBUSDT, *data)

    case cycles.SANDTWD:
        (*mapStore).Store(cycles.SANDTWD, *data)
    case cycles.SANDUSDT:
        (*mapStore).Store(cycles.SANDUSDT, *data)

    case cycles.RLYTWD:
        (*mapStore).Store(cycles.RLYTWD, *data)
    case cycles.RLYUSDT:
        (*mapStore).Store(cycles.RLYUSDT, *data)

    case cycles.LOOTTWD:
        (*mapStore).Store(cycles.LOOTTWD, *data)
    case cycles.LOOTUSDT:
        (*mapStore).Store(cycles.LOOTUSDT, *data)

    case cycles.APETWD:
        (*mapStore).Store(cycles.APETWD, *data)
    case cycles.APEUSDT:
        (*mapStore).Store(cycles.APEUSDT, *data)

    case cycles.BNBTWD:
        (*mapStore).Store(cycles.BNBTWD, *data)
    case cycles.BNBUSDT:
        (*mapStore).Store(cycles.BNBUSDT, *data)

    case cycles.ETCTWD:
        (*mapStore).Store(cycles.ETCTWD, *data)
    case cycles.ETCUSDT:
        (*mapStore).Store(cycles.ETCUSDT, *data)

    case cycles.ARBTWD:
        (*mapStore).Store(cycles.ARBTWD, *data)
    case cycles.ARBUSDT:
        (*mapStore).Store(cycles.ARBUSDT, *data)
    }
}

func setUserOrderData(userOrders *[]api.WsUserOrder, mapStore *sync.Map) {
    for _, userOrder := range (*userOrders) {
        mapStore.Swap(userOrder.ID, userOrder)
    }
}

func setBalanceData(balances *[]api.WsUserAccountBalance, mapStore *sync.Map) {
    for _, balance := range (*balances) {
        switch balance.Currency {
        case "twd":
            (*mapStore).Store(cycles.TWD, balance)
        case "max":
            (*mapStore).Store(cycles.MAX, balance)
        case "usdt":
            (*mapStore).Store(cycles.USDT, balance)
        case "btc":
            (*mapStore).Store(cycles.BTC, balance)
        case "eth":
            (*mapStore).Store(cycles.ETH, balance)
        case "ltc":
            (*mapStore).Store(cycles.LTC, balance)
        case "bch":
            (*mapStore).Store(cycles.BCH, balance)
        case "xrp":
            (*mapStore).Store(cycles.XRP, balance)
        case "bcnt":
            (*mapStore).Store(cycles.BCNT, balance)
        case "usdc":
            (*mapStore).Store(cycles.USDC, balance)
        case "link":
            (*mapStore).Store(cycles.LINK, balance)
        case "comp":
            (*mapStore).Store(cycles.COMP, balance)
        case "doge":
            (*mapStore).Store(cycles.DOGE, balance)
        case "ada":
            (*mapStore).Store(cycles.ADA, balance)
        case "dot":
            (*mapStore).Store(cycles.DOT, balance)
        case "matic":
            (*mapStore).Store(cycles.MATIC, balance)
        case "sol":
            (*mapStore).Store(cycles.SOL, balance)
        case "shib":
            (*mapStore).Store(cycles.SHIB, balance)
        case "sand":
            (*mapStore).Store(cycles.SAND, balance)
        case "rly":
            (*mapStore).Store(cycles.RLY, balance)
        case "loot":
            (*mapStore).Store(cycles.LOOT, balance)
        case "ape":
            (*mapStore).Store(cycles.APE, balance)
        case "bnb":
            (*mapStore).Store(cycles.BNB, balance)
        case "etc":
            (*mapStore).Store(cycles.ETC, balance)
        case "arb":
            (*mapStore).Store(cycles.ARB, balance)
        }
    }
}


func convertPairName(name string) string {
    name = strings.ToLower(name)
    names := strings.Split(name, "/")
    var sb = strings.Builder{}
    sb.WriteString(names[0])
    sb.WriteString(names[1])
    return sb.String()
}
