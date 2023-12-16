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

func (this *TradeEngine) StartTrade(ctx context.Context, cycle Cycle, quoteAmount float64, maxQuoteAmount float64, rate float64) {
	startTime := time.Now()

	symbolToCheck := cycle.GetSymbolsToCheck()
	symbols := cycle.GetSymbols()
	sides := cycle.GetSides()

	// 須符合小數點
	quotePrecision := this.tradingPairInfoHandler.Get(symbols[0]).QuoteUnitPrecision
	quoteAmount = roundToDecimalPlaces(quoteAmount, quotePrecision)

	limitBuyQuote := this.tradingPairInfoHandler.Get(symbols[0]).MinQuoteAmount
	if quoteAmount < limitBuyQuote {
		return
	}

	// 建立第一筆訂單
	this.notifyHandler.MsgChan <- fmt.Sprintf("%s [開始交易 %s] 1, 最大金額 %f, 初始金額 %f, rate: %f",
		cycle.GetName(), symbols[0], maxQuoteAmount, quoteAmount, rate)

	askPrice := this.depthHandler.GetDepth(symbols[0]).Asks[0].Price
	baseAmount := quoteAmount / askPrice
	basePrecision := this.tradingPairInfoHandler.Get(symbols[0]).BaseUnitPrecision
	baseAmount = roundToDecimalPlaces(baseAmount, basePrecision)

	limitBuyBase := this.tradingPairInfoHandler.Get(symbols[0]).MinBaseAmount
	if baseAmount < limitBuyBase {
		// this.notifyHandler.MsgChan <- fmt.Sprintf("<ERROR><%s> 初始金額 %fBase 小於最小金額 %fBase", cycle.GetName(), baseAmount, limitBuyBase)
		return
	}

	// order
	order, err := this.apiClient.NewCreateOrderMarketService().
		WithPair(symbols[0]).
		WithSide(api.BUY).
		WithBaseAmount(baseAmount).
		WithQuoteAmount(quoteAmount).
		Do(ctx)
	if err != nil {
		this.notifyHandler.MsgChan <- err.Error()
		this.apiClient.NewCancelAllOrderService().WithPair(symbols[0]).Do(ctx)
		return
	}

	//----------------------------------------------------------------
	// 抓到指定訂單
	for {
		data := this.userOrderHandler.Get(order.Id)
		if data == (api.WsUserOrder{}) || data.Status != api.OrderStatusDone {
			continue
		}
		goto second
	}

second:
	// 建立第二筆訂單
	this.notifyHandler.MsgChan <- fmt.Sprintf("[開始交易 %s] 2", symbols[1])

	// 等待檢查到錢
	amount := this.waitToGetAmount(ctx, symbolToCheck[1], symbols[1], sides[1])

	basePrecision = this.tradingPairInfoHandler.Get(symbols[1]).BaseUnitPrecision
	quotePrecision = this.tradingPairInfoHandler.Get(symbols[1]).QuoteUnitPrecision

	if sides[1] == api.BUY {
		askPrice = this.depthHandler.GetDepth(symbols[1]).Asks[0].Price
		// 須符合小數點
		quoteAmount = roundToDecimalPlaces(amount, quotePrecision)
		baseAmount = amount / askPrice
		baseAmount = roundToDecimalPlaces(baseAmount, basePrecision)

	} else if sides[1] == api.SELL {
		bidPrice := this.depthHandler.GetDepth(symbols[1]).Bids[0].Price
		// 須符合小數點
		baseAmount = roundToDecimalPlaces(amount, basePrecision)
		quoteAmount = roundToDecimalPlaces(amount * bidPrice, quotePrecision)
	}

	order, err = this.apiClient.NewCreateOrderMarketService().
		WithPair(symbols[1]).
		WithSide(sides[1]).
		WithBaseAmount(baseAmount).
		WithQuoteAmount(quoteAmount).
		Do(ctx)
	if err != nil {
		this.notifyHandler.MsgChan <- err.Error()
        this.notifyHandler.MsgChan <- fmt.Sprintf("Amount: %f, BaseAmount: %f, QuoteAmount: %f, Ask price: %f", amount, baseAmount, quoteAmount, askPrice)
		this.apiClient.NewCancelAllOrderService().WithPair(symbols[1]).Do(ctx)
		return
	}

	// 抓到指定訂單
	for {
		data := this.userOrderHandler.Get(order.Id)
		if data == (api.WsUserOrder{}) || data.Status != api.OrderStatusDone {
			continue
		}
		goto third
	}

third:
	// 建立第三筆訂單
	this.notifyHandler.MsgChan <- fmt.Sprintf("[開始交易 %s] 3", symbols[2])

	// 等待檢查到錢
	amount = this.waitToGetAmount(ctx, symbolToCheck[2], symbols[2], sides[2])

	basePrecision = this.tradingPairInfoHandler.Get(symbols[2]).BaseUnitPrecision
	quotePrecision = this.tradingPairInfoHandler.Get(symbols[2]).QuoteUnitPrecision

	if sides[2] == api.BUY {
		askPrice = this.depthHandler.GetDepth(symbols[2]).Asks[0].Price
		// 須符合小數點
		quoteAmount = roundToDecimalPlaces(amount, quotePrecision)
		baseAmount = amount / askPrice
		baseAmount = roundToDecimalPlaces(baseAmount, basePrecision)

	} else if sides[2] == api.SELL {
		bidPrice := this.depthHandler.GetDepth(symbols[2]).Bids[0].Price
		// 須符合小數點
		baseAmount = roundToDecimalPlaces(amount, basePrecision)
		quoteAmount = roundToDecimalPlaces(amount * bidPrice, quotePrecision)
	}

	order, err = this.apiClient.NewCreateOrderMarketService().
		WithPair(symbols[2]).
		WithSide(sides[2]).
		WithBaseAmount(baseAmount).
		WithQuoteAmount(quoteAmount).
		Do(ctx)
	if err != nil {
		this.notifyHandler.MsgChan <- err.Error()
		this.apiClient.NewCancelAllOrderService().WithPair(symbols[2]).Do(ctx)
		return
	}

	// 抓到指定訂單
	for {
		data := this.userOrderHandler.Get(order.Id)
		if data == (api.WsUserOrder{}) || data.Status != api.OrderStatusDone {
			continue
		}
		goto end
	}

end:
	this.notifyHandler.MsgChan <- fmt.Sprintf("[%s]: 完成訂單, 完成時間 %v ", cycle.GetName(), time.Since(startTime))
}

// 交易過程中可能會出現身上有多餘的幣沒有賣完, 此 function 會定期去檢查身上的幣, 若有多出來則會賣掉
func (this *TradeEngine) TradeEnd(ctx context.Context, isTrading *bool, currencyToCheck func() []string, convertPair func(string) string) {
	ticker := time.NewTicker(time.Second)

	for {
		for _, c := range currencyToCheck() {
			if *isTrading {
				continue
			}
			pair := convertPair(c)
			pairInfo := this.tradingPairInfoHandler.Get(pair)
			balance := this.balanceHandler.Get(c).Balance

			if balance > pairInfo.MinBaseAmount {
				price := this.depthHandler.GetDepth(pair).Bids[0].Price
				baseAmount := roundToDecimalPlaces(balance, pairInfo.BaseUnitPrecision)
				quoteAmount := roundToDecimalPlaces(baseAmount*price, pairInfo.QuoteUnitPrecision)
				this.apiClient.NewCreateOrderMarketService().
					WithPair(pair).
					WithSide(api.SELL).
					WithBaseAmount(baseAmount).
					WithQuoteAmount(quoteAmount).
					Do(ctx)

				this.notifyHandler.MsgChan <- fmt.Sprintf("完成未被賣完全的幣 %s, %f", c, balance)
			}
		}
		<-ticker.C
	}
}

func (this *TradeEngine) waitToGetAmount(ctx context.Context, symbolToCheck string, pair string, side api.Side) float64 {
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

func roundToDecimalPlaces(value float64, decimalPlaces int) float64 {
	precision := math.Pow10(decimalPlaces)
	return math.Floor(value*precision) / precision
}
