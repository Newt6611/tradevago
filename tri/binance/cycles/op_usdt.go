package cycles

import (
	"github.com/Newt6611/tradevago/pkg/api"
	_ "github.com/Newt6611/tradevago/tri"
)

type OpUsdt struct { }

func NewOpUsdt() OpUsdt {
    return OpUsdt {}
}

func (this OpUsdt) GetName() string {
    return "BTC-OP-USDT"
}

func (this OpUsdt) GetSymbolsToCheck() []string {
    return []string {
        BTC,
        OP,
        USDT,
    }
}

func (this OpUsdt) GetSymbols() []string {
    return []string {
        OPBTC,
        OPUSDT,
        BTCUSDT,
    }
}

func (this OpUsdt) GetSides() []api.Side {
    return []api.Side {
        api.BUY,
        api.SELL,
        api.BUY,
    }
}



type UsdtOp struct { }

func NewUsdtOp() UsdtOp {
    return UsdtOp {}
}

func (this UsdtOp) GetName() string {
    return "BTC-USDT-OP"
}

func (this UsdtOp) GetSymbolsToCheck() []string {
    return []string {
        BTC,
        USDT,
        OP,
    }
}

func (this UsdtOp) GetSymbols() []string {
    return []string {
        BTCUSDT,
        OPUSDT,
        OPBTC,
    }
}

func (this UsdtOp) GetSides() []api.Side {
    return []api.Side {
        api.SELL,
        api.BUY,
        api.SELL,
    }
}
