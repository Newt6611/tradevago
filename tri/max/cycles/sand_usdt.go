package cycles

import (
	"github.com/Newt6611/tradevago/tri"
	_ "github.com/Newt6611/tradevago/tri"
)

type SandUsdt struct { }

func NewSandUsdt() SandUsdt {
    return SandUsdt {}
}

func (this SandUsdt) GetName() string {
    return "TWD-SAND-USDT"
}

func (this SandUsdt) GetSymbolsToCheck() []string {
    return []string {
        TWD,
        SAND,
        USDT,
    }
}

func (this SandUsdt) GetSymbols() []string {
    return []string {
        SANDTWD,
        SANDUSDT,
        USDTTWD,
    }
}

func (this SandUsdt) GetSides() []tri.Side {
    return []tri.Side {
        tri.BUY,
        tri.SELL,
        tri.SELL,
    }
}


type UsdtSand struct { }

func NewUsdtSand() UsdtSand {
    return UsdtSand {}
}

func (this UsdtSand) GetName() string {
    return "TWD-USDT-SAND"
}

func (this UsdtSand) GetSymbolsToCheck() []string {
    return []string {
        TWD,
        USDT,
        SAND,
    }
}

func (this UsdtSand) GetSymbols() []string {
    return []string {
        USDTTWD,
        SANDUSDT,
        SANDTWD,
    }
}

func (this UsdtSand) GetSides() []tri.Side {
    return []tri.Side {
        tri.BUY,
        tri.BUY,
        tri.SELL,
    }
}
