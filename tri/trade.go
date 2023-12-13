package tri

import (
	"context"
	"math"

	"github.com/Newt6611/tradevago/pkg/api"
)

type TradeEngine struct {
    apiClient                   *api.Api
    depthHandler                *DepthHandler
    tradingPairInfoHandler      *TradingPairInfoHandler
    notifyHandler               *NotifyHandler
}

func NewTradeEngine(apiClient *api.Api, depthHandler *DepthHandler, tradingPairInfoHandler *TradingPairInfoHandler, notifyHandler *NotifyHandler) *TradeEngine {
    return &TradeEngine {
        apiClient: apiClient,
        depthHandler: depthHandler,
        tradingPairInfoHandler: tradingPairInfoHandler,
        notifyHandler: notifyHandler,
    }
}

func (this *TradeEngine) StartTrade(ctx context.Context, cycle Cycle, startAmount float64) {
    // startTime := time.Now()
    //
    // symbolToCheck := cycle.GetSymbolsToCheck()
    // symbols := cycle.GetSymbols()
    // sides := cycle.GetSides()
    // // clientId := rand.Intn(MaxClientId)
    //
    //
    // // 須符合小數點
    // p := this.tradingPairInfoHandler.Get(symbols[0]).BaseUnitPrecision
    // startAmount = roundToDecimalPlaces(startAmount, p)
    //
    // limitBuyQuote := this.tradingPairInfoHandler.Get(symbols[0]).MinBaseAmount
    // if startAmount < limitBuyQuote {
    //     // msgChan <- fmt.Sprintf("<ERROR><%s> 初始金額 %fTWD 小於最小金額 %fTWD", cycle.GetName(), amount, limitBuyQuote)
    //     return
    // }
    //
    // // 建立第一筆訂單
    // this.notifyHandler.MsgChan <- fmt.Sprintf("```%s [開始交易 %s] 1, 最大金額 %f, 初始金額 %f, rate: %f```", cycle.GetName(), symbols[0], initMaxOrderAmount, amount, rate)
    //
    // askPrice := this.depthHandler.GetDepth(symbols[0]).Asks[0].Price
    // // order
    // _, err := this.apiClient.NewCreateOrderMarketService().
    //                     WithPair(symbols[0]).
    //                     WithSide(api.BUY).
    //                     WithPrice(askPrice).
    //                     Do(ctx)
    // if err != nil {
    //     this.notifyHandler.MsgChan <- err.Error()
    //     this.apiClient.NewCancelAllOrderService().WithPair(symbols[0]).Do(ctx)
    //     return
    // }
//----------------------------------------------------------------

//     // 抓到指定訂單
//     for _, data := range this.userOrderHandler.OrderData.Data[symbols[0]] {
//         // 檢查訂單狀態
//         if data.ID != order.OrderID { continue }
//
//         // 1: In progress (Partial deal), 2: Completed
//         if data.Status == 2 {
//             this.bitoClient.CancelAll(symbols[0])
//             goto second
//         }
//     }
//
// second:
//     // 建立第二筆訂單
//     msgChan <- fmt.Sprintf("[開始交易 %s] 2", symbols[1])
//
//     clientId = rand.Intn(MaxClientId)
//
//     // 等待 redis 檢查到錢
//     amount, amountStr = this.waitToGetAmountFromRedis(ctx, symbolToCheck[1], symbols[1], sides[1])
//
//     if sides[1] == pkg.BUY {
//         // 須符合小數點
//         p, _ = strconv.Atoi(this.tradingPairInfoHandler.GetInfo(symbols[1]).AmountPrecision)
//
//         amount = roundToDecimalPlaces(amount, p)
//         amountStr = strconv.FormatFloat(amount, 'f', p, 64)
//         order = this.bitoClient.CreateOrderMarketBuy(clientId, symbols[1], amountStr)
//
//     } else if sides[1] == pkg.SELL {
//         // 須符合小數點
//         p, _ = strconv.Atoi(this.tradingPairInfoHandler.GetInfo(symbols[1]).BasePrecision)
//
//         amount = roundToDecimalPlaces(amount, p)
//         amountStr = strconv.FormatFloat(amount, 'f', p, 64)
//
//         order = this.bitoClient.CreateOrderMarketSell(clientId, symbols[1], amountStr)
//     }
//     if len(order.StatusCode.Error) > 0 {
//         msgChan <- order.StatusCode.Error
//         this.bitoClient.CancelAll(symbols[1])
//         return
//     }
//
//     // 抓到指定訂單
//     for _, data := range this.userOrderHandler.OrderData.Data[symbols[1]] {
//         // 檢查訂單狀態
//         if data.ID != order.OrderID { continue }
//
//         // 1: In progress (Partial deal), 2: Completed
//         if data.Status == 2 {
//             this.bitoClient.CancelAll(symbols[1])
//             goto third
//         }
//     }
//
// third:
//     // 建立第三筆訂單
//     msgChan <- fmt.Sprintf("[開始交易 %s] 3", symbols[2])
//
//     clientId = rand.Intn(MaxClientId)
//
//     // 等待 redis 檢查到錢
//     amount, amountStr = this.waitToGetAmountFromRedis(ctx, symbolToCheck[2], symbols[2], sides[2])
//
//
//     if sides[2] == pkg.BUY {
//         // 須符合小數點
//         p, _ = strconv.Atoi(this.tradingPairInfoHandler.GetInfo(symbols[2]).AmountPrecision)
//
//         amount = roundToDecimalPlaces(amount, p)
//         amountStr = strconv.FormatFloat(amount, 'f', p, 64)
//         order = this.bitoClient.CreateOrderMarketBuy(clientId, symbols[2], amountStr)
//
//     } else if sides[2] == pkg.SELL {
//         // 須符合小數點
//         p, _ = strconv.Atoi(this.tradingPairInfoHandler.GetInfo(symbols[2]).BasePrecision)
//
//         amount = roundToDecimalPlaces(amount, p)
//         amountStr := strconv.FormatFloat(amount, 'f', p, 64)
//
//         order = this.bitoClient.CreateOrderMarketSell(clientId, symbols[2], amountStr)
//     }
//     if len(order.StatusCode.Error) > 0 {
//         msgChan <- order.StatusCode.Error
//         this.bitoClient.CancelAll(symbols[2])
//         return
//     }
//
//     // 抓到指定訂單
//     for _, data := range this.userOrderHandler.OrderData.Data[symbols[2]] {
//         // 檢查訂單狀態
//         if data.ID != order.OrderID { continue }
//
//         // 1: In progress (Partial deal), 2: Completed
//         if data.Status == 2 {
//             this.bitoClient.CancelAll(symbols[2])
//             goto end
//         }
//     }
//
// end:
//     msgChan <- fmt.Sprintf("```[%s]: 完成訂單, 完成時間 %v ```", cycle.GetName(), time.Since(startTime))
}

func roundToDecimalPlaces(value float64, decimalPlaces int) float64 {
	precision := math.Pow10(decimalPlaces)
	return math.Floor(value * precision) / precision
}
