package api

import "context"

type OrderStatus string

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
    price           float64
    baseAmount      float64
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

func (this *CreateOrderService) WithPrice(price float64) *CreateOrderService {
    this.price = price
    return this
}

func (this *CreateOrderService) WithBaseAmount(amount float64) *CreateOrderService {
    this.baseAmount = amount
    return this
}

func (this *CreateOrderService) Do(ctx context.Context) (Order, error){
    return this.client.CreateOrderMarket(ctx, this.side, this.pair, this.price, this.baseAmount)
}
