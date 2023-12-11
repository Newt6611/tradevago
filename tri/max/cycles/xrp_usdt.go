package cycles

import (
	"github.com/Newt6611/tradevago/tri"
	_ "github.com/Newt6611/tradevago/tri"
)

type XrpUsdt struct { }

func NewXrpUsdt() XrpUsdt {
    return XrpUsdt {}
}

func (this XrpUsdt) GetName() string {
    return "TWD-XRP-USDT"
}

func (this XrpUsdt) GetSymbolsToCheck() []string {
    return []string {
        TWD,
        XRP,
        USDT,
    }
}

func (this XrpUsdt) GetSymbols() []string {
    return []string {
        XRPTWD,
        XRPUSDT,
        USDTTWD,
    }
}

func (this XrpUsdt) GetSides() []tri.Side {
    return []tri.Side {
        tri.BUY,
        tri.SELL,
        tri.SELL,
    }
}


type UsdtXrp struct { }

func NewUsdtXrp() UsdtXrp {
    return UsdtXrp {}
}

func (this UsdtXrp) GetName() string {
    return "TWD-USDT-XRP"
}

func (this UsdtXrp) GetSymbolsToCheck() []string {
    return []string {
        TWD,
        USDT,
        XRP,
    }
}

func (this UsdtXrp) GetSymbols() []string {
    return []string {
        USDTTWD,
        XRPUSDT,
        XRPTWD,
    }
}

func (this UsdtXrp) GetSides() []tri.Side {
    return []tri.Side {
        tri.BUY,
        tri.BUY,
        tri.SELL,
    }
}
