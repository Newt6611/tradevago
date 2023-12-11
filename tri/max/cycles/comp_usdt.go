package cycles

import (
	"github.com/Newt6611/tradevago/tri"
	_ "github.com/Newt6611/tradevago/tri"
)

type CompUsdt struct { }

func NewCompUsdt() CompUsdt {
    return CompUsdt {}
}

func (this CompUsdt) GetName() string {
    return "TWD-COMP-USDT"
}

func (this CompUsdt) GetSymbolsToCheck() []string {
    return []string {
        TWD,
        COMP,
        USDT,
    }
}

func (this CompUsdt) GetSymbols() []string {
    return []string {
        COMPTWD,
        COMPUSDT,
        USDTTWD,
    }
}

func (this CompUsdt) GetSides() []tri.Side {
    return []tri.Side {
        tri.BUY,
        tri.SELL,
        tri.SELL,
    }
}


type UsdtComp struct { }

func NewUsdtComp() UsdtComp {
    return UsdtComp {}
}

func (this UsdtComp) GetName() string {
    return "TWD-USDT-COMP"
}

func (this UsdtComp) GetSymbolsToCheck() []string {
    return []string {
        TWD,
        USDT,
        COMP,
    }
}

func (this UsdtComp) GetSymbols() []string {
    return []string {
        USDTTWD,
        COMPUSDT,
        COMPTWD,
    }
}

func (this UsdtComp) GetSides() []tri.Side {
    return []tri.Side {
        tri.BUY,
        tri.BUY,
        tri.SELL,
    }
}
