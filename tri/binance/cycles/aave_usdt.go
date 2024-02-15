package cycles


import (
	"github.com/Newt6611/tradevago/pkg/api"
	_ "github.com/Newt6611/tradevago/tri"
)

type AaveUsdt struct { }

func NewAaveUsdt() AaveUsdt {
    return AaveUsdt {}
}

func (this AaveUsdt) GetName() string {
    return "BTC-AAVE-USDT"
}

func (this AaveUsdt) GetSymbolsToCheck() []string {
    return []string {
        BTC,
        AAVE,
        USDT,
    }
}

func (this AaveUsdt) GetSymbols() []string {
    return []string {
        AAVEBTC,
        AAVEUSDT,
        BTCUSDT,
    }
}

func (this AaveUsdt) GetSides() []api.Side {
    return []api.Side {
        api.BUY,
        api.SELL,
        api.BUY,
    }
}



type UsdtAave struct { }

func NewUsdtAave() UsdtAave {
    return UsdtAave {}
}

func (this UsdtAave) GetName() string {
    return "BTC-USDT-AAVE"
}

func (this UsdtAave) GetSymbolsToCheck() []string {
    return []string {
        BTC,
        USDT,
        AAVE,
    }
}

func (this UsdtAave) GetSymbols() []string {
    return []string {
        BTCUSDT,
        AAVEUSDT,
        AAVEBTC,
    }
}

func (this UsdtAave) GetSides() []api.Side {
    return []api.Side {
        api.SELL,
        api.BUY,
        api.SELL,
    }
}
