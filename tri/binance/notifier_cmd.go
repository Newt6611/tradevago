package binance

import (
	"fmt"

	"github.com/Newt6611/tradevago/pkg/notify"
	"github.com/Newt6611/tradevago/tri"
	"github.com/Newt6611/tradevago/tri/binance/cycles"
)

func notifierCmds(balanceHandler *tri.BalanceHandler) map[string]func() string {
	return map[string]func() string{
		"!BTC": func() string {
			btc := balanceHandler.Get(cycles.BTC).Balance
			return fmt.Sprintf("%.8f", btc)
		},
		"!BNB": func() string {
			bnb := balanceHandler.Get(cycles.BNB).Balance
			return fmt.Sprintf("%.8f", bnb)
		},
		notify.Sticker: func() string {
			btc := balanceHandler.Get(cycles.BTC).Balance
			bnb := balanceHandler.Get(cycles.BNB).Balance
            return fmt.Sprintf("BTC: %.8f\nBNB: %.8f", btc, bnb)
		},
	}
}
