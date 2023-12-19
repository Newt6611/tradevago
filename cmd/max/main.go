package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/Newt6611/tradevago/pkg/api"
	"github.com/Newt6611/tradevago/pkg/api/max"
	"github.com/Newt6611/tradevago/pkg/notify/telegram"
	maxtri "github.com/Newt6611/tradevago/tri/max"
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
    termChan := make(chan os.Signal, 1)
    signal.Notify(termChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)

    var bot *telegram.TelegramClient

    go func() {
        // Notify bot
        tgToken := viper.GetString("TELEGRAM.MAX.TOKEN")
        tgChannelId := viper.GetInt64("TELEGRAM.MAX.CHANNEL_ID")
        bot = telegram.NewTelegramClient(tgToken, tgChannelId)

        // Setup Max Api key Secret key
        apiKey := viper.GetString("MAX.API_KEY")
        apiSecret := viper.GetString("MAX.API_SECRET")
        takerFee := viper.GetFloat64("MAX.TAKER_FEE")
        makerFee := viper.GetFloat64("MAX.MAKER_FEE")

        client := max.NewMaxClient(apiKey, apiSecret, takerFee, makerFee)
        apiClient := api.NewApi(client)

        wsclient := max.NewMaxWs(apiKey, apiSecret)
        apiws := api.NewWsApi(wsclient)
        maxtri.StartMaxTri(apiClient, apiws, bot)
    }()

    for {
        <-termChan
        bot.SendCodeMsg(context.Background(), "機器人已被關閉")
        os.Exit(0)
    }
}
