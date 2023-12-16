package binance

import (
	"context"
	"strconv"

	"github.com/Newt6611/tradevago/pkg/api"
	b "github.com/adshao/go-binance/v2"
)


func (this *BinanceWS) RunDepthConsumer(ctx context.Context, pairs []string, depth int) (chan api.WsDepth, chan struct{}) {
    depthDataChan := make(chan api.WsDepth, 500)

    symbolLevels := map[string]string{}
    for _, pair := range pairs {
        symbolLevels[pair] = strconv.Itoa(depth)
    }
    _, stopC, err := b.WsCombinedPartialDepthServe(symbolLevels, partialDepthHandler(depthDataChan), partialDepthErrorHandler(depthDataChan))
    if err != nil {
        depthDataChan <- api.WsDepth{ Err: err }
    }

    return depthDataChan, stopC
}
func partialDepthHandler(depthDataChan chan api.WsDepth) func(event *b.WsPartialDepthEvent) {
    return func(event *b.WsPartialDepthEvent) {
        askPrice, askAmount, _ := event.Asks[0].Parse()
        bidPrice, bidAmount, _ := event.Bids[0].Parse()
        depthData := api.WsDepth {
            Pair: event.Symbol,
            Asks: []api.DepthInfo {
                {
                    Price: askPrice,
                    Amount: askAmount,
                },
            },
            Bids: []api.DepthInfo {
                {
                    Price: bidPrice,
                    Amount: bidAmount,
                },
            },

            Err: nil,
        }

        depthDataChan <- depthData
    }
}

func partialDepthErrorHandler(depthDataChan chan api.WsDepth) func(err error) {
    return func (err error) {
        depthDataChan <- api.WsDepth{ Err: err }
    }
}
