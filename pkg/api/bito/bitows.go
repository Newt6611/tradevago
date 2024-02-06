package bito

import (
	"github.com/bitoex/bitopro-api-go/pkg/ws"
)


type BitoWs struct {
    apiKey              string
    apiSecret           string
    apiws               *ws.Ws
}

func NewBitoWs(apiKey string, apiSecret string, email string) *BitoWs {
    return &BitoWs{
        apiKey: apiKey,
        apiSecret: apiSecret,
        apiws: ws.NewWs(email, apiKey, apiSecret, "wss://stream.bitopro.com:443"),
    }
}
