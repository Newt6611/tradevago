package max

import (
	"context"

	"github.com/Newt6611/tradevago/pkg/api"
)

func (m *Max) CreateOrder(ctx context.Context, side api.Side, orderType api.OrderType, pair string, price float64, baseAmount float64) (api.Order, error) {

    return api.Order{}, nil
}
