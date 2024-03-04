package cycles

import (
	"github.com/Newt6611/tradevago/tri"
)

const (
    TWD         string = "MAX_TWD"
    MAX         string = "MAX_MAX"
    USDT        string = "MAX_USDT"
    BTC         string = "MAX_BTC"
    ETH         string = "MAX_ETH"
    LTC         string = "MAX_LTC"
    BCH         string = "MAX_BCH"
    XRP         string = "MAX_XRP"
    BCNT        string = "MAX_BCNT"
    USDC        string = "MAX_USDC"
    LINK        string = "MAX_LINK"
    COMP        string = "MAX_COMP"
    DOGE        string = "MAX_DOGE"
    ADA         string = "MAX_ADA"
    DOT         string = "MAX_DOT"
    MATIC       string = "MAX_MATIC"
    SOL         string = "MAX_SOL"
    SHIB        string = "MAX_SHIB"
    SAND        string = "MAX_SAND"
    RLY         string = "MAX_RLY"
    LOOT        string = "MAX_LOOT"
    APE         string = "MAX_APE"
    BNB         string = "MAX_BNB"
    ETC         string = "MAX_ETC"
    ARB         string = "MAX_ARB"
)

const (
    MAXTWD          string = "maxtwd"
    USDTTWD         string = "usdttwd"

    BTCTWD          string = "btctwd"
    BTCUSDT         string = "btcusdt"

    ETHTWD          string = "ethtwd"
    ETHUSDT         string = "ethusdt"

    LTCTWD          string = "ltctwd"
    LTCUSDT         string = "ltcusdt"

    BCHTWD          string = "bchtwd"
    BCHUSDT         string = "bchusdt"

    XRPTWD          string = "xrptwd"
    XRPUSDT         string = "xrpusdt"

    BCNTTWD         string = "bcnttwd"
    BCNTUSDT        string = "bcntusdt"

    USDCTWD         string = "usdctwd"
    USDCUSDT        string = "usdcusdt"

    LINKTWD         string = "linktwd"
    LINKUSDT        string = "linkusdt"

    COMPTWD         string = "comptwd"
    COMPUSDT        string = "compusdt"

    DOGETWD         string = "dogetwd"
    DOGEUSDT        string = "dogeusdt"

    ADATWD          string = "adatwd"
    ADAUSDT         string = "adausdt"

    DOTTWD          string = "dottwd"
    DOTUSDT         string = "dotusdt"

    MATICTWD        string = "matictwd"
    MATICUSDT       string = "maticusdt"

    SOLTWD          string = "soltwd"
    SOLUSDT         string = "solusdt"

    SHIBTWD         string = "shibtwd"
    SHIBUSDT        string = "shibusdt"

    SANDTWD         string = "sandtwd"
    SANDUSDT        string = "sandusdt"

    RLYTWD          string = "rlytwd"
    RLYUSDT         string = "rlyusdt"

    LOOTTWD         string = "loottwd"
    LOOTUSDT        string = "lootusdt"

    APETWD          string = "apetwd"
    APEUSDT         string = "apeusdt"

    BNBTWD          string = "bnbtwd"
    BNBUSDT         string = "bnbusdt"

    ETCTWD          string = "etctwd"
    ETCUSDT         string = "etcusdt"

    ARBTWD          string = "arbtwd"
    ARBUSDT         string = "arbusdt"
)

func GetPairs() []string {
    return []string {
        MAXTWD,
        USDTTWD,

        BTCTWD,
        BTCUSDT,

        ETHTWD,
        ETHUSDT,

        LTCTWD,
        LTCUSDT,

        BCHTWD,
        BCHUSDT,

        XRPTWD,
        XRPUSDT,

        BCNTTWD,
        BCNTUSDT,

        USDCTWD,
        USDCUSDT,

        LINKTWD,
        LINKUSDT,

        COMPTWD,
        COMPUSDT,

        DOGETWD,
        DOGEUSDT,

        ADATWD,
        ADAUSDT,

        DOTTWD,
        DOTUSDT,

        MATICTWD,
        MATICUSDT,

        SOLTWD,
        SOLUSDT,

        SHIBTWD,
        SHIBUSDT,

        SANDTWD,
        SANDUSDT,

        RLYTWD,
        RLYUSDT,

        LOOTTWD,
        LOOTUSDT,

        APETWD,
        APEUSDT,

        BNBTWD,
        BNBUSDT,

        ETCTWD,
        ETCUSDT,

        ARBTWD,
        ARBUSDT,
    }
}

func GetCycles() []tri.Cycle {
    return []tri.Cycle {
        NewBtcUsdt(),
        NewUsdtBtc(),

        NewEthUsdt(),
        NewUsdtEth(),

        NewLtcUsdt(),
        NewUsdtLtc(),

        NewBchUsdt(),
        NewUsdtBch(),

        NewXrpUsdt(),
        NewUsdtXrp(),

        NewBcntUsdt(),
        NewUsdtBcnt(),

        NewUsdcUsdt(),
        NewUsdtUsdc(),

        NewLinkUsdt(),
        NewUsdtLink(),

        NewCompUsdt(),
        NewUsdtComp(),

        NewDogeUsdt(),
        NewUsdtDoge(),

        NewAdaUsdt(),
        NewUsdtAda(),

        NewDotUsdt(),
        NewUsdtDot(),

        NewMaticUsdt(),
        NewUsdtMatic(),

        NewSolUsdt(),
        NewUsdtSol(),

        NewShibUsdt(),
        NewUsdtShib(),

        NewSandUsdt(),
        NewUsdtSand(),

        NewRlyUsdt(),
        NewUsdtRly(),

        NewLootUsdt(),
        NewUsdtLoot(),

        NewApeUsdt(),
        NewUsdtApe(),

        NewBnbUsdt(),
        NewUsdtBnb(),

        NewEtcUsdt(),
        NewUsdtEtc(),

        NewArbUsdt(),
        NewUsdtArb(),
    }
}
