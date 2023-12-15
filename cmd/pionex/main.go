package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Newt6611/tradevago/pkg/api/pionex"
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
    ctx := context.Background()

    apiKey := viper.GetString("PIONEX.API_KEY")
    apiSecret := viper.GetString("PIONEX.API_SECRET")
    pionexWs := pionex.NewPionexWs(apiKey, apiSecret)

    d, _ := pionexWs.RunDepthConsumer(ctx, []string{ "BTC_USDT", //"ETH_USDT", "RIB_USDT",
                                                        // "TIA_USDT", "JTO_USDT", "ADA_USDT",
                                                        // "FTT_USDT", "MEME_USDT", "SOL_USDT",
                                                        }, 1)
    for {
        a := <-d
        if a.Err != nil {
            fmt.Println(a.Err)
        }
    }
    fmt.Println("Hello")
}
