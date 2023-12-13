package api

import "context"

type WsUserOrder struct {
	ID              string
    Pair            string
	Status          OrderStatus
}

type WsUserOrderDatas struct {
    Datas           map[string]WsUserOrder
    Err             error
}

func (this *WSApi) RunUserOrderConsumer(ctx context.Context) (chan WsUserOrderDatas, chan struct{}) {
    return this.client.RunUserOrderConsumer(ctx)
}
