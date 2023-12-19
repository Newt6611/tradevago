package binance

import (
	"context"
	"strconv"

	"github.com/Newt6611/tradevago/pkg/api"
	b "github.com/adshao/go-binance/v2"
)

func (m *Binance) GetPairInfo(ctx context.Context, pairs []string) ([]api.PairInfo, error) {
    res, err := m.binanceClient.NewExchangeInfoService().Symbols(pairs...).Do(ctx)
    if err != nil {
        return []api.PairInfo{}, err
    }

    return mapPairInfo(res.Symbols), nil
}

func mapPairInfo(symbols []b.Symbol) []api.PairInfo {
    pairInfos := []api.PairInfo{}

    for _, symbol := range symbols {
        pairInfo := api.PairInfo {
            Name: symbol.Symbol,
            MarketStatus: symbol.Status,
            BaseUnit: symbol.BaseAsset,
            BaseUnitPrecision: symbol.BaseAssetPrecision,
            QuoteUnit: symbol.QuoteAsset,
            QuoteUnitPrecision: symbol.QuoteAssetPrecision,
        }
        minPriceStr := symbol.PriceFilter().MinPrice
        minBaseStr := symbol.LotSizeFilter().MinQuantity

        pairInfo.MinQuoteAmount, _ = strconv.ParseFloat(minPriceStr, 64)
        pairInfo.MinBaseAmount, _ = strconv.ParseFloat(minBaseStr, 64)

        pairInfos = append(pairInfos, pairInfo)
    }

    return pairInfos
}
