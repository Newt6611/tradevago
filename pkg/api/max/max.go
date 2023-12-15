package max

import m "github.com/maicoin/max-exchange-api-go"

const MAX_API_ENDPOINT = "https://max-api.maicoin.com"

type apiErrorResponse struct {
	Success bool                `json:"success"`
	Error   apiErrorDetails     `json:"error"`
}

type apiErrorDetails struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Max struct {
    apiKey      string
    apiSecret   string
    takerFee    float64
    makerFee    float64
    maxapi      m.API
}

func NewMaxClient(apiKey string, apiSecret string, takerFee float64, makerFee float64) *Max {
    return &Max {
        apiKey: apiKey,
        apiSecret: apiSecret,
        takerFee: takerFee,
        makerFee: makerFee,
        maxapi: m.NewClient(m.AuthToken(apiKey, apiSecret)),
    }
}

func (m *Max) GetName() string {
    return "Max"
}

func (m *Max) GetTakerFee() float64 {
    return m.takerFee
}

func (m *Max) GetMakerFee() float64 {
    return m.makerFee
}
