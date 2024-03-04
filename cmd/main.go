package main

import (
	"fmt"
	"os"

	"github.com/Newt6611/tradevago/cmd/binance"
	"github.com/Newt6611/tradevago/cmd/bito"
	"github.com/Newt6611/tradevago/cmd/max"
	"github.com/spf13/viper"
)

const (
    Max     string = "max"
    Binance string = "binance"
    Bito    string = "bito"
)

func main() {
    viper.AutomaticEnv()
    if !viper.GetBool("DOCKER") {
        workingDir, _ := os.Getwd()
        viper.SetConfigFile(workingDir + "/config.yaml")
        if err := viper.ReadInConfig(); err != nil {
            fmt.Println(err)
            os.Exit(1)
        }
    }

    platform := viper.GetString("PLATFORM")

    switch platform {
    case Max:
        fmt.Println("start max")
        max.EntryPoint()
    case Binance:
        binance.EntryPoint()
    case Bito:
        fmt.Println("bito")
        bito.EntryPoint()
    default:
        fmt.Printf("Unable to match any platform with %s\n", platform)
    }
}
