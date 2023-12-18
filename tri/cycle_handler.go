package tri

import (
	"math"

	"github.com/Newt6611/tradevago/pkg/api"
)

func CycleHandler(apiClient *api.Api, depthHandler *DepthHandler, cycle Cycle) (float64, float64) {
    var rate float64 = 1
    symbols := cycle.GetSymbols()
    sides := cycle.GetSides()
    orderAmount := math.MaxFloat64

    for i := 0; i < 3; i++ {
        depth := depthHandler.GetDepth(symbols[i])
        rate = resolveAmount(apiClient, rate, depth, sides[i])
        orderAmount = resolveOrderAmount(sides[i], depth, orderAmount)
    }

    return rate, orderAmount
}

func resolveAmount(apiClient *api.Api, baseAmount float64, depth api.WsDepth, side api.Side) float64 {
    if side == api.BUY {
        baseAmount /= depth.Asks[0].Price
    } else if side == api.SELL {
        baseAmount *= depth.Bids[0].Price
    }
    baseAmount -= baseAmount * apiClient.GetTakerFee()
    return baseAmount
}

func resolveOrderAmount(side api.Side, depth api.WsDepth, orderAmount float64) float64 {
    if side == api.SELL {
        if orderAmount > depth.Bids[0].Amount{
            orderAmount = depth.Bids[0].Amount
        }
        orderAmount *= depth.Bids[0].Price
    } else if side == api.BUY {
        t := depth.Asks[0].Price * depth.Asks[0].Amount
        if orderAmount > t {
            orderAmount = t
        }
        orderAmount /= depth.Asks[0].Price
    }
    return orderAmount
}
