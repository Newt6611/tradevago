package max

import (
	"github.com/Newt6611/tradevago/pkg/api"
)

const MAX_WS_ENDPOINT = "wss://max-stream.maicoin.com/ws"

type errorEvent struct {
	Event    string   `json:"e"`
	Errors   []string `json:"E"`
	ClientID string   `json:"i"`
	T        int64    `json:"T"`
}

type authenticationRequest struct {
	Action    string   `json:"action"`
	APIKey    string   `json:"apiKey"`
	Nonce     int64    `json:"nonce"`
	Signature string   `json:"signature"`
	ID        string   `json:"id"`
	Filters   []string `json:"filters"`
}

type subscriptionRequest struct {
	Action        string              `json:"action"`
	Subscriptions []subscriptionEntry `json:"subscriptions"`
	ID            string              `json:"id"`
}

type subscriptionEntry struct {
	Channel string `json:"channel"`
	Market  string `json:"market"`
	Depth   int    `json:"depth,omitempty"`
}

type MaxWs struct {
    apiKey              string
    apiSecret           string
    depthCache          map[string]api.WsDepth
}

func NewMaxWs(apiKey string, apiSecret string) *MaxWs {
    return &MaxWs {
        apiKey: apiKey,
        apiSecret: apiSecret,
        depthCache: map[string]api.WsDepth{},
    }
}
