package api

import "context"

type PairInfo struct {
	Name               string
	MarketStatus       string
	BaseUnit           string
	QuoteUnit          string
	QuoteUnitPrecision int
	MinBaseAmount      float64
	BaseUnitPrecision  int
	MinQuoteAmount     float64

	Binance            bool
	ApplyMaxToMarket   bool
	ApplyMinToMarket   bool
	MaxNotional        float64
	MinNotional        float64
	StepSize           float64
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
