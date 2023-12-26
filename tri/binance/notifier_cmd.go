package binance

import (
	"fmt"

	"github.com/Newt6611/tradevago/pkg/notify"
	"github.com/Newt6611/tradevago/tri"
	"github.com/Newt6611/tradevago/tri/max/cycles"
)

func notifierCmds(balanceHandler *tri.BalanceHandler) map[string]func() string {
	return map[string]func() string{
		"!BTC": func() string {
			btc := balanceHandler.Get(cycles.BTC).Balance
			return fmt.Sprintf("%.8f", btc)
		},
		notify.Sticker: func() string {
			btc := balanceHandler.Get(cycles.TWD).Balance
			return fmt.Sprintf("BTC: %.8f", btc)
		},
	}
}
