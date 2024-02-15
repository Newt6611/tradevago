package tri

import (
	"fmt"
	"math"

	"github.com/Newt6611/tradevago/pkg/api"
)


func applyStepSizeFilter(tradeEngine *TradeEngine, symbol string, quoteAmount, baseAmount float64) (float64, float64) {
    stepSize := tradeEngine.tradingPairInfoHandler.Get(symbol).StepSize

    quoteAmountMultiple := math.Floor(quoteAmount / stepSize)
    baseAmountMultiple := math.Floor(baseAmount / stepSize)

    quoteAmount = quoteAmountMultiple * stepSize
    baseAmount = baseAmountMultiple * stepSize

	return quoteAmount, baseAmount
}

func applyNotionalFilter(tradeEngine *TradeEngine, symbol string, side api.Side, quoteAmount float64, baseAmount float64) bool {
    pairInfo := tradeEngine.tradingPairInfoHandler.Get(symbol)
    depth := tradeEngine.depthHandler.GetDepth(symbol)
    var price float64
    var amount float64

    if side == api.BUY {
        price = depth.Asks[0].Price
        amount = price * quoteAmount

    } else if side == api.SELL {
        price = depth.Bids[0].Price
        amount = price * baseAmount
    }

    if pairInfo.ApplyMaxToMarket && amount > pairInfo.MaxNotional {
        return false
    }

    if pairInfo.ApplyMinToMarket && amount < pairInfo.MinNotional {
        return false
    }

    fmt.Println("----------------")
    fmt.Println(symbol)
    fmt.Printf("Price: %.8f\n", price)
    fmt.Printf("Quote Amount: %.8f\n", quoteAmount)
    fmt.Printf("Base Amount: %.8f\n", baseAmount)
    fmt.Printf("Amount: %.8f\n", amount)
    fmt.Printf("MaxNotinoal %.8f\n", pairInfo.MaxNotional)
    fmt.Printf("MinNotional %.8f\n", pairInfo.MinNotional)
    fmt.Println("----------------")
    return true
}
