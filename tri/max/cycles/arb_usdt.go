package cycles

import (
	"github.com/Newt6611/tradevago/tri"
	_ "github.com/Newt6611/tradevago/tri"
)

type ArbUsdt struct { }

func NewArbUsdt() ArbUsdt {
    return ArbUsdt {}
}

func (this ArbUsdt) GetName() string {
    return "TWD-ARB-USDT"
}

func (this ArbUsdt) GetSymbolsToCheck() []string {
    return []string {
        TWD,
        ARB,
        USDT,
    }
}

func (this ArbUsdt) GetSymbols() []string {
    return []string {
        ARBTWD,
        ARBUSDT,
        USDTTWD,
    }
}

func (this ArbUsdt) GetSides() []tri.Side {
    return []tri.Side {
        tri.BUY,
        tri.SELL,
        tri.SELL,
    }
}

type UsdtArb struct { }

func NewUsdtArb() UsdtArb {
    return UsdtArb {}
}

func (this UsdtArb) GetName() string {
    return "TWD-USDT-ARB"
}

func (this UsdtArb) GetSymbolsToCheck() []string {
    return []string {
        TWD,
        USDT,
        ARB,
    }
}

func (this UsdtArb) GetSymbols() []string {
    return []string {
        USDTTWD,
        ARBUSDT,
        ARBTWD,
    }
}

func (this UsdtArb) GetSides() []tri.Side {
    return []tri.Side {
        tri.BUY,
        tri.BUY,
        tri.SELL,
    }
}
