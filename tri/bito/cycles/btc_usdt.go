package cycles

import "github.com/Newt6611/tradevago/pkg/api"

type BtcUsdt struct { }

func NewBtcUsdt() BtcUsdt {
    return BtcUsdt{}
}

func (this BtcUsdt) GetName() string {
    return "TWD-BTC-USDT"
}

func (this BtcUsdt) GetSymbolsToCheck() []string {
    return []string {
        TWD,
        BTC,
        USDT,
    }
}

func (this BtcUsdt) GetSymbols() []string {
    return []string {
        BTCTWD,
        BTCUSDT,
        USDTTWD,
    }
}

func (this BtcUsdt) GetSides() []api.Side {
    return []api.Side {
        api.BUY,
        api.SELL,
        api.SELL,
    }
}

type UsdtBtc struct { }

func NewUsdtBtc() UsdtBtc {
    return UsdtBtc {}
}

func (this UsdtBtc) GetName() string {
    return "TWD-USDT-BTC"
}

func (this UsdtBtc) GetSymbolsToCheck() []string {
    return []string {
        TWD,
        USDT,
        BTC,
    }
}

func (this UsdtBtc) GetSymbols() []string {
    return []string {
        USDTTWD,
        BTCUSDT,
        BTCTWD,
    }
}

func (this UsdtBtc) GetSides() []api.Side {
    return []api.Side {
        api.BUY,
        api.BUY,
        api.SELL,
    }
}
