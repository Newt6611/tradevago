package api

import "context"

type Order struct {
    Id string
}

type CreateOrderService struct {
    client          Client
    side            Side
    orderType       OrderType
    pair            string
    price           float64
    baseAmount      float64
}

func (c *Api) NewCreateOrderService() *CreateOrderService {
    return &CreateOrderService {
        client: c.client,
    }
}

func (this *CreateOrderService) WithSide(side Side) *CreateOrderService {
    this.side = side
    return this
}

func (this *CreateOrderService) WithOrderType(orderType OrderType) *CreateOrderService {
    this.orderType = orderType
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
    return this.client.CreateOrder(ctx, this.side, this.orderType, this.pair, this.price, this.baseAmount)
}
