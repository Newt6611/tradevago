package binance

import (
	"context"

	"github.com/Newt6611/tradevago/pkg/api"
	"github.com/adshao/go-binance/v2"
)

func (this *Binance) GetDepth(ctx context.Context, depthService *api.DepthService) (api.Depth, error) {
    res, err := this.binanceClient.NewDepthService().Symbol(depthService.Pair).Limit(depthService.Limit).Do(ctx)
    if err != nil {
        return api.Depth{}, err
    }
    return mapDepth(res), nil
}

func mapDepth(res *binance.DepthResponse) api.Depth {
    depth := api.Depth{}

    for i := range res.Asks {
        askPrice, askAmount, _ := res.Asks[i].Parse()
        bidPrice, bidAmount, _ := res.Bids[i].Parse()

        depth.Asks = append(depth.Asks, api.DepthInfo {
            Price: askPrice,
            Amount: askAmount,
        })

        depth.Bids = append(depth.Bids, api.DepthInfo {
            Price: bidPrice,
            Amount: bidAmount,
        })
    }
    return depth
}
