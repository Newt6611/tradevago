package tri

import (
	"context"
	"fmt"
	"time"

	"github.com/Newt6611/tradevago/pkg/api"
)

// 交易過程中可能會出現身上有多餘的幣沒有賣完, 此 function 會定期去檢查身上的幣, 若有多出來則會賣掉
func (this *TradeEngine) TradeEnd(ctx context.Context, isTrading *bool, currencyToCheck func() []string, convertPair func(string) string) {
	ticker := time.NewTicker(time.Second)

	for {
		for _, c := range currencyToCheck() {
			if *isTrading {
				continue
			}
			pair := convertPair(c)
			pairInfo := this.tradingPairInfoHandler.Get(pair)
			balance := this.balanceHandler.Get(c).Balance

			if balance > pairInfo.MinBaseAmount {
				price := this.depthHandler.GetDepth(pair).Bids[0].Price
				baseAmount := roundToDecimalPlaces(balance, pairInfo.BaseUnitPrecision)
				quoteAmount := roundToDecimalPlaces(baseAmount*price, pairInfo.QuoteUnitPrecision)
				this.apiClient.NewCreateOrderMarketService().
					WithPair(pair).
					WithSide(api.SELL).
					WithBaseAmount(baseAmount).
					WithQuoteAmount(quoteAmount).
					Do(ctx)

				this.notifyHandler.SendMsg(fmt.Sprintf("完成未被賣完全的幣 %s, %f", c, balance))
			}
		}
		<-ticker.C
	}
}
