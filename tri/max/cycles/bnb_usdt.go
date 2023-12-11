package cycles

import (
	"github.com/Newt6611/tradevago/tri"
	_ "github.com/Newt6611/tradevago/tri"
)

type BnbUsdt struct { }

func NewBnbUsdt() BnbUsdt {
    return BnbUsdt {}
}

func (this BnbUsdt) GetName() string {
    return "TWD-BNB-USDT"
}

func (this BnbUsdt) GetSymbolsToCheck() []string {
    return []string {
        TWD,
        BNB,
        USDT,
    }
}

func (this BnbUsdt) GetSymbols() []string {
    return []string {
        BNBTWD,
        BNBUSDT,
        USDTTWD,
    }
}

func (this BnbUsdt) GetSides() []tri.Side {
    return []tri.Side {
        tri.BUY,
        tri.SELL,
        tri.SELL,
    }
}


type UsdtBnb struct { }

func NewUsdtBnb() UsdtBnb {
    return UsdtBnb {}
}

func (this UsdtBnb) GetName() string {
    return "TWD-USDT-BNB"
}

func (this UsdtBnb) GetSymbolsToCheck() []string {
    return []string {
        TWD,
        USDT,
        BNB,
    }
}

func (this UsdtBnb) GetSymbols() []string {
    return []string {
        USDTTWD,
        BNBUSDT,
        BNBTWD,
    }
}

func (this UsdtBnb) GetSides() []tri.Side {
    return []tri.Side {
        tri.BUY,
        tri.BUY,
        tri.SELL,
    }
}
