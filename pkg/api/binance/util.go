package binance

import "github.com/Newt6611/tradevago/pkg/api"


func getOrderStatus(state string) api.OrderStatus {
    switch state {
    case "NEW":
        return api.OrderStatusWait
    case "FILLED":
        return api.OrderStatusDone
    case "CANCELED":
        return api.OrderStatusCancel
    }

    return api.OrderStatusUnknow
}
