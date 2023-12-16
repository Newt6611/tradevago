package binance

import (
	"github.com/adshao/go-binance/v2"
)

type Binance struct {
    binanceClient   *binance.Client
    apiKey          string
    apiSecret       string
    takerFee        float64
    makerFee        float64
}

func NewBinance(apiKey string, apiSecret string, takerFee float64, makerFee float64) *Binance {
    client := binance.NewClient(apiKey, apiSecret)
    return &Binance {
        binanceClient: client,
        apiKey: apiKey,
        apiSecret: apiSecret,
        takerFee: takerFee,
        makerFee: makerFee,
    }
}

func (this *Binance) GetName() string {
    return "Binance"
}

func (m *Binance) GetTakerFee() float64 {
    return m.takerFee
}

func (m *Binance) GetMakerFee() float64 {
    return m.makerFee
}
