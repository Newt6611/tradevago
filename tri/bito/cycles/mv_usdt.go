package cycles

import "github.com/Newt6611/tradevago/pkg/api"

type MvUsdt struct { }

func NewMvUsdt() MvUsdt {
    return MvUsdt{}
}

func (this MvUsdt) GetName() string {
    return "TWD-MV-USDT"
}

func (this MvUsdt) GetSymbolsToCheck() []string {
    return []string {
        TWD,
        MV,
        USDT,
    }
}

func (this MvUsdt) GetSymbols() []string {
    return []string {
        MVTWD,
        MVUSDT,
        USDTTWD,
    }
}

func (this MvUsdt) GetSides() []api.Side {
    return []api.Side {
        api.BUY,
        api.SELL,
        api.SELL,
    }
}

type UsdtMv struct { }

func NewUsdtMv() UsdtMv {
    return UsdtMv {}
}

func (this UsdtMv) GetName() string {
    return "TWD-USDT-MV"
}

func (this UsdtMv) GetSymbolsToCheck() []string {
    return []string {
        TWD,
        USDT,
        MV,
    }
}

func (this UsdtMv) GetSymbols() []string {
    return []string {
        USDTTWD,
        MVUSDT,
        MVTWD,
    }
}

func (this UsdtMv) GetSides() []api.Side {
    return []api.Side {
        api.BUY,
        api.BUY,
        api.SELL,
    }
}
