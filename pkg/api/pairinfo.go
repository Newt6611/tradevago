package api

import "context"

type PairInfo struct {
	Name                 string  `json:"name"`
	MarketStatus         string  `json:"market_status"`
	BaseUnit             string  `json:"base_unit"`
	BaseUnitPrecision    int     `json:"base_unit_precision"`
	MinBaseAmount        float64 `json:"min_base_amount"`
	QuoteUnit            string  `json:"quote_unit"`
	QuoteUnitPrecision   int     `json:"quote_unit_precision"`
	MinQuoteAmount       float64 `json:"min_quote_amount"`
}


type PairInfoService struct {
    client Client
}


func (c *Api) NewPairInfoService() *PairInfoService {
    return &PairInfoService{
        client: c.client,
    }
}

func (this *PairInfoService) Do(ctx context.Context) ([]PairInfo, error) {
    return this.client.GetPairInfo(ctx)
}
