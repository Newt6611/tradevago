package cycles

import (
	"github.com/Newt6611/tradevago/tri"
	_ "github.com/Newt6611/tradevago/tri"
)

type DotUsdt struct { }

func NewDotUsdt() DotUsdt {
    return DotUsdt {}
}

func (this DotUsdt) GetName() string {
    return "TWD-DOT-USDT"
}

func (this DotUsdt) GetSymbolsToCheck() []string {
    return []string {
        TWD,
        DOT,
        USDT,
    }
}

func (this DotUsdt) GetSymbols() []string {
    return []string {
        DOTTWD,
        DOTUSDT,
        USDTTWD,
    }
}

func (this DotUsdt) GetSides() []tri.Side {
    return []tri.Side {
        tri.BUY,
        tri.SELL,
        tri.SELL,
    }
}


type UsdtDot struct { }

func NewUsdtDot() UsdtDot {
    return UsdtDot {}
}

func (this UsdtDot) GetName() string {
    return "TWD-USDT-DOT"
}

func (this UsdtDot) GetSymbolsToCheck() []string {
    return []string {
        TWD,
        USDT,
        DOT,
    }
}

func (this UsdtDot) GetSymbols() []string {
    return []string {
        USDTTWD,
        DOTUSDT,
        DOTTWD,
    }
}

func (this UsdtDot) GetSides() []tri.Side {
    return []tri.Side {
        tri.BUY,
        tri.BUY,
        tri.SELL,
    }
}
