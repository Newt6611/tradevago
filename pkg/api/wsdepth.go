package api

import "context"

type WsDepth struct {
    Pair string
    Asks []DepthInfo
    Bids []DepthInfo
    Err error
}

func (this *WSApi) RunDepthConsumer(ctx context.Context, pairs []string, depth int) (chan WsDepth, chan struct{}) {
    return this.client.RunDepthConsumer(ctx, pairs, depth)
}
