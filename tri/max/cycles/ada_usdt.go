package cycles

import (
	"github.com/Newt6611/tradevago/pkg/api"
	_ "github.com/Newt6611/tradevago/tri"
)

type AdaUsdt struct { }

func NewAdaUsdt() AdaUsdt {
    return AdaUsdt {}
}

func (this AdaUsdt) GetName() string {
    return "TWD-ADA-USDT"
}

func (this AdaUsdt) GetSymbolsToCheck() []string {
    return []string {
        TWD,
        ADA,
        USDT,
    }
}

func (this AdaUsdt) GetSymbols() []string {
    return []string {
        ADATWD,
        ADAUSDT,
        USDTTWD,
    }
}

func (this AdaUsdt) GetSides() []api.Side {
    return []api.Side {
        api.BUY,
        api.SELL,
        api.SELL,
    }
}

type UsdtAda struct { }

func NewUsdtAda() UsdtAda {
    return UsdtAda {}
}

func (this UsdtAda) GetName() string {
    return "TWD-USDT-ADA"
}

func (this UsdtAda) GetSymbolsToCheck() []string {
    return []string {
        TWD,
        USDT,
        ADA,
    }
}

func (this UsdtAda) GetSymbols() []string {
    return []string {
        USDTTWD,
        ADAUSDT,
        ADATWD,
    }
}

func (this UsdtAda) GetSides() []api.Side {
    return []api.Side {
        api.BUY,
        api.BUY,
        api.SELL,
    }
}
