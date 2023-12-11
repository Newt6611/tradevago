package tri

type Cycle interface {
    GetName() string
    GetSymbolsToCheck() []string // BTC, ETH
    GetSymbols() []string // BTCETH, BTCUSDT
    GetSides() []Side // SELL, BUY
}
