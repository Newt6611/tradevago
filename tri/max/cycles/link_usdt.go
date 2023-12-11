package cycles

import (
	"github.com/Newt6611/tradevago/tri"
	_ "github.com/Newt6611/tradevago/tri"
)

type LinkUsdt struct { }

func NewLinkUsdt() LinkUsdt {
    return LinkUsdt {}
}

func (this LinkUsdt) GetName() string {
    return "TWD-LINK-USDT"
}

func (this LinkUsdt) GetSymbolsToCheck() []string {
    return []string {
        TWD,
        LINK,
        USDT,
    }
}

func (this LinkUsdt) GetSymbols() []string {
    return []string {
        LINKTWD,
        LINKUSDT,
        USDTTWD,
    }
}

func (this LinkUsdt) GetSides() []tri.Side {
    return []tri.Side {
        tri.BUY,
        tri.SELL,
        tri.SELL,
    }
}


type UsdtLink struct { }

func NewUsdtLink() UsdtLink {
    return UsdtLink {}
}

func (this UsdtLink) GetName() string {
    return "TWD-USDT-LINK"
}

func (this UsdtLink) GetSymbolsToCheck() []string {
    return []string {
        TWD,
        USDT,
        LINK,
    }
}

func (this UsdtLink) GetSymbols() []string {
    return []string {
        USDTTWD,
        LINKUSDT,
        LINKTWD,
    }
}

func (this UsdtLink) GetSides() []tri.Side {
    return []tri.Side {
        tri.BUY,
        tri.BUY,
        tri.SELL,
    }
}
