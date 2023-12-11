package cycles

import (
	"github.com/Newt6611/tradevago/tri"
	_ "github.com/Newt6611/tradevago/tri"
)

type DogeUsdt struct { }

func NewDogeUsdt() DogeUsdt {
    return DogeUsdt {}
}

func (this DogeUsdt) GetName() string {
    return "TWD-DOGE-USDT"
}

func (this DogeUsdt) GetSymbolsToCheck() []string {
    return []string {
        TWD,
        DOGE,
        USDT,
    }
}

func (this DogeUsdt) GetSymbols() []string {
    return []string {
        DOGETWD,
        DOGEUSDT,
        USDTTWD,
    }
}

func (this DogeUsdt) GetSides() []tri.Side {
    return []tri.Side {
        tri.BUY,
        tri.SELL,
        tri.SELL,
    }
}


type UsdtDoge struct { }

func NewUsdtDoge() UsdtDoge {
    return UsdtDoge {}
}

func (this UsdtDoge) GetName() string {
    return "TWD-USDT-DOGE"
}

func (this UsdtDoge) GetSymbolsToCheck() []string {
    return []string {
        TWD,
        USDT,
        DOGE,
    }
}

func (this UsdtDoge) GetSymbols() []string {
    return []string {
        USDTTWD,
        DOGEUSDT,
        DOGETWD,
    }
}

func (this UsdtDoge) GetSides() []tri.Side {
    return []tri.Side {
        tri.BUY,
        tri.BUY,
        tri.SELL,
    }
}
