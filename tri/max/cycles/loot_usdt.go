package cycles

import (
	"github.com/Newt6611/tradevago/pkg/api"
	_ "github.com/Newt6611/tradevago/tri"
)

type LootUsdt struct { }

func NewLootUsdt() LootUsdt {
    return LootUsdt {}
}

func (this LootUsdt) GetName() string {
    return "TWD-LOOT-USDT"
}

func (this LootUsdt) GetSymbolsToCheck() []string {
    return []string {
        TWD,
        LOOT,
        USDT,
    }
}

func (this LootUsdt) GetSymbols() []string {
    return []string {
        LOOTTWD,
        LOOTUSDT,
        USDTTWD,
    }
}

func (this LootUsdt) GetSides() []api.Side {
    return []api.Side {
        api.BUY,
        api.SELL,
        api.SELL,
    }
}


type UsdtLoot struct { }

func NewUsdtLoot() UsdtLoot {
    return UsdtLoot {}
}

func (this UsdtLoot) GetName() string {
    return "TWD-USDT-LOOT"
}

func (this UsdtLoot) GetSymbolsToCheck() []string {
    return []string {
        TWD,
        USDT,
        LOOT,
    }
}

func (this UsdtLoot) GetSymbols() []string {
    return []string {
        USDTTWD,
        LOOTUSDT,
        LOOTTWD,
    }
}

func (this UsdtLoot) GetSides() []api.Side {
    return []api.Side {
        api.BUY,
        api.BUY,
        api.SELL,
    }
}
