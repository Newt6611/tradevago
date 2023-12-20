package tri

import (
	"context"
	"fmt"
	"math"
	"time"

	"github.com/Newt6611/tradevago/pkg/api"
)

type TradeEngine struct {
	apiClient              *api.Api
	depthHandler           *DepthHandler
	tradingPairInfoHandler *TradingPairInfoHandler
	balanceHandler         *BalanceHandler
	userOrderHandler       *UserOrderHandler
	notifyHandler          *NotifyHandler
}

func NewTradeEngine(apiClient *api.Api,
	depthHandler *DepthHandler,
	tradingPairInfoHandler *TradingPairInfoHandler,
	balanceHandler *BalanceHandler,
	notifyHandler *NotifyHandler,
	userOrderHandler *UserOrderHandler) *TradeEngine {
	return &TradeEngine{
		apiClient:              apiClient,
		depthHandler:           depthHandler,
		tradingPairInfoHandler: tradingPairInfoHandler,
		balanceHandler:         balanceHandler,
		userOrderHandler:       userOrderHandler,
		notifyHandler:          notifyHandler,
	}
}

func (this *TradeEngine) StartTrade(ctx context.Context, cycle Cycle, startAmount float64, maxStartAmount float64, rate float64) {
	startTime := time.Now()

	symbolToCheck := cycle.GetSymbolsToCheck()
	symbols := cycle.GetSymbols()
	sides := cycle.GetSides()


    var quoteAmount, baseAmount float64
    for i := 0; i < 3; i++ {
        amount := startAmount
        if i > 0 {
            if sides[i - 1] == api.BUY {
                amount = baseAmount
            } else if sides[i - 1] == api.SELL {
                amount = quoteAmount
            }
        }
        quoteAmount, baseAmount = this.updateQuoteBaseAmount(symbols[i], amount, sides[i])

        limitBuyQuote, limitSellBase := this.getLimitBuySellQty(symbols[i])

        if quoteAmount < limitBuyQuote || baseAmount < limitSellBase {
            return
        }
    }

    //--
    for i := 0; i < 3; i++ {
        ask := this.depthHandler.GetDepth(symbols[i]).Asks[0]
        bid := this.depthHandler.GetDepth(symbols[i]).Bids[0]
        this.notifyHandler.SendMsg(fmt.Sprintf("%s Ask Price: %f, Amount: %f", 
                                                symbols[i], ask.Price, ask.Amount))

        this.notifyHandler.SendMsg(fmt.Sprintf("Bid Price: %f, Amount: %f", 
                                                bid.Price, bid.Amount))
    }
    //--

    this.notifyHandler.SendMsg(fmt.Sprintf("%s [開始交易 %s], 初始金額 %f, 最大金額 %f, rate: %f",
        cycle.GetName(), symbols[0], startAmount, maxStartAmount, rate))

    for i := 0; i < 3; i++ {
        //--
        ask := this.depthHandler.GetDepth(symbols[i]).Asks[0]
        bid := this.depthHandler.GetDepth(symbols[i]).Bids[0]
        this.notifyHandler.SendMsg(fmt.Sprintf("%s Ask Price: %f, Amount: %f", 
                                                symbols[i], ask.Price, ask.Amount))

        this.notifyHandler.SendMsg(fmt.Sprintf("Bid Price: %f, Amount: %f", 
                                                bid.Price, bid.Amount))
        //--

        if i > 0 { // 第一 round 不執行
            startAmount = this.waitToGetBalanceAmount(ctx, symbolToCheck[i], symbols[i], sides[i])
        }

        quoteAmount, baseAmount = this.updateQuoteBaseAmount(symbols[i], startAmount, sides[i])
        this.notifyHandler.SendMsg(fmt.Sprintf("[開始交易 %s] %d", symbols[i], i + 1))
        order, err := this.createOrder(ctx, symbols[i], sides[i], baseAmount, quoteAmount)

        if err != nil {
            this.notifyHandler.SendMsg(err.Error())
            return
        }

        this.waitToGetCertainOrderDone(order.Id)
        this.notifyHandler.SendMsg(fmt.Sprintf("第 %d 單完成", i + 1))

    }

	this.notifyHandler.SendMsg(fmt.Sprintf("[%s]: 完成訂單, 完成時間 %v ", cycle.GetName(), time.Since(startTime)))
}

func (this *TradeEngine) waitToGetBalanceAmount(ctx context.Context, symbolToCheck string, pair string, side api.Side) float64 {
	pairInfo := this.tradingPairInfoHandler.Get(pair)
	var amountToCheck float64
	if side == api.BUY {
		amountToCheck = pairInfo.MinQuoteAmount
	} else if side == api.SELL {
		amountToCheck = pairInfo.MinBaseAmount
	}

	for { // 一直等到能抓到最新的錢
		accountBalance := this.balanceHandler.Get(symbolToCheck)
		if accountBalance.Balance >= amountToCheck {
			return accountBalance.Balance
		}
	}
}

func (this *TradeEngine) createOrder(ctx context.Context, symbol string, side api.Side, baseAmount float64, quoteAmount float64) (api.Order, error) {
    return this.apiClient.NewCreateOrderMarketService().
		WithPair(symbol).
		WithSide(side).
		WithBaseAmount(baseAmount).
		WithQuoteAmount(quoteAmount).
		Do(ctx)
}

func (this *TradeEngine) updateQuoteBaseAmount(symbol string, amount float64, side api.Side) (float64, float64) {
    quotePrecision, basePrecision := this.getQuoteBasePrecisions(symbol)
    askPrice, bidPrice := this.getAskBidPrice(symbol)
    var quoteAmount, baseAmount float64
    switch side {
    case api.BUY:
        quoteAmount = roundToDecimalPlaces(amount, quotePrecision)
        baseAmount = roundToDecimalPlaces(quoteAmount / askPrice, basePrecision)
    case api.SELL:
        baseAmount = roundToDecimalPlaces(amount, basePrecision)
        quoteAmount = roundToDecimalPlaces(baseAmount * bidPrice, quotePrecision)
    }
    return quoteAmount, baseAmount
}

func (this *TradeEngine) waitToGetCertainOrderDone(orderId string) {
	for {
		data := this.userOrderHandler.Get(orderId)
		if data != (api.WsUserOrder{}) && data.Status == api.OrderStatusDone {
            break
		}
	}
}

func (this *TradeEngine) getQuoteBasePrecisions(symbol string) (int, int) {
	quotePrecision := this.tradingPairInfoHandler.Get(symbol).QuoteUnitPrecision
    basePrecision := this.tradingPairInfoHandler.Get(symbol).BaseUnitPrecision
    return quotePrecision, basePrecision
}

func (this *TradeEngine) getLimitBuySellQty(symbol string) (float64, float64) {
    limitBuyQuote := this.tradingPairInfoHandler.Get(symbol).MinQuoteAmount
    limitSellBase := this.tradingPairInfoHandler.Get(symbol).MinBaseAmount
    return limitBuyQuote, limitSellBase
}

func (this *TradeEngine) getAskBidPrice(symbol string) (float64, float64) {
    askPrice := this.depthHandler.GetDepth(symbol).Asks[0].Price
    bidPrice := this.depthHandler.GetDepth(symbol).Bids[0].Price
    return askPrice, bidPrice
}

func roundToDecimalPlaces(value float64, decimalPlaces int) float64 {
	precision := math.Pow10(decimalPlaces)
	return math.Floor(value*precision) / precision
}
