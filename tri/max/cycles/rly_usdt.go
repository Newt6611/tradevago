package cycles

import (
	"github.com/Newt6611/tradevago/tri"
	_ "github.com/Newt6611/tradevago/tri"
)

type RlyUsdt struct { }

func NewRlyUsdt() RlyUsdt {
    return RlyUsdt {}
}

func (this RlyUsdt) GetName() string {
    return "TWD-RLY-USDT"
}

func (this RlyUsdt) GetSymbolsToCheck() []string {
    return []string {
        TWD,
        RLY,
        USDT,
    }
}

func (this RlyUsdt) GetSymbols() []string {
    return []string {
        RLYTWD,
        RLYUSDT,
        USDTTWD,
    }
}

func (this RlyUsdt) GetSides() []tri.Side {
    return []tri.Side {
        tri.BUY,
        tri.SELL,
        tri.SELL,
    }
}


type UsdtRly struct { }

func NewUsdtRly() UsdtRly {
    return UsdtRly {}
}

func (this UsdtRly) GetName() string {
    return "TWD-USDT-RLY"
}

func (this UsdtRly) GetSymbolsToCheck() []string {
    return []string {
        TWD,
        USDT,
        RLY,
    }
}

func (this UsdtRly) GetSymbols() []string {
    return []string {
        USDTTWD,
        RLYUSDT,
        RLYTWD,
    }
}

func (this UsdtRly) GetSides() []tri.Side {
    return []tri.Side {
        tri.BUY,
        tri.BUY,
        tri.SELL,
    }
}
