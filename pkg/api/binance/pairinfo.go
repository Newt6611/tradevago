package binance

import (
	"context"
	"fmt"

	"github.com/Newt6611/tradevago/pkg/api"
)

func (m *Binance) GetPairInfo(ctx context.Context, pairs []string) ([]api.PairInfo, error) {
    fmt.Println("binance get pair info not implment yet")
    m.binanceClient.NewExchangeInfoService().Symbols(pairs...).Do(ctx)

    return []api.PairInfo{}, nil
}
