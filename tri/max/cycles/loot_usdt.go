package cycles

import (
	"github.com/Newt6611/tradevago/tri"
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

func (this LootUsdt) GetSides() []tri.Side {
    return []tri.Side {
        tri.BUY,
        tri.SELL,
        tri.SELL,
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

func (this UsdtLoot) GetSides() []tri.Side {
    return []tri.Side {
        tri.BUY,
        tri.BUY,
        tri.SELL,
    }
}
