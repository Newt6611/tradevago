package bito

import (
	"context"
	"fmt"
	"strconv"

	"github.com/Newt6611/tradevago/pkg/api"
	"github.com/bitoex/bitopro-api-go/pkg/ws"
)


func (this *BitoWs) RunDepthConsumer(ctx context.Context, pairs []string, depth int) (chan api.WsDepth, chan struct{}) {
    formatedPairs := []string{}

    for _, pair := range pairs {
        p := fmt.Sprintf("%s:%d", pair, depth)
        formatedPairs = append(formatedPairs, p)
    }

    depthDataCh := make(chan api.WsDepth, 500)
    closeC := make(chan struct{})

    go func() {
        orderBooksCh, bitoCloseCh := this.apiws.RunOrderbookWsConsumer(ctx, formatedPairs)

        for {
            select {
            case <-closeC:
                close(bitoCloseCh)
                goto done
            case orderBooks := <- orderBooksCh:
                if orderBooks.Err != nil {
                    depthDataCh <- api.WsDepth { Err: orderBooks.Err }
                } else {
                    d := mapDepthData(orderBooks)
                    depthDataCh <- d
                }
            }
        }
        done:
    }()

    return depthDataCh, closeC
}

func mapDepthData(input ws.OrderBookData) (data api.WsDepth) {
    // Price   float64
    // Amount  float64
    askPrice, _ := strconv.ParseFloat(input.Asks[0].Price, 64)
    askAmount, _ := strconv.ParseFloat(input.Asks[0].Amount, 64)
    bidPrice, _ := strconv.ParseFloat(input.Bids[0].Price, 64)
    bidAmount, _ := strconv.ParseFloat(input.Bids[0].Amount, 64)

    return api.WsDepth{
        Pair: input.Pair,
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
    }
}
