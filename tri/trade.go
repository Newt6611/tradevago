package tri

import (
	"context"
	"errors"
	"fmt"
	"math"
	"time"

	"github.com/Newt6611/tradevago/pkg/api"
)

const (
    OrderWaitTime time.Duration = time.Minute * 1
)

var (
    OrderWaitTooLong error = errors.New("訂單等待時間過長")
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

func (this *TradeEngine) StartTrade(ctx context.Context, cycle Cycle, startAmount float64, maxStartAmount float64, currentHoldBalance float64, rate float64) {
	symbolToCheck := cycle.GetSymbolsToCheck()
	symbols := cycle.GetSymbols()
	sides := cycle.GetSides()


    var quoteAmount, baseAmount float64
    for i := 0; i < 3; i++ {
        amount := startAmount
        if i > 0 {
            switch sides[i - 1] {
            case api.BUY:
                amount = baseAmount
            case api.SELL:
                amount = quoteAmount
            }
        }
        quoteAmount, baseAmount = this.updateQuoteBaseAmount(symbols[i], amount, sides[i])

        limitBuyQuote, limitSellBase := this.getLimitBuySellQty(symbols[i])

        if quoteAmount < limitBuyQuote || baseAmount < limitSellBase {
            return
        }
    }

	startTime := time.Now()
    this.notifyHandler.SendMsg(fmt.Sprintf("%s [開始交易 %s], 初始金額 %f, 最大金額 %f, rate: %f, 目前餘額: %f",
        cycle.GetName(), symbols[0], startAmount, maxStartAmount, rate, currentHoldBalance))

    for i := 0; i < 3; i++ {
        if i > 0 { // 第一 round 不執行
            startAmount = this.waitToGetBalanceAmount(ctx, symbolToCheck[i], symbols[i], sides[i])
        }
        quoteAmount, baseAmount = this.updateQuoteBaseAmount(symbols[i], startAmount, sides[i])
trade:
        this.notifyHandler.SendMsg(fmt.Sprintf("[開始交易 %s] %d", symbols[i], i + 1))

        order, err := this.createOrder(ctx, symbols[i], sides[i], baseAmount, quoteAmount)

        if err != nil && errors.Is(err, api.ErrorBalanceNotEnougth) {
            quoteAmount *= 0.95
            baseAmount *= 0.95
            this.notifyHandler.SendMsg(fmt.Sprintf("%s resent quoteAmount: %f, baseAmount: %f", err.Error(), quoteAmount, baseAmount))
            goto trade
        } else if err != nil {
            this.notifyHandler.SendMsg(err.Error())
            return
        }

        err = this.waitToGetCertainOrderDone(order.Id, startTime)
        if err != nil {
            this.notifyHandler.SendMsg(err.Error())
            return
        }
        err = this.cancelOrder(ctx, symbols[i])
        if err != nil {
            this.notifyHandler.SendMsg(err.Error())
            return
        }

        this.notifyHandler.SendMsg(fmt.Sprintf("第 %d 單完成", i + 1))
    }

	this.notifyHandler.SendMsg(fmt.Sprintf("[%s]: 完成訂單, 完成時間 %v ", cycle.GetName(), time.Since(startTime)))
}

func (this *TradeEngine) waitToGetBalanceAmount(ctx context.Context, symbolToCheck string, pair string, side api.Side) float64 {
	pairInfo := this.tradingPairInfoHandler.Get(pair)
	var amountToCheck float64
    switch side {
    case api.BUY:
		amountToCheck = pairInfo.MinQuoteAmount
    case api.SELL:
		amountToCheck = pairInfo.MinBaseAmount
    }

	for { // 一直等到能抓到最新的錢
		accountBalance := this.balanceHandler.Get(symbolToCheck)
		if accountBalance.Balance >= amountToCheck {
			return accountBalance.Balance
		}
	}
}

func (this *TradeEngine) createOrder(ctx context.Context, pair string, side api.Side, baseAmount float64, quoteAmount float64) (api.Order, error) {
    var price float64
    if side == api.SELL {
        price = this.depthHandler.GetDepth(pair).Bids[0].Price
    } else {
        price = this.depthHandler.GetDepth(pair).Asks[0].Price
    }
    return this.apiClient.NewCreateOrderMarketService().
		WithPair(pair).
		WithSide(side).
		WithBaseAmount(baseAmount).
		WithQuoteAmount(quoteAmount).
        WithPrice(price).
		Do(ctx)
}

func (this *TradeEngine) cancelOrder(ctx context.Context, pair string) error {
    return this.apiClient.NewCancelAllOrderService().WithPair(pair).Do(ctx)
}

func (this *TradeEngine) updateQuoteBaseAmount(symbol string, amount float64, side api.Side) (float64, float64) {
    quotePrecision, basePrecision := this.getQuoteBasePrecisions(symbol)
    askPrice, bidPrice := this.getAskBidPrice(symbol)
    stepSize := this.getStepSize(symbol)

    var quoteAmount, baseAmount float64
    switch side {
    case api.BUY:
        quoteAmount = roundToDecimalPlaces(amount, quotePrecision)
        baseAmount = roundToDecimalPlaces(quoteAmount / askPrice, basePrecision)
    case api.SELL:
        baseAmount = roundToDecimalPlaces(amount, basePrecision)
        quoteAmount = roundToDecimalPlaces(baseAmount * bidPrice, quotePrecision)
    }
    if stepSize > 0 {
        // quoteAmount = applyStepSize(quoteAmount, stepSize)
        baseAmount = applyStepSize(baseAmount, stepSize)
    }
    return quoteAmount, baseAmount
}

func (this *TradeEngine) waitToGetCertainOrderDone(orderId string, startTime time.Time) error {
	for {
        if time.Since(startTime) > OrderWaitTime {
            return OrderWaitTooLong
        }
		data := this.userOrderHandler.Get(orderId)
		if data != (api.WsUserOrder{}) && (data.Status == api.OrderStatusDone || data.Status == api.OrderStatusPartial || data.Status == api.OrderStatusCompletePartial) {
            break
		}
	}
    return nil
}

func (this *TradeEngine) getQuoteBasePrecisions(symbol string) (int, int) {
	quotePrecision := this.tradingPairInfoHandler.Get(symbol).QuoteUnitPrecision
    basePrecision := this.tradingPairInfoHandler.Get(symbol).BaseUnitPrecision
    return quotePrecision, basePrecision
}

func (this *TradeEngine) getStepSize(symbol string) float64 {
    return this.tradingPairInfoHandler.Get(symbol).StepSize
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

func applyStepSize(value, stepsize float64) float64 {
	nearestMultiple := math.Floor(value / stepsize)
	result := nearestMultiple * stepsize
	return result
}
