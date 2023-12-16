package binance

import "context"


func (this *Binance) CancelAllOrder(ctx context.Context, pair string) error {
    _, err := this.binanceClient.NewCancelOpenOrdersService().Symbol(pair).Do(ctx)
    return err
}
