package binance

import (
	"context"
	"fmt"

	"github.com/Newt6611/tradevago/pkg/api"
)

func (this *Binance) CreateOrderMarket(ctx context.Context, side api.Side, pair string, baseAmount float64, quoteAmount float64) (api.Order, error) {
    fmt.Println("binance create market order not implment yet")

    return api.Order{}, nil
}
