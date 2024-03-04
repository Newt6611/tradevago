package bito

import (
	"fmt"

	"github.com/Newt6611/tradevago/pkg/notify"
	"github.com/Newt6611/tradevago/tri"
	"github.com/Newt6611/tradevago/tri/bito/cycles"
)

func notifierCmds(balanceHandler *tri.BalanceHandler, depthHandler *tri.DepthHandler) map[string]func() string {
	return map[string]func() string{
		"!TWD": func() string {
			twd := balanceHandler.Get(cycles.TWD).Balance
			return fmt.Sprintf("%.8f", twd)
		},
		"!BITO": func() string {
			max := balanceHandler.Get(cycles.BITO).Balance
			return fmt.Sprintf("%.8f", max)
		},
		notify.Sticker: func() string {
			twd := balanceHandler.Get(cycles.TWD).Balance
			bito := balanceHandler.Get(cycles.BITO).Balance
            bitoPrice := depthHandler.GetDepth(cycles.BITOTWD).Bids[0].Price
			return fmt.Sprintf("TWD: %.8f\nBITO: %.8f\n(%.8f TWD)", twd, bito, bito * bitoPrice)
		},
	}
}
