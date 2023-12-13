package max

import (
	"context"

	m "github.com/maicoin/max-exchange-api-go"
)

func (this *Max) CancelAllOrder(ctx context.Context, pair string) error {
    _, err := this.maxapi.CancelOrders(ctx, m.Market(pair))
    return err
}
