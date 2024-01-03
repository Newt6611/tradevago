package cycles

import (
	"github.com/Newt6611/tradevago/pkg/api"
	_ "github.com/Newt6611/tradevago/tri"
)

type AcaUsdt struct { }

func NewAcaUsdt() AcaUsdt {
    return AcaUsdt {}
}

func (this AcaUsdt) GetName() string {
    return "BTC-ACA-USDT"
}

func (this AcaUsdt) GetSymbolsToCheck() []string {
    return []string {
        BTC,
        ACA,
        USDT,
    }
}

func (this AcaUsdt) GetSymbols() []string {
    return []string {
        ACABTC,
        ACAUSDT,
        BTCUSDT,
    }
}

func (this AcaUsdt) GetSides() []api.Side {
    return []api.Side {
        api.BUY,
        api.SELL,
        api.BUY,
    }
}



type UsdtAca struct { }

func NewUsdtAca() UsdtAca {
    return UsdtAca {}
}

func (this UsdtAca) GetName() string {
    return "BTC-USDT-ACA"
}

func (this UsdtAca) GetSymbolsToCheck() []string {
    return []string {
        BTC,
        USDT,
        ACA,
    }
}

func (this UsdtAca) GetSymbols() []string {
    return []string {
        BTCUSDT,
        ACAUSDT,
        ACABTC,
    }
}

func (this UsdtAca) GetSides() []api.Side {
    return []api.Side {
        api.SELL,
        api.BUY,
        api.SELL,
    }
}
