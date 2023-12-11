package cycles

import (
	"github.com/Newt6611/tradevago/pkg/api"
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

func (this BchUsdt) GetSides() []api.Side {
    return []api.Side {
        api.BUY,
        api.SELL,
        api.SELL,
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

func (this UsdtBch) GetSides() []api.Side {
    return []api.Side {
        api.BUY,
        api.BUY,
        api.SELL,
    }
}
