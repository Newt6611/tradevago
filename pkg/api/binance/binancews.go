package binance

type BinanceWS struct {
    apiKey string
    apiSecret string
}

func NewBinanceWs(apiKey string, apiSecret string) *BinanceWS {
    return &BinanceWS {
        apiKey: apiKey,
        apiSecret: apiSecret,
    }
}

