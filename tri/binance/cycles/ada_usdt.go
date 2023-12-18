package cycles

import (
	"github.com/Newt6611/tradevago/pkg/api"
	_ "github.com/Newt6611/tradevago/tri"
)

type AdaUsdt struct { }

func NewAdaUsdt() AdaUsdt {
    return AdaUsdt {}
}

func (this AdaUsdt) GetName() string {
    return "BTC-ADA-USDT"
}

func (this AdaUsdt) GetSymbolsToCheck() []string {
    return []string {
        BTC,
        ADA,
        USDT,
    }
}

func (this AdaUsdt) GetSymbols() []string {
    return []string {
        ADABTC,
        ADAUSDT,
        BTCUSDT,
    }
}

func (this AdaUsdt) GetSides() []api.Side {
    return []api.Side {
        api.BUY,
        api.SELL,
        api.BUY,
    }
}



type UsdtAda struct { }

func NewUsdtAda() UsdtAda {
    return UsdtAda {}
}

func (this UsdtAda) GetName() string {
    return "BTC-USDT-ADA"
}

func (this UsdtAda) GetSymbolsToCheck() []string {
    return []string {
        BTC,
        USDT,
        ADA,
    }
}

func (this UsdtAda) GetSymbols() []string {
    return []string {
        BTCUSDT,
        ADAUSDT,
        ADABTC,
    }
}

func (this UsdtAda) GetSides() []api.Side {
    return []api.Side {
        api.SELL,
        api.BUY,
        api.SELL,
    }
}
