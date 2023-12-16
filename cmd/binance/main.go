package main

import (
	"fmt"
	"os"

	"github.com/Newt6611/tradevago/pkg/api"
	"github.com/Newt6611/tradevago/pkg/api/binance"
	tri "github.com/Newt6611/tradevago/tri/binance"
	"github.com/spf13/viper"
)

func init() {
	workingDir, _ := os.Getwd()
	viper.SetConfigFile(workingDir + "/config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	// tgToken := viper.GetString("TELEGRAM.TOKEN")
	// tgChannelId := viper.GetInt64("TELEGRAM.MAX.CHANNEL_ID")
	// bot := telegram.NewTelegramClient(tgToken, tgChannelId)

	// Setup Max Api key Secret key
	// apiKey := viper.GetString("MAX.API_KEY")
	// apiSecret := viper.GetString("MAX.API_SECRET")
	// takerFee := viper.GetFloat64("MAX.TAKER_FEE")
	// makerFee := viper.GetFloat64("MAX.MAKER_FEE")

	client := binance.NewBinance("", "", 0, 0)
	apiClient := api.NewApi(client)

	wsclient := binance.NewBinanceWs("", "")
	apiws := api.NewWsApi(wsclient)

    tri.StartBinanceTri(apiClient, apiws, nil)
}
