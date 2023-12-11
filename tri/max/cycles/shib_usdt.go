package cycles

import (
	"github.com/Newt6611/tradevago/pkg/api"
	_ "github.com/Newt6611/tradevago/tri"
)

type ShibUsdt struct { }

func NewShibUsdt() ShibUsdt {
    return ShibUsdt {}
}

func (this ShibUsdt) GetName() string {
    return "TWD-SHIB-USDT"
}

func (this ShibUsdt) GetSymbolsToCheck() []string {
    return []string {
        TWD,
        SHIB,
        USDT,
    }
}

func (this ShibUsdt) GetSymbols() []string {
    return []string {
        SHIBTWD,
        SHIBUSDT,
        USDTTWD,
    }
}

func (this ShibUsdt) GetSides() []api.Side {
    return []api.Side {
        api.BUY,
        api.SELL,
        api.SELL,
    }
}


type UsdtShib struct { }

func NewUsdtShib() UsdtShib {
    return UsdtShib {}
}

func (this UsdtShib) GetName() string {
    return "TWD-USDT-SHIB"
}

func (this UsdtShib) GetSymbolsToCheck() []string {
    return []string {
        TWD,
        USDT,
        SHIB,
    }
}

func (this UsdtShib) GetSymbols() []string {
    return []string {
        USDTTWD,
        SHIBUSDT,
        SHIBTWD,
    }
}

func (this UsdtShib) GetSides() []api.Side {
    return []api.Side {
        api.BUY,
        api.BUY,
        api.SELL,
    }
}
