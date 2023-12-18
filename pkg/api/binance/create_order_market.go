package binance

import (
	"context"
	"fmt"

	"github.com/Newt6611/tradevago/pkg/api"
	b "github.com/adshao/go-binance/v2"
)

func (this *Binance) CreateOrderMarket(ctx context.Context, side api.Side, pair string, baseAmount float64, quoteAmount float64) (api.Order, error) {
    fmt.Println("binance create market order not implment yet")

    binanceSide := b.SideTypeBuy
    if side == api.SELL {
        binanceSide = b.SideTypeSell
    }
    this.binanceClient.NewCreateOrderService().
        NewOrderRespType(b.NewOrderRespTypeACK).
        Type(b.OrderTypeMarket).
        Side(binanceSide).
        Do(ctx)

    return api.Order{}, nil
}
