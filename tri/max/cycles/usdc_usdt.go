package cycles

import (
	"github.com/Newt6611/tradevago/tri"
	_ "github.com/Newt6611/tradevago/tri"
)

type UsdcUsdt struct { }

func NewUsdcUsdt() UsdcUsdt {
    return UsdcUsdt {}
}

func (this UsdcUsdt) GetName() string {
    return "TWD-USDC-USDT"
}

func (this UsdcUsdt) GetSymbolsToCheck() []string {
    return []string {
        TWD,
        USDC,
        USDT,
    }
}

func (this UsdcUsdt) GetSymbols() []string {
    return []string {
        USDCTWD,
        USDCUSDT,
        USDTTWD,
    }
}

func (this UsdcUsdt) GetSides() []tri.Side {
    return []tri.Side {
        tri.BUY,
        tri.SELL,
        tri.SELL,
    }
}


type UsdtUsdc struct { }

func NewUsdtUsdc() UsdtUsdc {
    return UsdtUsdc {}
}

func (this UsdtUsdc) GetName() string {
    return "TWD-USDT-USDC"
}

func (this UsdtUsdc) GetSymbolsToCheck() []string {
    return []string {
        TWD,
        USDT,
        USDC,
    }
}

func (this UsdtUsdc) GetSymbols() []string {
    return []string {
        USDTTWD,
        USDCUSDT,
        USDCTWD,
    }
}

func (this UsdtUsdc) GetSides() []tri.Side {
    return []tri.Side {
        tri.BUY,
        tri.BUY,
        tri.SELL,
    }
}
