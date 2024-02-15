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
    OrderWaitTime time.Duration = time.Minute * 5
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

        quoteAmount, baseAmount, _ := this.updateQuoteBaseAmount(symbols[i], amount, sides[i])
        // if !trade {
        //     return
        // }

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
        quoteAmount, baseAmount, _ := this.updateQuoteBaseAmount(symbols[i], startAmount, sides[i])
        // if !trade {
        //     return
        // }
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

        processToLong := this.waitToGetCertainOrderDone(order.Id, startTime)
        if processToLong != nil {
            this.notifyHandler.SendMsg(processToLong.Error())
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

func (this *TradeEngine) createOrder(ctx context.Context, symbol string, side api.Side, baseAmount float64, quoteAmount float64) (api.Order, error) {
    return this.apiClient.NewCreateOrderMarketService().
		WithPair(symbol).
		WithSide(side).
		WithBaseAmount(baseAmount).
		WithQuoteAmount(quoteAmount).
		Do(ctx)
}

func (this *TradeEngine) updateQuoteBaseAmount(symbol string, amount float64, side api.Side) (float64, float64, bool) {
    quotePrecision, basePrecision := this.getQuoteBasePrecisions(symbol)
    askPrice, bidPrice := this.getAskBidPrice(symbol)
    trade := true

    var quoteAmount, baseAmount float64
    if side == api.BUY {
        quoteAmount = roundToDecimalPlaces(amount, quotePrecision)
        baseAmount = roundToDecimalPlaces(quoteAmount / askPrice, basePrecision)
    } else {
        baseAmount = roundToDecimalPlaces(amount, basePrecision)
        quoteAmount = roundToDecimalPlaces(baseAmount * bidPrice, quotePrecision)
    }

    // apply binance filter
    isBinance := this.tradingPairInfoHandler.Get(symbol).Binance
    if isBinance {
        quoteAmount, baseAmount, trade = this.applyBinanceFilter(symbol, side, quoteAmount, baseAmount)
    }
    return quoteAmount, baseAmount, trade
}

func (this *TradeEngine) waitToGetCertainOrderDone(orderId string, startTime time.Time) error {
	for {
        if time.Since(startTime) > OrderWaitTime {
            return errors.New("訂單等待時間過長")
        }
		data := this.userOrderHandler.Get(orderId)
		if data != (api.WsUserOrder{}) && data.Status == api.OrderStatusDone {
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

func (this *TradeEngine) applyBinanceFilter(symbol string, side api.Side, quoteAmount, baseAmount float64) (float64, float64, bool) {
    quoteAmount, baseAmount = applyStepSizeFilter(this, symbol, quoteAmount, baseAmount)
    trade := applyNotionalFilter(this, symbol, side, quoteAmount, baseAmount)
    return quoteAmount, baseAmount, trade
}

func roundToDecimalPlaces(value float64, decimalPlaces int) float64 {
	precision := math.Pow10(decimalPlaces)
	return math.Floor(value*precision) / precision
}
