package cycles

import (
	"github.com/Newt6611/tradevago/tri"
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

func (this AdaUsdt) GetSides() []tri.Side {
    return []tri.Side {
        tri.BUY,
        tri.SELL,
        tri.SELL,
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

func (this UsdtAda) GetSides() []tri.Side {
    return []tri.Side {
        tri.BUY,
        tri.BUY,
        tri.SELL,
    }
}
