package cycles

import (
	"github.com/Newt6611/tradevago/pkg/api"
	_ "github.com/Newt6611/tradevago/tri"
)

type EosUsdt struct { }

func NewEosUsdt() EosUsdt {
    return EosUsdt {}
}

func (this EosUsdt) GetName() string {
    return "BTC-EOS-USDT"
}

func (this EosUsdt) GetSymbolsToCheck() []string {
    return []string {
        BTC,
        EOS,
        USDT,
    }
}

func (this EosUsdt) GetSymbols() []string {
    return []string {
        EOSBTC,
        EOSUSDT,
        BTCUSDT,
    }
}

func (this EosUsdt) GetSides() []api.Side {
    return []api.Side {
        api.BUY,
        api.SELL,
        api.BUY,
    }
}



type UsdtEos struct { }

func NewUsdtEos() UsdtEos {
    return UsdtEos {}
}

func (this UsdtEos) GetName() string {
    return "BTC-USDT-EOS"
}

func (this UsdtEos) GetSymbolsToCheck() []string {
    return []string {
        BTC,
        USDT,
        EOS,
    }
}

func (this UsdtEos) GetSymbols() []string {
    return []string {
        BTCUSDT,
        EOSUSDT,
        EOSBTC,
    }
}

func (this UsdtEos) GetSides() []api.Side {
    return []api.Side {
        api.SELL,
        api.BUY,
        api.SELL,
    }
}
