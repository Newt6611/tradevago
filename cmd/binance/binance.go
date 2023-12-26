package binance

import (

	"github.com/Newt6611/tradevago/pkg/api"
	"github.com/Newt6611/tradevago/pkg/api/binance"
	"github.com/Newt6611/tradevago/pkg/notify/telegram"
	tri "github.com/Newt6611/tradevago/tri/binance"
	"github.com/spf13/viper"
)

func EntryPoint() {
	tgToken := viper.GetString("TELEGRAM.BINANCE.TOKEN")
	tgChannelId := viper.GetInt64("TELEGRAM.BINANCE.CHANNEL_ID")
	bot := telegram.NewTelegramClient(tgToken, tgChannelId)

	// Setup Max Api key Secret key
	apiKey := viper.GetString("BINANCE.API_KEY")
	apiSecret := viper.GetString("BINANCE.API_SECRET")
	takerFee := viper.GetFloat64("BINANCE.TAKER_FEE")
	makerFee := viper.GetFloat64("BINANCE.MAKER_FEE")

	client := binance.NewBinance(apiKey, apiSecret, takerFee, makerFee)
	apiClient := api.NewApi(client)

	wsclient := binance.NewBinanceWs(apiKey, apiSecret)
	apiws := api.NewWsApi(wsclient)

    tri.StartBinanceTri(apiClient, apiws, bot)
}
