package max

import (
	"fmt"

	"github.com/Newt6611/tradevago/tri"
	"github.com/Newt6611/tradevago/tri/max/cycles"
)

func notifierCmds(balanceHandler *tri.BalanceHandler) map[string]func() string {
	return map[string]func() string{
		"!TWD": func() string {
			twd := balanceHandler.Get(cycles.TWD).Balance
			return fmt.Sprintf("%.8f", twd)
		},
		"!MAX": func() string {
			max := balanceHandler.Get(cycles.MAX).Balance
			return fmt.Sprintf("%.8f", max)
		},
	}
}
