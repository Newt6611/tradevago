package tri

import "github.com/Newt6611/tradevago/pkg/api"

type Cycle interface {
    GetName() string
    GetSymbolsToCheck() []string // BTC, ETH
    GetSymbols() []string // BTCETH, BTCUSDT
    GetSides() []api.Side // SELL, BUY
}
