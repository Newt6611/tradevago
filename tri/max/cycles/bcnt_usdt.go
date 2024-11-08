package cycles

import (
	"github.com/Newt6611/tradevago/pkg/api"
	_ "github.com/Newt6611/tradevago/tri"
)

type BcntUsdt struct { }

func NewBcntUsdt() BcntUsdt {
    return BcntUsdt {}
}

func (this BcntUsdt) GetName() string {
    return "TWD-BCNT-USDT"
}

func (this BcntUsdt) GetSymbolsToCheck() []string {
    return []string {
        TWD,
        BCNT,
        USDT,
    }
}

func (this BcntUsdt) GetSymbols() []string {
    return []string {
        BCNTTWD,
        BCNTUSDT,
        USDTTWD,
    }
}

func (this BcntUsdt) GetSides() []api.Side {
    return []api.Side {
        api.BUY,
        api.SELL,
        api.SELL,
    }
}


type UsdtBcnt struct { }

func NewUsdtBcnt() UsdtBcnt {
    return UsdtBcnt {}
}

func (this UsdtBcnt) GetName() string {
    return "TWD-USDT-BCNT"
}

func (this UsdtBcnt) GetSymbolsToCheck() []string {
    return []string {
        TWD,
        USDT,
        BCNT,
    }
}

func (this UsdtBcnt) GetSymbols() []string {
    return []string {
        USDTTWD,
        BCNTUSDT,
        BCNTTWD,
    }
}

func (this UsdtBcnt) GetSides() []api.Side {
    return []api.Side {
        api.BUY,
        api.BUY,
        api.SELL,
    }
}
