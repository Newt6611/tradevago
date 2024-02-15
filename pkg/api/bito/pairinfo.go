package bito

import (
	"context"
	"encoding/json"
    "strconv"

	"github.com/Newt6611/tradevago/internal"
	"github.com/Newt6611/tradevago/pkg/api"
)

type tradingPairInfo struct {
    Pair                        string  `json:"pair"`
    Base                        string  `json:"base"`
    Quote                       string  `json:"quote"`
    BasePrecision               string  `json:"basePrecision"`
    QuotePrecision              string  `json:"quotePrecision"`
    MinLimitBaseAmount          string  `json:"minLimitBaseAmount"`
    MaxLimitBaseAmount          string  `json:"maxLimitBaseAmount"`
    MinMarketBuyQuoteAmount     string  `json:"minMarketBuyQuoteAmount"`
    OrderOpenLimit              string  `json:"orderOpenLimit"`
    Maintain                    bool    `json:"maintain"`
    OrderBookQuotePrecision     string  `json:"orderBookQuotePrecision"`
    OrderBookQuoteScaleLevel    string  `json:"orderBookQuoteScaleLevel"`
    AmountPrecision             string  `json:"amountPrecision"`
}

type tradingPairInfos struct {
    Data []tradingPairInfo `json:"data,omitempty"`
}

func (m *Bito) GetPairInfo(ctx context.Context, pairs []string) ([]api.PairInfo, error) {
    path := m.proxy + "/v3/provisioning/trading-pairs"
    bs, err := internal.Get(ctx, path, nil, nil)
    var data tradingPairInfos
    if err != nil {
        return nil, err
    } else {
        if err = json.Unmarshal([]byte(bs), &data); err != nil {
            return nil, err
        }
    }
    info := mapTradingInfo(&data)
    return info, nil
}

func mapTradingInfo(tradingPairInfos *tradingPairInfos) (infos []api.PairInfo) {
    for _, data := range tradingPairInfos.Data {
        baseUnitPrecision, _ := strconv.Atoi(data.BasePrecision)
        minBaseAmount, _ := strconv.ParseFloat(data.MinLimitBaseAmount, 64)
        quoteUnitPrecision, _ := strconv.Atoi(data.AmountPrecision)
        minQuoteAmount, _ := strconv.ParseFloat(data.MinMarketBuyQuoteAmount, 64)

        infos = append(infos, api.PairInfo {
            Name: data.Pair,
            BaseUnit: data.Base,
            BaseUnitPrecision: baseUnitPrecision,
            MinBaseAmount: minBaseAmount,
            QuoteUnit: data.Quote,
            QuoteUnitPrecision: quoteUnitPrecision,
            MinQuoteAmount: minQuoteAmount,
        })
    }
    return infos
}
