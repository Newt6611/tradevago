package cycles

import "github.com/Newt6611/tradevago/tri"

const (
    TWD         string = "TWD"
    BITO        string = "BITO"

    USDT        string = "USDT"
    BTC         string = "BTC"
    ETH         string = "ETH"
    SOL         string = "SOL"
    DOGE        string = "DOGE"
    USDC        string = "USDC"
    ADA         string = "ADA"
    XRP         string = "XRP"
    APE         string = "APE"
    MATIC       string = "MATIC"
    TRX         string = "TRX"
    MV          string = "MV"
    LTC         string = "LTC"
    EOS         string = "EOS"
    SHIB        string = "SHIB"
    BNB         string = "BNB"
    BCH         string = "BCH"
    TON         string = "TON"
    YFI         string = "YFI"
)

const (
    BTCTWD      string = "btc_twd"
    BITOTWD     string = "bito_twd"

    BTCUSDT     string = "btc_usdt"
    USDTTWD     string = "usdt_twd"
    ETHTWD      string = "eth_twd"
    ETHUSDT     string = "eth_usdt"
    SOLTWD      string = "sol_twd"
    SOLUSDT     string = "sol_usdt"
    DOGETWD     string = "doge_twd"
    DOGEUSDT    string = "doge_usdt"
    USDCTWD     string = "usdc_twd"
    USDCUSDT    string = "usdc_usdt"
    ADATWD      string = "ada_twd"
    ADAUSDT     string = "ada_usdt"
    XRPTWD      string = "xrp_twd"
    APETWD      string = "ape_twd"
    APEUSDT     string = "ape_usdt"
    MATICTWD    string = "matic_twd"
    MATICUSDT   string = "matic_usdt"
    TRXTWD      string = "trx_twd"
    MVTWD       string = "mv_twd"
    MVUSDT      string = "mv_usdt"
    LTCTWD      string = "ltc_twd"
    LTCUSDT     string = "ltc_usdt"
    EOSTWD      string = "eos_twd"
    EOSUSDT     string = "eos_usdt"
    SHIBTWD     string = "shib_twd"
    SHIBUSDT    string = "shib_usdt"
    BNBTWD      string = "bnb_twd"
    BCHTWD      string = "bch_twd"
    BCHUSDT     string = "bch_usdt"
    TONTWD      string = "ton_twd"
    TONUSDT     string = "ton_usdt"
    YFITWD      string = "yfi_twd"
    YFIUSDT     string = "yfi_usdt"
)

func GetPairs() []string {
    return []string {
        BITOTWD,
        BTCTWD,
        BTCUSDT,
        USDTTWD,
        ETHTWD,
        ETHUSDT,
        SOLTWD,
        SOLUSDT,
        DOGETWD,
        DOGEUSDT,
        USDCTWD,
        USDCUSDT,
        ADATWD,
        ADAUSDT,
        XRPTWD,
        APETWD,
        APEUSDT,
        MATICTWD,
        MATICUSDT,
        TRXTWD,
        MVTWD,
        MVUSDT,
        LTCTWD,
        LTCUSDT,
        EOSTWD,
        EOSUSDT,
        SHIBTWD,
        SHIBUSDT,
        BNBTWD,
        BCHTWD,
        BCHUSDT,
        TONTWD,
        TONUSDT,
        YFITWD,
        YFIUSDT,
    }
}

func GetCycles() []tri.Cycle {
    return []tri.Cycle {
        NewBtcUsdt(),
        NewUsdtBtc(),

        NewAdaUsdt(),
        NewUsdtAda(),

        NewApeUsdt(),
        NewUsdtApe(),

        NewBchUsdt(),
        NewUsdtBch(),

        NewDogeUsdt(),
        NewUsdtDoge(),

        NewEosUsdt(),
        NewUsdtEos(),

        NewEthUsdt(),
        NewUsdtEth(),

        NewLtcUsdt(),
        NewUsdtLtc(),

        NewMaticUsdt(),
        NewUsdtMatic(),

        NewMvUsdt(),
        NewUsdtMv(),

        NewShibUsdt(),
        NewUsdtShib(),

        NewSolUsdt(),
        // NewUsdtSol(),

        NewTonUsdt(),
        NewUsdtTon(),

        NewUsdcUsdt(),
        NewUsdtUsdc(),

        NewYfiUsdt(),
        NewUsdtYfi(),
    }
}
