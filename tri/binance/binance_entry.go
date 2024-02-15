package binance

import (
	"context"
	"fmt"
	"time"

	"github.com/Newt6611/tradevago/pkg/api"
	"github.com/Newt6611/tradevago/pkg/notify"
	"github.com/Newt6611/tradevago/tri"
	"github.com/Newt6611/tradevago/tri/binance/cycles"
)

// TODO: Fix NOTIONAL filter

func StartBinanceTri(api *api.Api, apiws *api.WSApi, msgBot notify.Notifier) {
	ctx := context.Background()

	//------------- init -------------//
	notifyHandler := tri.NewNotifyHandler(msgBot)
	go notifyHandler.Handle(ctx)
	defer notifyHandler.Stop()

	depthHandler := tri.NewDepthHandler(apiws, notifyHandler)
	go depthHandler.Handle(ctx, cycles.GetPairs(), 5, setupDepthData)
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

	for !depthHandler.IsReady() {}
    fmt.Println("Depth Handler Ready")

	for !balanceHandler.IsReady(cycles.BTC) {}
    fmt.Println("Balance Handlder Ready")

    for !pairInfoHandler.IsReady() {}
    fmt.Println("PairInfo Handlder Ready")

	go tradeEngine.TradeEnd(ctx, &isTrading, getAllCurrencyToCheck, getBtcQuotePair)
	notifyHandler.SendMsg("開始運作")


	cycless := cycles.GetCycles()

	for {
		tri.ClearScreen()
		for _, cycle := range cycless {
			bnbAmount := balanceHandler.Get(cycles.BNB).Balance
            if bnbAmount < 0.0064 { // ~= 58.71488TWD, Price: 9,174.2TWD
				notifyHandler.SendMsg(fmt.Sprintf("BNB幣少於 0.0064, 請趕快補充 %f", bnbAmount))
				continue
			}

			startTime := time.Now()
			rate, maxOrderAmount := tri.CycleHandler(api, depthHandler, cycle)
			if rate > 1.003 {
				currentBtcBalance := balanceHandler.Get(cycles.BTC).Balance

                // 0.0009BTC ~= 1179.7452TWD, Price: 1,311,546TWD
				if currentBtcBalance <= 0 || currentBtcBalance < 0.0009 {
					notifyHandler.SendMsg(fmt.Sprintf("BTC 餘額不足(0.0009), %f", currentBtcBalance))
					continue
				}

				initMaxOrderAmount := maxOrderAmount

				if currentBtcBalance < maxOrderAmount {
					maxOrderAmount = currentBtcBalance
				}

				// 開始交易
				isTrading = true
				tradeEngine.StartTrade(ctx, cycle, maxOrderAmount, initMaxOrderAmount, currentBtcBalance, rate)
				isTrading = false
			}
			fmt.Printf("[%s] rate: %v, maxOrderAmount: %v, %v\n", cycle.GetName(), rate, maxOrderAmount, time.Since(startTime))
		}
		<-ticker.C
	}
}
