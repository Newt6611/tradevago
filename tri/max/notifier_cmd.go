package max

import (
	"fmt"

	"github.com/Newt6611/tradevago/pkg/notify"
	"github.com/Newt6611/tradevago/tri"
	"github.com/Newt6611/tradevago/tri/max/cycles"
)

func notifierCmds(balanceHandler *tri.BalanceHandler, depthHandler *tri.DepthHandler) map[string]func() string {
	return map[string]func() string{
		"!TWD": func() string {
			twd := balanceHandler.Get(cycles.TWD).Balance
			return fmt.Sprintf("%.8f", twd)
		},
		"!MAX": func() string {
			max := balanceHandler.Get(cycles.MAX).Balance
			return fmt.Sprintf("%.8f", max)
		},
		notify.Sticker: func() string {
            maxPrice := depthHandler.GetDepth(cycles.MAXTWD).Bids[0].Price
			twd := balanceHandler.Get(cycles.TWD).Balance
			max := balanceHandler.Get(cycles.MAX).Balance
			return fmt.Sprintf("TWD: %.8f\nMAX: %.8f\n(%.8f TWD)", twd, max, max * maxPrice)
		},
	}
}
