package max

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/Newt6611/tradevago/internal"
	"github.com/Newt6611/tradevago/pkg/api"
	"github.com/gorilla/websocket"
)

type accountEvent struct {
	Channel         string       `json:"c"`
	Event           string       `json:"e"`
	Balances        []balance    `json:"B"`
	Timestamp       int64        `json:"T"`
}

type balance struct {
	Currency   string  `json:"cu"`
	Available  float64 `json:"av,string"`
	Locked     float64 `json:"l,string"`
	Staked     *float64 `json:"stk,string"`
	UpdateTime int     `json:"TU"`
}

func (this *MaxWs) RunAccountConsumer(ctx context.Context) (chan api.WsUserAccountDatas, chan struct{}){
    userAccountChan := make(chan api.WsUserAccountDatas, 500)

    nonce := generateNonce()
    subscriptions := authenticationRequest {
        Action: "auth",
        APIKey: this.apiKey,
        Nonce: nonce,
        Signature: generateSignature(this.apiSecret, fmt.Sprintf("%d", nonce)),
        ID: "clientAccountId",
        Filters: []string { "account" },
    }

    close := internal.RunWsClient(ctx, MAX_WS_ENDPOINT, nil, subscriptions, func(ws *websocket.Conn, t int, b []byte, err error) {
        if err != nil {
            userAccountChan <- api.WsUserAccountDatas { Err: err }
            return
        }

        var errEvent errorEvent
        err = json.Unmarshal(b, &errEvent)
        if err == nil && errEvent.Event == "error" {
            userAccountChan <- api.WsUserAccountDatas { Err: errors.New(errEvent.Errors[0]) }
            return
        }

        var accountEvent accountEvent
        err = json.Unmarshal(b, &accountEvent)
        if err == nil {
            accountBalances := handleUserAccountEvent(accountEvent.Balances)
            userAccountChan <- api.WsUserAccountDatas { Datas: accountBalances, Err: nil }
            return
        }
    })

    return userAccountChan, close
}

func handleUserAccountEvent(balances []balance) []api.WsUserAccountBalance {
    accountBalances := []api.WsUserAccountBalance{}

    for _, balance := range balances {
        accountBalances = append(accountBalances, api.WsUserAccountBalance {
            Currency: balance.Currency,
            Balance: balance.Available,
        })
    }

    return accountBalances
}
