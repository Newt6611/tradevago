package api

import "context"

type Client interface {
    GetName() string
    GetTakerFee() float64
    GetMakerFee() float64
    GetDepth(ctx context.Context, depthService *DepthService) (Depth, error)
    GetPairInfo(ctx context.Context) ([]PairInfo, error)
    CreateOrder(ctx context.Context, side Side, orderType OrderType, pair string, price float64, baseAmount float64) (Order, error)
}

type Api struct {
    client Client
}

func NewApi(client Client) *Api {
    return &Api {
        client: client,
    }
}

func (this *Api) GetName() string {
    return this.client.GetName()
}

func (this *Api) GetTakerFee() float64 {
    return this.client.GetTakerFee()
}

func (this *Api) GetMakerFee() float64 {
    return this.client.GetMakerFee()
}
