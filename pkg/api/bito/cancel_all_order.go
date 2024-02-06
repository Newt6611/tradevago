package bito

import (
	"context"
	"errors"
)

func (this *Bito) CancelAllOrder(ctx context.Context, pair string) error {
    res := this.bitoapi.CancelAll(pair)
    if len(res.Error) != 0 {
        return errors.New(res.Error)
    }
    return nil
}
