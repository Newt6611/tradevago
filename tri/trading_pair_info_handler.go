package tri

import (
	"context"
	"fmt"
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
            fmt.Println(err)
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
    info = this.pairInfos[name]
    this.mutex.Unlock()
    return info
}
