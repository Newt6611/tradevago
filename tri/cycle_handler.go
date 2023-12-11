package tri

import "github.com/Newt6611/tradevago/pkg/api"

func CycleHandler(apiClient *api.Api, depthHandler *DepthHandler, cycle Cycle) (float64, float64) {
    var baseAmount float64 = 1
    symbols := cycle.GetSymbols()
    sides := cycle.GetSides()

    // 1
    depth := depthHandler.GetDepth(symbols[0])
    baseAmount = resolveAmount(apiClient, baseAmount, depth, sides[0])
    orderAmount := depth.Asks[0].Amount

    // 2
    depth = depthHandler.GetDepth(symbols[1])
    baseAmount = resolveAmount(apiClient, baseAmount, depth, sides[1])
    orderAmount = resolveOrderAmount(sides[1], depth, orderAmount)

    // 3
    depth = depthHandler.GetDepth(symbols[2])
    baseAmount = resolveAmount(apiClient, baseAmount, depth, sides[2])
    orderAmount = resolveOrderAmount(sides[2], depth, orderAmount)

    return baseAmount, orderAmount
}

func resolveAmount(apiClient *api.Api, baseAmount float64, depth api.WsDepth, side Side) float64 {
    if side == BUY{
        baseAmount /= depth.Asks[0].Price
    } else if side == SELL {
        baseAmount *= depth.Bids[0].Price
    }
    baseAmount -= baseAmount * apiClient.GetTakerFee()
    return baseAmount
}

func resolveOrderAmount(side Side, depth api.WsDepth, orderAmount float64) float64 {
    if side == SELL {
        if orderAmount > depth.Bids[0].Amount{
            orderAmount = depth.Bids[0].Amount
        }
        orderAmount *= depth.Bids[0].Price
    } else if side == BUY {
        t := depth.Asks[0].Price * depth.Asks[0].Amount
        if orderAmount > t {
            orderAmount = t
        }
        orderAmount /= depth.Asks[0].Price
    }
    return orderAmount
}
