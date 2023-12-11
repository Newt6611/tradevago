package max

import "github.com/Newt6611/tradevago/pkg/api"

const MAX_WS_ENDPOINT = "wss://max-stream.maicoin.com/ws"

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
    depthCache map[string]api.WsDepth
}

func NewMaxWs() *MaxWs {
    return &MaxWs {
        depthCache: map[string]api.WsDepth{},
    }
}
