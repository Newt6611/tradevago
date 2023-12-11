package internal

import (
	"context"
	"net/http"

	"github.com/gorilla/websocket"
)

func RunWsClient(ctx context.Context, url string, header http.Header, subscription interface{}, callback func(int, []byte, error)) chan struct{} {
    close := make(chan struct{})
    go func() {
        ws:
        ws, _, err := websocket.DefaultDialer.Dial(url, header)
        if err != nil {
            callback(-1, nil, err)
            return
        }
        ws.WriteJSON(subscription)

        for {
            select {
            case <-ctx.Done():
                ws.Close()
                return
            case <- close:
                ws.Close()
                return
            default:
                t, b, err := ws.ReadMessage()
                callback(t, b, err)
                if err != nil {
                    ws.Close()
                    goto ws
                }
            }
        }

    }()

    return close
}
