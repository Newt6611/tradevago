package cycles

import (
	"github.com/Newt6611/tradevago/tri"
	_ "github.com/Newt6611/tradevago/tri"
)

type SolUsdt struct { }

func NewSolUsdt() SolUsdt {
    return SolUsdt {}
}

func (this SolUsdt) GetName() string {
    return "TWD-SOL-USDT"
}

func (this SolUsdt) GetSymbolsToCheck() []string {
    return []string {
        TWD,
        SOL,
        USDT,
    }
}

func (this SolUsdt) GetSymbols() []string {
    return []string {
        SOLTWD,
        SOLUSDT,
        USDTTWD,
    }
}

func (this SolUsdt) GetSides() []tri.Side {
    return []tri.Side {
        tri.BUY,
        tri.SELL,
        tri.SELL,
    }
}


type UsdtSol struct { }

func NewUsdtSol() UsdtSol {
    return UsdtSol {}
}

func (this UsdtSol) GetName() string {
    return "TWD-USDT-SOL"
}

func (this UsdtSol) GetSymbolsToCheck() []string {
    return []string {
        TWD,
        USDT,
        SOL,
    }
}

func (this UsdtSol) GetSymbols() []string {
    return []string {
        USDTTWD,
        SOLUSDT,
        SOLTWD,
    }
}

func (this UsdtSol) GetSides() []tri.Side {
    return []tri.Side {
        tri.BUY,
        tri.BUY,
        tri.SELL,
    }
}
