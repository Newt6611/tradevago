package tri

import (
	"context"
	"sync"
	"time"

	"github.com/Newt6611/tradevago/pkg/api"
)

type TradingPairInfoHandler struct {
    apiClient   *api.Api
    pairInfos    map[string]api.PairInfo
    mutex       sync.Mutex
}

func NewTradingPairInfoHandler(apiClient *api.Api) *TradingPairInfoHandler {
    return &TradingPairInfoHandler{
        apiClient: apiClient,
        pairInfos: make(map[string]api.PairInfo),
    }
}

func (this *TradingPairInfoHandler) Handle(ctx context.Context, convertPairInfoName func(string)string) {
    ticker := time.NewTicker(time.Minute * 30)

    for {
        datas, err := this.apiClient.NewPairInfoService().Do(ctx)
        if err != nil {
            // TODO: notify message
            continue
        }

        for _, data := range datas {
            this.mutex.Lock()
            name := convertPairInfoName(data.Name)
            this.pairInfos[name] = data
            this.mutex.Unlock()
        }
        <-ticker.C
    }
}

func (this *TradingPairInfoHandler) Get(name string) api.PairInfo {
    var info api.PairInfo
    this.mutex.Lock()
    defer this.mutex.Unlock()
    info = this.pairInfos[name]
    return info
}

func (this *TradingPairInfoHandler) IsReady() bool {
    return len(this.pairInfos) > 0
}
