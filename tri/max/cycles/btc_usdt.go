package cycles

import (
	"github.com/Newt6611/tradevago/tri"
	_ "github.com/Newt6611/tradevago/tri"
)

type BtcUsdt struct { }

func NewBtcUsdt() BtcUsdt {
    return BtcUsdt {}
}

func (this BtcUsdt) GetName() string {
    return "TWD-BTC-USDT"
}

func (this BtcUsdt) GetSymbolsToCheck() []string {
    return []string {
        TWD,
        BTC,
        USDT,
    }
}

func (this BtcUsdt) GetSymbols() []string {
    return []string {
        BTCTWD,
        BTCUSDT,
        USDTTWD,
    }
}

func (this BtcUsdt) GetSides() []tri.Side {
    return []tri.Side {
        tri.BUY,
        tri.SELL,
        tri.SELL,
    }
}


type UsdtBtc struct { }

func NewUsdtBtc() UsdtBtc {
    return UsdtBtc {}
}

func (this UsdtBtc) GetName() string {
    return "TWD-USDT-BTC"
}

func (this UsdtBtc) GetSymbolsToCheck() []string {
    return []string {
        TWD,
        USDT,
        BTC,
    }
}

func (this UsdtBtc) GetSymbols() []string {
    return []string {
        USDTTWD,
        BTCUSDT,
        BTCTWD,
    }
}

func (this UsdtBtc) GetSides() []tri.Side {
    return []tri.Side {
        tri.BUY,
        tri.BUY,
        tri.SELL,
    }
}
