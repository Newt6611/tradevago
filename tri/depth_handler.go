package tri

import (
	"context"
	"fmt"
	"sync"

	"github.com/Newt6611/tradevago/pkg/api"
)

type DepthHandler struct {
    wsapi       *api.WSApi
    depthTable  sync.Map
    c           chan struct{}
}

func NewDepthHandler(apiws *api.WSApi) *DepthHandler {
    return &DepthHandler {
        wsapi: apiws,
        depthTable: sync.Map{},
    }
}

func (this *DepthHandler) Handle(ctx context.Context, pairs []string, setdataf func(d *api.WsDepth, m *sync.Map)) {
    var depthDataChan chan api.WsDepth
    var c chan struct{}
ws:
    depthDataChan, c = this.wsapi.RunDepthConsumer(ctx, pairs, 1)
    for {
        depthData := <- depthDataChan
        if depthData.Err != nil {
            // TODO: Notify
            fmt.Println("[DepthHandler]: " + depthData.Err.Error())
            close(c)
            goto ws
        }

        setdataf(&depthData, &this.depthTable)
    }
}

func (this *DepthHandler) GetDepth(key string) api.WsDepth {
    value, _ := this.depthTable.Load(key)
    return value.(api.WsDepth)
}

func (this *DepthHandler) Stop() {
    if this.c != nil {
        close(this.c)
    }
}
