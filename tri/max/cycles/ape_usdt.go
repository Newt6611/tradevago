package cycles

import (
	"github.com/Newt6611/tradevago/pkg/api"
	_ "github.com/Newt6611/tradevago/tri"
)

type ApeUsdt struct { }

func NewApeUsdt() ApeUsdt {
    return ApeUsdt {}
}

func (this ApeUsdt) GetName() string {
    return "TWD-APE-USDT"
}

func (this ApeUsdt) GetSymbolsToCheck() []string {
    return []string {
        TWD,
        APE,
        USDT,
    }
}

func (this ApeUsdt) GetSymbols() []string {
    return []string {
        APETWD,
        APEUSDT,
        USDTTWD,
    }
}

func (this ApeUsdt) GetSides() []api.Side {
    return []api.Side {
        api.BUY,
        api.SELL,
        api.SELL,
    }
}

type UsdtApe struct { }

func NewUsdtApe() UsdtApe {
    return UsdtApe {}
}

func (this UsdtApe) GetName() string {
    return "TWD-USDT-APE"
}

func (this UsdtApe) GetSymbolsToCheck() []string {
    return []string {
        TWD,
        USDT,
        APE,
    }
}

func (this UsdtApe) GetSymbols() []string {
    return []string {
        USDTTWD,
        APEUSDT,
        APETWD,
    }
}

func (this UsdtApe) GetSides() []api.Side {
    return []api.Side {
        api.BUY,
        api.BUY,
        api.SELL,
    }
}
