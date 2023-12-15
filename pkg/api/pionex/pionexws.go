package pionex

const (
    PIONEX_WS_PRIVATE   string = "wss://ws.pionex.com/ws"
    PIONEX_WS_PUBLIC    string = "wss://ws.pionex.com/wsPub"
)

type PionexWs struct {
    apiKey              string
    apiSecret           string
    depthCloseChan      chan struct{}
    depthCloseChans     []chan struct{}
}

func NewPionexWs(apiKey string, apiSecret string) *PionexWs {
    return &PionexWs {
        apiKey: apiKey,
        apiSecret: apiSecret,

        depthCloseChans: []chan struct{}{},
    }
}

type pingMessage struct {
	Operation  string `json:"op"`
	Timestamp  int64  `json:"timestamp"`
}
