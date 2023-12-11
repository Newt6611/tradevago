package api

type Side string

const (
    SELL Side = "SELL"
    BUY  Side = "BUY"
)

type OrderType string
const (
    LIMIT   OrderType = "LIMIT"
    MARKET  OrderType = "MARKET"
)
