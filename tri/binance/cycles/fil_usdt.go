package cycles

import (
	"github.com/Newt6611/tradevago/pkg/api"
	_ "github.com/Newt6611/tradevago/tri"
)

type FilUsdt struct { }

func NewFilUsdt() FilUsdt {
    return FilUsdt {}
}

func (this FilUsdt) GetName() string {
    return "BTC-FIL-USDT"
}

func (this FilUsdt) GetSymbolsToCheck() []string {
    return []string {
        BTC,
        FIL,
        USDT,
    }
}

func (this FilUsdt) GetSymbols() []string {
    return []string {
        FILBTC,
        FILUSDT,
        BTCUSDT,
    }
}

func (this FilUsdt) GetSides() []api.Side {
    return []api.Side {
        api.BUY,
        api.SELL,
        api.BUY,
    }
}



type UsdtFil struct { }

func NewUsdtFil() UsdtFil {
    return UsdtFil {}
}

func (this UsdtFil) GetName() string {
    return "BTC-USDT-FIL"
}

func (this UsdtFil) GetSymbolsToCheck() []string {
    return []string {
        BTC,
        USDT,
        FIL,
    }
}

func (this UsdtFil) GetSymbols() []string {
    return []string {
        BTCUSDT,
        FILUSDT,
        FILBTC,
    }
}

func (this UsdtFil) GetSides() []api.Side {
    return []api.Side {
        api.SELL,
        api.BUY,
        api.SELL,
    }
}
