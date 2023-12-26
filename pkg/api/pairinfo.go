package api

import "context"

type PairInfo struct {
	Name                 string
    StepSize             float64 // Binance
	MarketStatus         string
	BaseUnit             string
	BaseUnitPrecision    int
	MinBaseAmount        float64
	QuoteUnit            string
	QuoteUnitPrecision   int
	MinQuoteAmount       float64
}


type PairInfoService struct {
    client Client
    pairs []string
}


func (c *Api) NewPairInfoService() *PairInfoService {
    return &PairInfoService{
        client: c.client,
    }
}

func (this *PairInfoService) WithPairs(pairs []string) *PairInfoService {
    this.pairs = pairs
    return this
}

func (this *PairInfoService) Do(ctx context.Context) ([]PairInfo, error) {
    return this.client.GetPairInfo(ctx, this.pairs)
}
