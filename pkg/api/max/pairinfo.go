package max

import (
	"context"
	"encoding/json"

	"github.com/Newt6611/tradevago/internal"
	"github.com/Newt6611/tradevago/pkg/api"
)

type pairInfo struct {
	ID                   string  `json:"id"`
	Name                 string  `json:"name"`
	MarketStatus         string  `json:"market_status"`
	BaseUnit             string  `json:"base_unit"`
	BaseUnitPrecision    int     `json:"base_unit_precision"`
	MinBaseAmount        float64 `json:"min_base_amount"`
	QuoteUnit            string  `json:"quote_unit"`
	QuoteUnitPrecision   int     `json:"quote_unit_precision"`
	MinQuoteAmount       float64 `json:"min_quote_amount"`
	MWalletSupported     bool    `json:"m_wallet_supported"`
}

func (m *Max) GetPairInfo(ctx context.Context) ([]api.PairInfo, error) {
    res, err := internal.Get(ctx, MAX_API_ENDPOINT +  "/api/v2/markets", nil)
    if err != nil {
        return []api.PairInfo{}, err
    }

    var data []pairInfo
	if err = json.Unmarshal([]byte(res), &data); err != nil {
        return []api.PairInfo{}, err
	}

    return mapPairInfo(data), nil
}

func mapPairInfo(res []pairInfo) []api.PairInfo {
    var pairinfos = []api.PairInfo{}
    for _, data := range res {
        pairinfos = append(pairinfos, api.PairInfo {
            Name: data.Name,
            MarketStatus: data.MarketStatus,
            BaseUnit: data.BaseUnit,
            BaseUnitPrecision: data.BaseUnitPrecision,
            MinBaseAmount: data.MinBaseAmount,
            QuoteUnit: data.QuoteUnit,
            QuoteUnitPrecision: data.QuoteUnitPrecision,
            MinQuoteAmount: data.MinQuoteAmount,
        })
    }
    return pairinfos
}
