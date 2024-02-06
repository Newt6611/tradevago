package cycles

import "github.com/Newt6611/tradevago/pkg/api"

type YfiUsdt struct { }

func NewYfiUsdt() YfiUsdt {
    return YfiUsdt{}
}

func (this YfiUsdt) GetName() string {
    return "TWD-YFI-USDT"
}

func (this YfiUsdt) GetSymbolsToCheck() []string {
    return []string {
        TWD,
        YFI,
        USDT,
    }
}

func (this YfiUsdt) GetSymbols() []string {
    return []string {
        YFITWD,
        YFIUSDT,
        USDTTWD,
    }
}

func (this YfiUsdt) GetSides() []api.Side {
    return []api.Side {
        api.BUY,
        api.SELL,
        api.SELL,
    }
}

type UsdtYfi struct { }

func NewUsdtYfi() UsdtYfi {
    return UsdtYfi {}
}

func (this UsdtYfi) GetName() string {
    return "TWD-USDT-YFI"
}

func (this UsdtYfi) GetSymbolsToCheck() []string {
    return []string {
        TWD,
        USDT,
        YFI,
    }
}

func (this UsdtYfi) GetSymbols() []string {
    return []string {
        USDTTWD,
        YFIUSDT,
        YFITWD,
    }
}

func (this UsdtYfi) GetSides() []api.Side {
    return []api.Side {
        api.BUY,
        api.BUY,
        api.SELL,
    }
}
