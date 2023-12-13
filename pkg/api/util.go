package api

type Side string

const (
    SELL Side = "SELL"
    BUY  Side = "BUY"
)

type OrderType string
const (
    OrderTypeLimit   OrderType = "LIMIT"
    OrderTypeMarket  OrderType = "MARKET"
)
