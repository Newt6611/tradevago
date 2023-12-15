package api

import "context"

type WsUserAccountBalance struct {
    Currency    string
    Balance     float64
}

type WsUserAccountDatas struct {
    Datas   []WsUserAccountBalance
    Err     error
}

func (this *WSApi) RunAccountConsumer(ctx context.Context) (chan WsUserAccountDatas, chan struct{}) {
    return this.client.RunAccountConsumer(ctx)
}
