package cycles

import (
	"github.com/Newt6611/tradevago/pkg/api"
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

func (this BnbUsdt) GetSides() []api.Side {
    return []api.Side {
        api.BUY,
        api.SELL,
        api.SELL,
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

func (this UsdtBnb) GetSides() []api.Side {
    return []api.Side {
        api.BUY,
        api.BUY,
        api.SELL,
    }
}
