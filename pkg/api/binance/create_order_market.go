package binance

import (
	"context"
	"strconv"

	"github.com/Newt6611/tradevago/pkg/api"
	b "github.com/adshao/go-binance/v2"
)

func (this *Binance) CreateOrderMarket(ctx context.Context, side api.Side, pair string, baseAmount float64, quoteAmount float64) (api.Order, error) {
    orderService := this.binanceClient.NewCreateOrderService().
        NewOrderRespType(b.NewOrderRespTypeACK).
        Symbol(pair).
        Type(b.OrderTypeMarket)

    if side == api.BUY {
        orderService.QuoteOrderQty(strconv.FormatFloat(quoteAmount, 'f', -1, 32)).Side(b.SideTypeBuy)
    } else {
        orderService.Quantity(strconv.FormatFloat(baseAmount, 'f', -1, 32)).Side(b.SideTypeSell)
    }

    order, err := orderService.Do(ctx)

    if err != nil {
        return api.Order{}, err
    }

    return api.Order {
        Id: strconv.FormatInt(order.OrderID, 10),
        OrderStatus: getOrderStatus(string(order.Status)),
    }, nil
}
