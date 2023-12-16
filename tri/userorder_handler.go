package tri

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/Newt6611/tradevago/pkg/api"
)

type UserOrderHandler struct {
    wsapi           *api.WSApi
    c               chan struct{}
    notifyHandler   *NotifyHandler
    userOrders      sync.Map
}

func NewUserOrderHandler(apiws *api.WSApi, notifyHandler *NotifyHandler) *UserOrderHandler {
    return &UserOrderHandler {
        wsapi: apiws,
        notifyHandler: notifyHandler,
        userOrders: sync.Map{},
    }
}
func (this *UserOrderHandler) Loop() {
    this.userOrders.Range(func(key, value interface{}) bool {
        fmt.Printf("%v: %v\n", key, value)
        return true
    })
}

func (this *UserOrderHandler) Handle(ctx context.Context, setdataf func(*[]api.WsUserOrder, *sync.Map)) {
ws:
    var userOrderDataChan chan api.WsUserOrderDatas
    userOrderDataChan, this.c = this.wsapi.RunUserOrderConsumer(ctx)
    for {
        userOrder := <- userOrderDataChan
        if userOrder.Err != nil {
            // this.notifyHandler.MsgChan <- fmt.Sprintf("[UserOrderHandler] %s", userOrder.Err.Error())
            close(this.c)
            goto ws
        }
        setdataf(&userOrder.Datas, &this.userOrders)
    }
}

func (this *UserOrderHandler) DeleteCompletedOrder() {
    ticker := time.NewTicker(time.Minute * 10)
    for {
        <-ticker.C
        this.userOrders.Range(func(key, value any)bool {
            data := value.(api.WsUserOrder)
            if data.Status == api.OrderStatusDone {
                this.userOrders.Delete(key)
            }
            return true
        })
    }
}

func (this *UserOrderHandler) Get(id string) api.WsUserOrder {
    value, ok := this.userOrders.Load(id)
    if !ok {
        return api.WsUserOrder{}
    }
    data := value.(api.WsUserOrder)
    return data
}

func (this *UserOrderHandler) Stop() {
    if this.c != nil {
        close(this.c)
    }
}
