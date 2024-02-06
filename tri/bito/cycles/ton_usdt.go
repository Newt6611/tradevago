package cycles

import "github.com/Newt6611/tradevago/pkg/api"

type TonUsdt struct { }

func NewTonUsdt() TonUsdt {
    return TonUsdt{}
}

func (this TonUsdt) GetName() string {
    return "TWD-TON-USDT"
}

func (this TonUsdt) GetSymbolsToCheck() []string {
    return []string {
        TWD,
        TON,
        USDT,
    }
}

func (this TonUsdt) GetSymbols() []string {
    return []string {
        TONTWD,
        TONUSDT,
        USDTTWD,
    }
}

func (this TonUsdt) GetSides() []api.Side {
    return []api.Side {
        api.BUY,
        api.SELL,
        api.SELL,
    }
}

type UsdtTon struct { }

func NewUsdtTon() UsdtTon {
    return UsdtTon {}
}

func (this UsdtTon) GetName() string {
    return "TWD-USDT-TON"
}

func (this UsdtTon) GetSymbolsToCheck() []string {
    return []string {
        TWD,
        USDT,
        TON,
    }
}

func (this UsdtTon) GetSymbols() []string {
    return []string {
        USDTTWD,
        TONUSDT,
        TONTWD,
    }
}

func (this UsdtTon) GetSides() []api.Side {
    return []api.Side {
        api.BUY,
        api.BUY,
        api.SELL,
    }
}
