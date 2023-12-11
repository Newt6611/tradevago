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
    notiftHandler := tri.NewNotifyHandler(msgBot)
    go notiftHandler.Handle(ctx)
    defer notiftHandler.Stop()

    depthHandler := tri.NewDepthHandler(apiws)
    go depthHandler.Handle(ctx, cycles.GetPairs(), setupDepthData)
    defer depthHandler.Stop()

    pairInfoHandler := tri.NewTradingPairInfoHandler(api)
    go pairInfoHandler.Handle(ctx, convertPairName)
    //-------------------------------//

    time.Sleep(time.Second * 1)
    notiftHandler.MsgChan <- "開始運作"

    ticker := time.NewTicker(time.Millisecond * 500)
    cycles := cycles.GetCycles()

    for {
        tri.ClearScreen()
        for _, cycle := range cycles {
            startTime := time.Now()
            rate, maxOrderAmount := tri.CycleHandler(api, depthHandler, cycle)
            if rate > 1.015 {
                msg := fmt.Sprintf("[%s] rate: %v, maxOrderAmount: %v", cycle.GetName(), rate, maxOrderAmount)
                notiftHandler.MsgChan <- msg

            }

            fmt.Printf("[%s] rate: %v, maxOrderAmount: %v, %v\n", cycle.GetName(), rate, maxOrderAmount, time.Since(startTime))
        }

        <-ticker.C
    }
}
