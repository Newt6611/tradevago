package bito

import "github.com/bitoex/bitopro-api-go/pkg/bitopro"

type Bito struct {
	bitoapi   bitopro.AuthAPI
	apiKey    string
	apiSecret string
	email     string
	takerFee  float64
	makerFee  float64
    proxy     string
}

func NewBitoClient(apiKey string, apiSecret string, takerFee float64, makerFee float64, email string) *Bito {
    return &Bito {
        apiKey: apiKey,
        apiSecret: apiSecret,
        takerFee: takerFee,
        makerFee: makerFee,
        email: email,
        bitoapi: *bitopro.GetAuthClient(email, apiKey, apiSecret),
        proxy: "https://api.bitopro.com",
    }
}

func (m *Bito) GetName() string {
    return "Bito"
}

func (m *Bito) GetTakerFee() float64 {
    return m.takerFee
}

func (m *Bito) GetMakerFee() float64 {
    return m.makerFee
}
