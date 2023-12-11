package cycles

import (
	"github.com/Newt6611/tradevago/tri"
	_ "github.com/Newt6611/tradevago/tri"
)

type LtcUsdt struct { }

func NewLtcUsdt() LtcUsdt {
    return LtcUsdt {}
}

func (this LtcUsdt) GetName() string {
    return "TWD-LTC-USDT"
}

func (this LtcUsdt) GetSymbolsToCheck() []string {
    return []string {
        TWD,
        LTC,
        USDT,
    }
}

func (this LtcUsdt) GetSymbols() []string {
    return []string {
        LTCTWD,
        LTCUSDT,
        USDTTWD,
    }
}

func (this LtcUsdt) GetSides() []tri.Side {
    return []tri.Side {
        tri.BUY,
        tri.SELL,
        tri.SELL,
    }
}


type UsdtLtc struct { }

func NewUsdtLtc() UsdtLtc {
    return UsdtLtc {}
}

func (this UsdtLtc) GetName() string {
    return "TWD-USDT-LTC"
}

func (this UsdtLtc) GetSymbolsToCheck() []string {
    return []string {
        TWD,
        USDT,
        LTC,
    }
}

func (this UsdtLtc) GetSymbols() []string {
    return []string {
        USDTTWD,
        LTCUSDT,
        LTCTWD,
    }
}

func (this UsdtLtc) GetSides() []tri.Side {
    return []tri.Side {
        tri.BUY,
        tri.BUY,
        tri.SELL,
    }
}
