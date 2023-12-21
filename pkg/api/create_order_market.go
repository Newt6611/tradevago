package api

import (
	"context"
	"errors"
)

type OrderStatus string

var (
    ErrorBalanceNotEnougth = errors.New("Error Balance Not Enougth")
)
const (
    OrderStatusWait     OrderStatus = "Wait"
    OrderStatusDone     OrderStatus = "Done"
    OrderStatusCancel   OrderStatus = "Cancel"

    OrderStatusUnknow   OrderStatus = "Unknow"
)

type Order struct {
    Id          string
    OrderStatus OrderStatus
}

type CreateOrderService struct {
    client          Client
    side            Side
    pair            string
    baseAmount      float64
    quoteAmount     float64
}

func (c *Api) NewCreateOrderMarketService() *CreateOrderService {
    return &CreateOrderService {
        client: c.client,
    }
}

func (this *CreateOrderService) WithSide(side Side) *CreateOrderService {
    this.side = side
    return this
}

func (this *CreateOrderService) WithPair(pair string) *CreateOrderService {
    this.pair = pair
    return this
}

func (this *CreateOrderService) WithQuoteAmount (quote float64) *CreateOrderService {
    this.quoteAmount = quote
    return this
}

func (this *CreateOrderService) WithBaseAmount(amount float64) *CreateOrderService {
    this.baseAmount = amount
    return this
}

func (this *CreateOrderService) Do(ctx context.Context) (Order, error){
    return this.client.CreateOrderMarket(ctx, this.side, this.pair, this.baseAmount, this.quoteAmount)
}
