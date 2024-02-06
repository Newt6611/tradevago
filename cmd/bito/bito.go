package bito

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/Newt6611/tradevago/pkg/api"
	b "github.com/Newt6611/tradevago/pkg/api/bito"
	"github.com/Newt6611/tradevago/pkg/notify/telegram"
	bitotri "github.com/Newt6611/tradevago/tri/bito"
	"github.com/spf13/viper"
)

func EntryPoint() {
    termChan := make(chan os.Signal, 1)
    signal.Notify(termChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)

    var bot *telegram.TelegramClient

    go func() {
        // Notify bot
        tgToken := viper.GetString("TELEGRAM.BITO.TOKEN")
        tgChannelId := viper.GetInt64("TELEGRAM.BITO.CHANNEL_ID")
        bot = telegram.NewTelegramClient(tgToken, tgChannelId)

        // Setup Bito Api key Secret key
        apiKey := viper.GetString("BITO.API_KEY")
        apiSecret := viper.GetString("BITO.API_SECRET")
        email := viper.GetString("BITO.EMAIL")
        takerFee := viper.GetFloat64("BITO.TAKER_FEE")
        makerFee := viper.GetFloat64("BITO.MAKER_FEE")

        client := b.NewBitoClient(apiKey, apiSecret, takerFee, makerFee, email)
        apiClient := api.NewApi(client)
        wsclient := b.NewBitoWs(apiKey, apiSecret, email)
        apiws := api.NewWsApi(wsclient)
        bitotri.StartBitoTri(apiClient, apiws, bot)
    }()

    for {
        <-termChan
        bot.SendCodeMsg(context.Background(), "機器人已被關閉")
        os.Exit(0)
    }
}
