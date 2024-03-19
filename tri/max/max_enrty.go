package max

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/Newt6611/tradevago/pkg/api"
	"github.com/Newt6611/tradevago/tri"
	"github.com/Newt6611/tradevago/tri/max/cycles"
)

const (
    PairStartTradeTimes int = 10
)

func StartMaxTri(backend *tri.Backend) {
	ctx := context.Background()

	//------------- init -------------//
	notifyHandler := tri.NewNotifyHandler(backend.MsgBot)
	go notifyHandler.Handle(ctx)
	defer notifyHandler.Stop()

	depthHandler := tri.NewDepthHandler(backend.Apiws, notifyHandler)
	go depthHandler.Handle(ctx, cycles.GetPairs(), 1, setupDepthData)
	defer depthHandler.Stop()

	time.Sleep(time.Millisecond * 100)

	balanceHandler := tri.NewBalanceHandler(backend.Apiws, notifyHandler)
	go balanceHandler.Handle(ctx, setBalanceData)
	defer balanceHandler.Stop()

	time.Sleep(time.Millisecond * 100)

	userOrderhandler := tri.NewUserOrderHandler(backend.Apiws, notifyHandler)
	go userOrderhandler.Handle(ctx, setUserOrderData)
	defer userOrderhandler.Stop()

	time.Sleep(time.Millisecond * 100)

	pairInfoHandler := tri.NewTradingPairInfoHandler(backend.Api)
	go pairInfoHandler.Handle(ctx, convertPairName)

	tradeEngine := tri.NewTradeEngine(backend.Api, depthHandler, pairInfoHandler, balanceHandler, notifyHandler, userOrderhandler)
	isTrading := false
	go notifyHandler.HandleMessage(notifierCmds(balanceHandler, depthHandler))

    tradeSignal := tri.NewTradeSignalHandler()
    go tradeSignal.Clear(time.Second * 10)
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
			if maxAmount < 5 {
				// notifyHandler.SendMsg(fmt.Sprintf("MAX幣少於 5, 請趕快補充 %f", maxAmount))
				notifyHandler.SendMsg(fmt.Sprintf("開始購買 Max 幣"))
                err := buyMaxFee(ctx, backend.Api, userOrderhandler, notifyHandler)
                if err != nil {
                    notifyHandler.SendMsg(err.Error())
                    continue
                }
                notifyHandler.SendMsg(fmt.Sprintf("購買完畢"))
			}

			startTime := time.Now()
			rate, maxOrderAmount := tri.CycleHandler(backend.Api, depthHandler, cycle)
			if rate > 1.001 {
                _, do := tradeSignal.StartTradeOrNot(cycle.GetName());
                if !do {
                    continue
                }

				currentTwdBalance := balanceHandler.Get(cycles.TWD).Balance

				if currentTwdBalance <= 0 || currentTwdBalance < 200 {
					notifyHandler.SendMsg(fmt.Sprintf("TWD 餘額不足(200), %f", currentTwdBalance))
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

// 一次購買 50 MAX
func buyMaxFee(ctx context.Context, apii *api.Api, userOrderhandler *tri.UserOrderHandler, notifyHandler *tri.NotifyHandler) error {
    order, err := apii.NewCreateOrderMarketService().
		WithPair(cycles.MAXTWD).
		WithSide(api.BUY).
		WithBaseAmount(50).
		Do(ctx)
    if err != nil {
        return err
    }

    startTime := time.Now()
    for {
        if time.Since(startTime) > time.Minute * 1 {
            return errors.New("購買 Max 時間等待過長")
        }
        o := userOrderhandler.Get(order.Id)
        if o.Status == api.OrderStatusDone || 
            o.Status == api.OrderStatusPartial || 
            o.Status == api.OrderStatusCompletePartial {
            break
        }
    }
    return nil
}
