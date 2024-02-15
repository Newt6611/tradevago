package tri

import (
	"context"
	"fmt"
	"time"

	"github.com/Newt6611/tradevago/pkg/api"
)

// 交易過程中可能會出現身上有多餘的幣沒有賣完, 此 function 會定期去檢查身上的幣, 若有多出來則會賣掉
func (this *TradeEngine) TradeEnd(ctx context.Context, isTrading *bool, currencyToCheck func() map[string]api.Side, convertPair func(string) string) {
	ticker := time.NewTicker(time.Second)

	for {
		for currency, side := range currencyToCheck() {
			if *isTrading {
				continue
			}
			pair := convertPair(currency)
			pairInfo := this.tradingPairInfoHandler.Get(pair)
			balance := this.balanceHandler.Get(currency).Balance

            quoteAmount, baseAmount := this.updateQuoteBaseAmount(pair, balance, side)

            if quoteAmount > pairInfo.MinQuoteAmount || baseAmount > pairInfo.MinBaseAmount {
                fmt.Println(balance)
                fmt.Printf("Symbol: %s\n", pair)
                fmt.Printf("Side: %v", side)
                fmt.Printf("MinQuoteAmount: %f\n", pairInfo.MinQuoteAmount)
                fmt.Printf("MinBaseAmount: %f\n", pairInfo.MinBaseAmount)
                fmt.Printf("quoteAmount: %f\n", quoteAmount)
                fmt.Printf("baseAmount: %f\n", baseAmount)
                fmt.Printf("==============\n")
                _, err := this.createOrder(ctx, pair, side, baseAmount, quoteAmount)
                if err != nil {
                    this.notifyHandler.SendMsg(fmt.Sprintf("賣餘幣時發生錯誤 %s", err.Error()))
                } else {
                    this.notifyHandler.SendMsg(fmt.Sprintf("完成未被賣完全的幣 %s, %f", currency, balance))
                }
            }
		}
		<-ticker.C
	}
}
