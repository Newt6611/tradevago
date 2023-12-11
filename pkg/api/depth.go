package api

import "context"

type DepthInfo struct {
    Price   float64
    Amount  float64
}

type Depth struct {
    Asks []DepthInfo
    Bids []DepthInfo
}

type DepthService struct {
    client Client
    Pair string
    Limit int
    SortByPrice bool
}

func (c *Api) NewDepthService() *DepthService {
    return &DepthService {
        client: c.client,
        Limit: 300,
        SortByPrice: true,
    }
}

func (this *DepthService) WithPair(pair string) *DepthService {
    this.Pair = pair
    return this
}

func (this *DepthService) WithLimit(limit int) *DepthService {
    this.Limit = limit
    return this
}

func (this *DepthService) WithSortByPrice(sort bool) *DepthService {
    this.SortByPrice = sort
    return this
}

func (this *DepthService) Do(ctx context.Context) (Depth, error) {
    return this.client.GetDepth(ctx, this)
}
