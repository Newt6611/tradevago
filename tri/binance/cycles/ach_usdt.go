package cycles

import (
	"github.com/Newt6611/tradevago/pkg/api"
	_ "github.com/Newt6611/tradevago/tri"
)

type AchUsdt struct { }

func NewAchUsdt() AchUsdt {
    return AchUsdt {}
}

func (this AchUsdt) GetName() string {
    return "BTC-ACH-USDT"
}

func (this AchUsdt) GetSymbolsToCheck() []string {
    return []string {
        BTC,
        ACH,
        USDT,
    }
}

func (this AchUsdt) GetSymbols() []string {
    return []string {
        ACHBTC,
        ACHUSDT,
        BTCUSDT,
    }
}

func (this AchUsdt) GetSides() []api.Side {
    return []api.Side {
        api.BUY,
        api.SELL,
        api.BUY,
    }
}



type UsdtAch struct { }

func NewUsdtAch() UsdtAch {
    return UsdtAch {}
}

func (this UsdtAch) GetName() string {
    return "BTC-USDT-ACH"
}

func (this UsdtAch) GetSymbolsToCheck() []string {
    return []string {
        BTC,
        USDT,
        ACH,
    }
}

func (this UsdtAch) GetSymbols() []string {
    return []string {
        BTCUSDT,
        ACHUSDT,
        ACHBTC,
    }
}

func (this UsdtAch) GetSides() []api.Side {
    return []api.Side {
        api.SELL,
        api.BUY,
        api.SELL,
    }
}
