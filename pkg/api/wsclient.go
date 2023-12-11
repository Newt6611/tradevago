package api

import "context"

type WSClient interface {
    RunDepthConsumer(ctx context.Context, pairs []string, depth int) (chan WsDepth, chan struct{})
}

type WSApi struct {
    client WSClient
}

func NewWsApi(client WSClient) *WSApi {
    return &WSApi {
        client: client,
    }
}
