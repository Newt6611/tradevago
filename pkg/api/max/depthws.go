package max

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/Newt6611/tradevago/internal"
	"github.com/Newt6611/tradevago/pkg/api"
)

type orderBookSnapshot struct {
	Channel string    `json:"c"`
	Event   string    `json:"e"`
	Market  string    `json:"M"`
	Asks    [][]string `json:"a"`
	Bids    [][]string `json:"b"`
	Timestamp int64    `json:"T"`
}

type orderBookUpdate struct {
	Channel string    `json:"c"`
	Event   string    `json:"e"`
	Market  string    `json:"M"`
	Asks    [][]string `json:"a"`
	Bids    [][]string `json:"b"`
	Timestamp int64    `json:"T"`
}

func (this *MaxWs) RunDepthConsumer(ctx context.Context, pairs []string, depth int) (chan api.WsDepth, chan struct{}) {
    subEntry := []subscriptionEntry{}
    for _, pair := range pairs {
        subEntry = append(subEntry, subscriptionEntry{
            Channel: "book",
            Market: pair,
            Depth: depth,
        })
    }
    subscriptions := subscriptionRequest {
        Action: "sub",
        Subscriptions: subEntry,
        ID: "client1",
    }

    depthDataChan := make(chan api.WsDepth, 1000)

    close := internal.RunWsClient(ctx, MAX_WS_ENDPOINT, nil, subscriptions, func(t int, b []byte, err error) {
        if err != nil {
            depthDataChan <- api.WsDepth { Err: err }
            return
        }

        var snapshot orderBookSnapshot
        err = json.Unmarshal(b, &snapshot)
        if err == nil && snapshot.Event == "snapshot" {
            wsdepth := handleDepthSnapShot(snapshot)
            wsdepth.Pair = snapshot.Market
            this.depthCache[wsdepth.Pair] = wsdepth
            depthDataChan <- wsdepth
            return
        }

        var update orderBookUpdate
        err = json.Unmarshal(b, &update)
        if err == nil && update.Event == "update" {
            wsdepth := this.depthCache[update.Market]
            wsdepth.Err = nil
            handleDepthUpdate(update, &wsdepth)
            this.depthCache[wsdepth.Pair] = wsdepth
            depthDataChan <- wsdepth
            return
        }
    })

    return depthDataChan, close
}

func handleDepthSnapShot(snapshot orderBookSnapshot) api.WsDepth {
    depth := api.WsDepth{}
    for i := range snapshot.Asks {
        askPrice, _ := strconv.ParseFloat(snapshot.Asks[i][0], 64)
        askAmount, _ := strconv.ParseFloat(snapshot.Asks[i][1], 64)

        bidPrice, _ := strconv.ParseFloat(snapshot.Bids[i][0], 64)
        bidAmount, _ := strconv.ParseFloat(snapshot.Bids[i][1], 64)

        depth.Asks = append(depth.Asks, api.DepthInfo {
            Price: askPrice,
            Amount: askAmount,
        })

        depth.Bids = append(depth.Bids, api.DepthInfo {
            Price: bidPrice,
            Amount: bidAmount,
        })
    }
    return depth
}

func handleDepthUpdate(update orderBookUpdate, depth *api.WsDepth) {
    if len(update.Asks) > 0 {
        var askPrice float64
        askAmount, _ := strconv.ParseFloat(update.Asks[0][1], 64)

        if askAmount != 0 {
            askPrice, _ = strconv.ParseFloat(update.Asks[0][0], 64)
        } else {
            askPrice, _ = strconv.ParseFloat(update.Asks[1][0], 64)
            askAmount, _ = strconv.ParseFloat(update.Asks[1][1], 64)
        }

        (*depth).Asks[0].Price = askPrice
        (*depth).Asks[0].Amount = askAmount
    }

    if len(update.Bids) > 0 {
        var bidPrice float64
        bidAmount, _ := strconv.ParseFloat(update.Bids[0][1], 64)

        if bidAmount != 0 {
            bidPrice, _ = strconv.ParseFloat(update.Bids[0][0], 64)
        } else {
            bidPrice, _ = strconv.ParseFloat(update.Bids[1][0], 64)
            bidAmount, _ = strconv.ParseFloat(update.Bids[1][1], 64)
        }

        (*depth).Bids[0].Price = bidPrice
        (*depth).Bids[0].Amount = bidAmount
    }
}
