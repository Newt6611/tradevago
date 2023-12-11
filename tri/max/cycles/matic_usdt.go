package cycles

import (
	"github.com/Newt6611/tradevago/tri"
	_ "github.com/Newt6611/tradevago/tri"
)

type MaticUsdt struct { }

func NewMaticUsdt() MaticUsdt {
    return MaticUsdt {}
}

func (this MaticUsdt) GetName() string {
    return "TWD-MATIC-USDT"
}

func (this MaticUsdt) GetSymbolsToCheck() []string {
    return []string {
        TWD,
        MATIC,
        USDT,
    }
}

func (this MaticUsdt) GetSymbols() []string {
    return []string {
        MATICTWD,
        MATICUSDT,
        USDTTWD,
    }
}

func (this MaticUsdt) GetSides() []tri.Side {
    return []tri.Side {
        tri.BUY,
        tri.SELL,
        tri.SELL,
    }
}


type UsdtMatic struct { }

func NewUsdtMatic() UsdtMatic {
    return UsdtMatic {}
}

func (this UsdtMatic) GetName() string {
    return "TWD-USDT-MATIC"
}

func (this UsdtMatic) GetSymbolsToCheck() []string {
    return []string {
        TWD,
        USDT,
        MATIC,
    }
}

func (this UsdtMatic) GetSymbols() []string {
    return []string {
        USDTTWD,
        MATICUSDT,
        MATICTWD,
    }
}

func (this UsdtMatic) GetSides() []tri.Side {
    return []tri.Side {
        tri.BUY,
        tri.BUY,
        tri.SELL,
    }
}
