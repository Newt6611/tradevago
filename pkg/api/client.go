package api

import "context"

type Client interface {
    GetName() string
    GetTakerFee() float64
    GetMakerFee() float64
    GetDepth(ctx context.Context, depthService *DepthService) (Depth, error)
    GetPairInfo(ctx context.Context, pairs []string) ([]PairInfo, error)
    CreateOrderMarket(ctx context.Context, side Side, pair string, baseAmount float64, quoteAmount float64) (Order, error)
    CancelAllOrder(ctx context.Context, pair string) (error)
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
