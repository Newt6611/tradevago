package pionex

import (
	"context"
	"fmt"
	"strconv"

	"github.com/Newt6611/tradevago/internal"
	"github.com/Newt6611/tradevago/pkg/api"
	"github.com/gorilla/websocket"
)

type depthWsRequest struct {
    Operation string `json:"op"`
    Topic     string `json:"topic"`
    Symbol    string `json:"symbol"`
    Limit     int    `json:"limit"`
}

type depthDataInfo struct {
	Bids [][]string `json:"bids"`
	Asks [][]string `json:"asks"`
}

type depthDatas struct {
	Topic     string        `json:"topic"`
	Symbol    string        `json:"symbol"`
	Data      depthDataInfo `json:"data"`
	Timestamp int64         `json:"timestamp"`
}

func (this *PionexWs) RunDepthConsumer(ctx context.Context, pairs []string, depth int) (chan api.WsDepth, chan struct{}) {
    this.depthCloseChan = make(chan struct{})
    depthDataChan := make(chan api.WsDepth, 300)

    for _, pair := range pairs {
        depthWsRequest := depthWsRequest {
            Operation: "SUBSCRIBE",
            Topic: "DEPTH",
            Symbol: pair,
            Limit: depth,
        }
        c := internal.RunWsClient(ctx, PIONEX_WS_PUBLIC, nil, depthWsRequest, func(ws *websocket.Conn, t int, b []byte, err error) {
            if err != nil {
                depthDataChan <- api.WsDepth { Err: err }
                return
            }

            // var ping pingMessage
            // err = json.Unmarshal(b, &ping)
            // if err == nil && ping.Operation == "PING" {
            //     ws.WriteJSON(pingMessage{
            //         Operation: "PONG",
            //         Timestamp: time.Now().Unix(),
            //     })
            //     return
            // }
            fmt.Println(string(b))

            // var depthDatas depthDatas
            // err = json.Unmarshal(b, &depthDatas)
            // if err == nil && depthDatas.Timestamp > 0 {
            //     depthDataChan <- handleDepthResponse(depthDatas)
            //     return
            // }
        })
        this.depthCloseChans = append(this.depthCloseChans, c)
    }

    return depthDataChan, this.depthCloseChan
}

func (this *PionexWs) handleClose() {
    for {
        _, ok := <-this.depthCloseChan
        if !ok {
            for _, c := range this.depthCloseChans {
                if c != nil {
                    close(c)
                }
            }
        }
    }
}

func handleDepthResponse(depthDatas depthDatas) api.WsDepth {
    askPrice, _ := strconv.ParseFloat(depthDatas.Data.Asks[0][0], 64)
    askAmount, _ := strconv.ParseFloat(depthDatas.Data.Asks[0][1], 64)
    bidPrice, _ := strconv.ParseFloat(depthDatas.Data.Bids[0][0], 64)
    bidAmount, _ := strconv.ParseFloat(depthDatas.Data.Bids[0][1], 64)

    return api.WsDepth {
        Pair: depthDatas.Symbol,
        Asks: []api.DepthInfo {
            {
                Price: askPrice,
                Amount: askAmount,
            },
        },
        Bids: []api.DepthInfo {
            {
                Price: bidPrice,
                Amount: bidAmount,
            },
        },
    }
}
