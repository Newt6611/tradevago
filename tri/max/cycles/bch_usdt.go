package cycles

import (
	"github.com/Newt6611/tradevago/tri"
	_ "github.com/Newt6611/tradevago/tri"
)

type BchUsdt struct { }

func NewBchUsdt() BchUsdt {
    return BchUsdt {}
}

func (this BchUsdt) GetName() string {
    return "TWD-BCH-USDT"
}

func (this BchUsdt) GetSymbolsToCheck() []string {
    return []string {
        TWD,
        BCH,
        USDT,
    }
}

func (this BchUsdt) GetSymbols() []string {
    return []string {
        BCHTWD,
        BCHUSDT,
        USDTTWD,
    }
}

func (this BchUsdt) GetSides() []tri.Side {
    return []tri.Side {
        tri.BUY,
        tri.SELL,
        tri.SELL,
    }
}


type UsdtBch struct { }

func NewUsdtBch() UsdtBch {
    return UsdtBch {}
}

func (this UsdtBch) GetName() string {
    return "TWD-USDT-BCH"
}

func (this UsdtBch) GetSymbolsToCheck() []string {
    return []string {
        TWD,
        USDT,
        BCH,
    }
}

func (this UsdtBch) GetSymbols() []string {
    return []string {
        USDTTWD,
        BCHUSDT,
        BCHTWD,
    }
}

func (this UsdtBch) GetSides() []tri.Side {
    return []tri.Side {
        tri.BUY,
        tri.BUY,
        tri.SELL,
    }
}
