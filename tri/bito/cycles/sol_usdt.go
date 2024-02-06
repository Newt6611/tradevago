package cycles

import "github.com/Newt6611/tradevago/pkg/api"

type SolUsdt struct { }

func NewSolUsdt() SolUsdt {
    return SolUsdt{}
}

func (this SolUsdt) GetName() string {
    return "TWD-SOL-USDT"
}

func (this SolUsdt) GetSymbolsToCheck() []string {
    return []string {
        TWD,
        SOL,
        USDT,
    }
}

func (this SolUsdt) GetSymbols() []string {
    return []string {
        SOLTWD,
        SOLUSDT,
        USDTTWD,
    }
}

func (this SolUsdt) GetSides() []api.Side {
    return []api.Side {
        api.BUY,
        api.SELL,
        api.SELL,
    }
}

type UsdtSol struct { }

func NewUsdtSol() UsdtSol {
    return UsdtSol {}
}

func (this UsdtSol) GetName() string {
    return "TWD-USDT-SOL"
}

func (this UsdtSol) GetSymbolsToCheck() []string {
    return []string {
        TWD,
        USDT,
        SOL,
    }
}

func (this UsdtSol) GetSymbols() []string {
    return []string {
        USDTTWD,
        SOLUSDT,
        SOLTWD,
    }
}

func (this UsdtSol) GetSides() []api.Side {
    return []api.Side {
        api.BUY,
        api.BUY,
        api.SELL,
    }
}
