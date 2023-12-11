package cycles

import (
	"github.com/Newt6611/tradevago/tri"
	_ "github.com/Newt6611/tradevago/tri"
)

type EthUsdt struct { }

func NewEthUsdt() EthUsdt {
    return EthUsdt {}
}

func (this EthUsdt) GetName() string {
    return "TWD-ETH-USDT"
}

func (this EthUsdt) GetSymbolsToCheck() []string {
    return []string {
        TWD,
        ETH,
        USDT,
    }
}

func (this EthUsdt) GetSymbols() []string {
    return []string {
        ETHTWD,
        ETHUSDT,
        USDTTWD,
    }
}

func (this EthUsdt) GetSides() []tri.Side {
    return []tri.Side {
        tri.BUY,
        tri.SELL,
        tri.SELL,
    }
}


type UsdtEth struct { }

func NewUsdtEth() UsdtEth {
    return UsdtEth {}
}

func (this UsdtEth) GetName() string {
    return "TWD-USDT-ETH"
}

func (this UsdtEth) GetSymbolsToCheck() []string {
    return []string {
        TWD,
        USDT,
        ETH,
    }
}

func (this UsdtEth) GetSymbols() []string {
    return []string {
        USDTTWD,
        ETHUSDT,
        ETHTWD,
    }
}

func (this UsdtEth) GetSides() []tri.Side {
    return []tri.Side {
        tri.BUY,
        tri.BUY,
        tri.SELL,
    }
}
