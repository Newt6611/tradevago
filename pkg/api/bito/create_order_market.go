package bito

import (
	"context"
	"errors"
	"fmt"
	"math/rand"

	"github.com/Newt6611/tradevago/pkg/api"
	"github.com/bitoex/bitopro-api-go/pkg/bitopro"
)


func (this *Bito) CreateOrderMarket(ctx context.Context, side api.Side, pair string, price float64, baseAmount float64, quoteAmount float64) (api.Order, error) {
    // The base amount of the order for the trading pair. 
    // However, when executing a market buy order, this field represents the order's quote value. 
    // Please check this doc for limitation.
    // https://github.com/bitoex/bitopro-offical-api-docs/blob/master/api/v3/private/create_an_order.md
    id := rand.Intn(1000)
    var bitoOrder *bitopro.CreateOrder
    if (side == api.BUY) {
        // bitoOrder = this.bitoapi.CreateOrderLimitBuy(id, pair, fmt.Sprintf("%v", price), fmt.Sprintf("%v", baseAmount))
        bitoOrder = this.bitoapi.CreateOrderMarketBuy(id, pair, fmt.Sprintf("%v", quoteAmount))
    } else if (side == api.SELL){
        // bitoOrder = this.bitoapi.CreateOrderLimitSell(id, pair, fmt.Sprintf("%v", price), fmt.Sprintf("%v", baseAmount))
        bitoOrder = this.bitoapi.CreateOrderMarketSell(id, pair, fmt.Sprintf("%v", baseAmount))
    }
    if len(bitoOrder.StatusCode.Error) != 0 {
        return api.Order{}, errors.New(bitoOrder.StatusCode.Error)
    }
    order := api.Order {
        Id: bitoOrder.OrderID,
        OrderStatus: api.OrderStatusWait,
    }

    return order, nil
}
