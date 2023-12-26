package cycles

import "github.com/Newt6611/tradevago/tri"

const (
    BTCUSDT     string = "BTCUSDT"
    ADABTC      string = "ADABTC"
    ADAUSDT     string = "ADAUSDT"
)

const (
    BTC         string = "BTC"
    USDT        string = "USDT"
    ADA         string = "ADA"
)

func GetPairs() []string {
    return []string {
        BTCUSDT,

        ADABTC,
        ADAUSDT,
    }
}

func GetCycles() []tri.Cycle {
    return []tri.Cycle {
        NewAdaUsdt(),
        // NewUsdtAda(),
    }
}
