package cycles

import "github.com/Newt6611/tradevago/pkg/api"

type EosUsdt struct { }

func NewEosUsdt() EosUsdt {
    return EosUsdt{}
}

func (this EosUsdt) GetName() string {
    return "TWD-EOS-USDT"
}

func (this EosUsdt) GetSymbolsToCheck() []string {
    return []string {
        TWD,
        EOS,
        USDT,
    }
}

func (this EosUsdt) GetSymbols() []string {
    return []string {
        EOSTWD,
        EOSUSDT,
        USDTTWD,
    }
}

func (this EosUsdt) GetSides() []api.Side {
    return []api.Side {
        api.BUY,
        api.SELL,
        api.SELL,
    }
}

type UsdtEos struct { }

func NewUsdtEos() UsdtEos {
    return UsdtEos {}
}

func (this UsdtEos) GetName() string {
    return "TWD-USDT-EOS"
}

func (this UsdtEos) GetSymbolsToCheck() []string {
    return []string {
        TWD,
        USDT,
        EOS,
    }
}

func (this UsdtEos) GetSymbols() []string {
    return []string {
        USDTTWD,
        EOSUSDT,
        EOSTWD,
    }
}

func (this UsdtEos) GetSides() []api.Side {
    return []api.Side {
        api.BUY,
        api.BUY,
        api.SELL,
    }
}
