package cycles

import (
	"github.com/Newt6611/tradevago/pkg/api"
	_ "github.com/Newt6611/tradevago/tri"
)

type EtcUsdt struct { }

func NewEtcUsdt() EtcUsdt {
    return EtcUsdt {}
}

func (this EtcUsdt) GetName() string {
    return "TWD-ETC-USDT"
}

func (this EtcUsdt) GetSymbolsToCheck() []string {
    return []string {
        TWD,
        ETC,
        USDT,
    }
}

func (this EtcUsdt) GetSymbols() []string {
    return []string {
        ETCTWD,
        ETCUSDT,
        USDTTWD,
    }
}

func (this EtcUsdt) GetSides() []api.Side {
    return []api.Side {
        api.BUY,
        api.SELL,
        api.SELL,
    }
}


type UsdtEtc struct { }

func NewUsdtEtc() UsdtEtc {
    return UsdtEtc {}
}

func (this UsdtEtc) GetName() string {
    return "TWD-USDT-ETC"
}

func (this UsdtEtc) GetSymbolsToCheck() []string {
    return []string {
        TWD,
        USDT,
        ETC,
    }
}

func (this UsdtEtc) GetSymbols() []string {
    return []string {
        USDTTWD,
        ETCUSDT,
        ETCTWD,
    }
}

func (this UsdtEtc) GetSides() []api.Side {
    return []api.Side {
        api.BUY,
        api.BUY,
        api.SELL,
    }
}
