package binance

import b "github.com/adshao/go-binance/v2"

type BinanceWS struct {
    apiKey string
    apiSecret string
    binanceClient   *b.Client
}

func NewBinanceWs(apiKey string, apiSecret string) *BinanceWS {
    client := b.NewClient(apiKey, apiSecret)
    return &BinanceWS {
        apiKey: apiKey,
        apiSecret: apiSecret,
        binanceClient: client,
    }
}

