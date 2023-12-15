package pionex

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/Newt6611/tradevago/internal"
	"github.com/Newt6611/tradevago/pkg/api"
	"github.com/gorilla/websocket"
)

func (this *PionexWs) RunAccountConsumer(ctx context.Context) (chan api.WsUserAccountDatas, chan struct{}){
    userAccountChan := make(chan api.WsUserAccountDatas, 1000)

    timestamp := time.Now().UnixNano() / int64(time.Millisecond)
    requestUrl := fmt.Sprintf("?key=%s&timestamp=%d", this.apiKey, timestamp)
    signature := generateSignature(this.apiSecret, "/ws" + requestUrl + "websocket_auth")
    // wss://ws.pionex.com/ws?key=OElNn5D_Frnf5MR0ChjYdG7PunK0AOgHTvevwzWS&timestamp=1655896754515&signature=3e901247350e744353f4a7a479fd67181184a627b119352ec1b7a432925e772c
    requestUrl = fmt.Sprintf("%s%s&signature=%s", PIONEX_WS_PRIVATE, requestUrl, signature)
    fmt.Println(requestUrl)

    subscriptions := map[string]string {
        "op": "SUBSCRIBE",
        "topic": "BALANCE",
    }

    close := internal.RunWsClient(ctx, requestUrl, nil, subscriptions, func(ws *websocket.Conn, t int, b []byte, err error) {
        if err != nil {
            userAccountChan <- api.WsUserAccountDatas { Err: err }
            return
        }

        var ping pingMessage
        err = json.Unmarshal(b, &ping)
        if err == nil && ping.Operation == "PING" {
            ws.WriteJSON(pingMessage{
                Operation: "PONG",
                Timestamp: time.Now().Unix(),
            })
        }
        fmt.Println(string(b))
    })

    return userAccountChan, close
}
