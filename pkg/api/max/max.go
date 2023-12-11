package max

const MAX_API_ENDPOINT = "https://max-api.maicoin.com/"

type Max struct {
    apiKey      string
    apiSecret   string
    takerFee    float64
    makerFee    float64
}

func NewMaxClient(apiKey string, apiSecret string, takerFee float64, makerFee float64) *Max {
    return &Max {
        apiKey: apiKey,
        apiSecret: apiSecret,
        takerFee: takerFee,
        makerFee: makerFee,
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
