package tri

import (
	"context"
	"math"
	"sync"

	"github.com/Newt6611/tradevago/pkg/api"
)

type DepthHandler struct {
    wsapi               *api.WSApi
    depthTable          sync.Map
    c                   chan struct{}
    notifyHandler       *NotifyHandler
    pairsCount          int
}

func NewDepthHandler(apiws *api.WSApi, notifyHandler *NotifyHandler) *DepthHandler {
    return &DepthHandler {
        wsapi: apiws,
        depthTable: sync.Map{},
        notifyHandler: notifyHandler,
        pairsCount: math.MaxInt32,
    }
}

func (this *DepthHandler) Handle(ctx context.Context, pairs []string, depth int, setdataf func(d *api.WsDepth, m *sync.Map)) {
    var depthDataChan chan api.WsDepth
    this.pairsCount = len(pairs)
ws:
    depthDataChan, this.c = this.wsapi.RunDepthConsumer(ctx, pairs, depth)
    for {
        depthData := <- depthDataChan
        if depthData.Err != nil {
            // this.notifyHandler.MsgChan <- fmt.Sprintf("[DepthHandler]: %s", depthData.Err.Error())
            close(this.c)
            goto ws
        }

        setdataf(&depthData, &this.depthTable)
    }
}

func (this *DepthHandler) IsReady() bool {
    length := 0
    this.depthTable.Range(func(key, value any) bool {
        length += 1
        return true
    })
    return length >= this.pairsCount
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
