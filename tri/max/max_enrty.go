package max

import (
	"context"
	"fmt"
	"time"

	"github.com/Newt6611/tradevago/pkg/api"
	"github.com/Newt6611/tradevago/pkg/notify"
	"github.com/Newt6611/tradevago/tri"
	"github.com/Newt6611/tradevago/tri/max/cycles"
)

func StartMaxTri(api *api.Api, apiws *api.WSApi, msgBot notify.Notifier) {
	ctx := context.Background()

	//------------- init -------------//
	notifyHandler := tri.NewNotifyHandler(msgBot)
	go notifyHandler.Handle(ctx)
	defer notifyHandler.Stop()

	depthHandler := tri.NewDepthHandler(apiws, notifyHandler)
	go depthHandler.Handle(ctx, cycles.GetPairs(), 1, setupDepthData)
	defer depthHandler.Stop()

	time.Sleep(time.Millisecond * 100)

	balanceHandler := tri.NewBalanceHandler(apiws, notifyHandler)
	go balanceHandler.Handle(ctx, setBalanceData)
	defer balanceHandler.Stop()

	time.Sleep(time.Millisecond * 100)

	userOrderhandler := tri.NewUserOrderHandler(apiws, notifyHandler)
	go userOrderhandler.Handle(ctx, setUserOrderData)
	defer userOrderhandler.Stop()

	time.Sleep(time.Millisecond * 100)

	pairInfoHandler := tri.NewTradingPairInfoHandler(api)
	go pairInfoHandler.Handle(ctx, convertPairName)

	tradeEngine := tri.NewTradeEngine(api, depthHandler, pairInfoHandler, balanceHandler, notifyHandler, userOrderhandler)
	isTrading := false
	go notifyHandler.HandleMessage(notifierCmds(balanceHandler))
	//-------------------------------//

    go userOrderhandler.DeleteCompletedOrder(&isTrading)

	ticker := time.NewTicker(time.Millisecond * 500)
	cycless := cycles.GetCycles()

	for !depthHandler.IsReady() {}
    fmt.Println("Depth Handler Ready")

	for !balanceHandler.IsReady(cycles.TWD) {}
    fmt.Println("Balance Handlder Ready")

    for !pairInfoHandler.IsReady() {}
    fmt.Println("PairInfo Handlder Ready")

	go tradeEngine.TradeEnd(ctx, &isTrading, getAllCurrencyToCheck, getTwdQuotePair)
	notifyHandler.SendMsg("開始運作")

	for {
		tri.ClearScreen()
		for _, cycle := range cycless {
			maxAmount := balanceHandler.Get(cycles.MAX).Balance
			if maxAmount < 10 {
				notifyHandler.SendMsg(fmt.Sprintf("MAX幣少於 10, 請趕快補充 %f", maxAmount))
				continue
			}

			startTime := time.Now()
			rate, maxOrderAmount := tri.CycleHandler(api, depthHandler, cycle)
			if rate > 1.0001 {
				currentTwdBalance := balanceHandler.Get(cycles.TWD).Balance

				if currentTwdBalance <= 0 || currentTwdBalance < 800 {
					notifyHandler.SendMsg(fmt.Sprintf("TWD 餘額不足(800), %f", currentTwdBalance))
					continue
				}

				initMaxOrderAmount := maxOrderAmount

				if currentTwdBalance < maxOrderAmount {
					maxOrderAmount = currentTwdBalance
				}

				// 開始交易
				isTrading = true
				tradeEngine.StartTrade(ctx, cycle, maxOrderAmount, initMaxOrderAmount, currentTwdBalance, rate)
				isTrading = false
			}
			fmt.Printf("[%s] rate: %v, maxOrderAmount: %v, %v\n", cycle.GetName(), rate, maxOrderAmount, time.Since(startTime))
		}
		<-ticker.C
	}
}
