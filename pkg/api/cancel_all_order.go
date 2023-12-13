package api

import "context"


type CancelAllOrderService struct {
    client  Client
    pair    string
}

func (c *Api) NewCancelAllOrderService() *CancelAllOrderService {
    return &CancelAllOrderService {
        client: c.client,
    }
}

func (this *CancelAllOrderService) WithPair(pair string) *CancelAllOrderService {
    this.pair = pair
    return this
}

func (this *CancelAllOrderService) Do(ctx context.Context) error {
    return this.client.CancelAllOrder(ctx, this.pair)
}
