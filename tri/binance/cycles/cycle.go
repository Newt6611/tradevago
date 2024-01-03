package cycles

import "github.com/Newt6611/tradevago/tri"

const (
    BTCUSDT     string = "BTCUSDT"

    ADABTC      string = "ADABTC"
    ADAUSDT     string = "ADAUSDT"

    ACABTC      string = "ACABTC"
    ACAUSDT     string = "ACAUSDT"

    AAVEBTC     string = "AAVEBTC"
    AAVEUSDT    string = "AAVEUSDT"

    ACHBTC      string = "ACHBTC"
    ACHUSDT     string = "ACHUSDT"

    OPBTC       string = "OPBTC"
    OPUSDT      string = "OPUSDT"

    EOSBTC      string = "EOSBTC"
    EOSUSDT     string = "EOSUSDT"

    FILBTC      string = "FILBTC"
    FILUSDT     string = "FILUSDT"
)

const (
    BTC         string = "BTC"
    BNB         string = "BNB"

    USDT        string = "USDT"
    ADA         string = "ADA"
    ACA         string = "ACA"
    AAVE        string = "AAVE"
    ACH         string = "ACH"
    OP          string = "OP"
    EOS         string = "EOS"
    FIL         string = "FIL"
)

func GetPairs() []string {
    return []string {
        BTCUSDT,

        ADABTC,
        ADAUSDT,

        ACABTC,
        ACAUSDT,

        AAVEBTC,
        AAVEUSDT,

        ACHBTC,
        ACHUSDT,

        OPBTC,
        OPUSDT,

        EOSBTC,
        EOSUSDT,

        FILBTC,
        FILUSDT,
    }
}

func GetCycles() []tri.Cycle {
    return []tri.Cycle {
        // NewAdaUsdt(),
        // NewUsdtAda(),
        //
        // NewAcaUsdt(),
        // NewUsdtAca(),
        //
        // NewAaveUsdt(),
        // NewUsdtAave(),
        //
        // NewAchUsdt(),
        // NewUsdtAch(),

        NewUsdtOp(),
        NewOpUsdt(),

        // NewEosUsdt(),
        // NewUsdtEos(),
    }
}
