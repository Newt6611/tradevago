package tri

import (
	"context"
	"sync"

	"github.com/Newt6611/tradevago/pkg/api"
)


type BalanceHandler struct {
    wsapi           *api.WSApi
    notifyHandler   *NotifyHandler
    balances        sync.Map
    c               chan struct{}
}

func NewBalanceHandler(apiws *api.WSApi, notifyHandler *NotifyHandler) *BalanceHandler {
    return &BalanceHandler{
        wsapi: apiws,
        notifyHandler: notifyHandler,
        balances: sync.Map{},
    }
}

func (this *BalanceHandler) Handle(ctx context.Context, setdataf func(*[]api.WsUserAccountBalance, *sync.Map)) {
ws:
    var userAccountChan chan api.WsUserAccountDatas

    userAccountChan, this.c = this.wsapi.RunAccountConsumer(ctx)
    for {
        userAccountData := <-userAccountChan
        if userAccountData.Err != nil {
            // this.notifyHandler.MsgChan <- fmt.Sprintf("[BalanceHandler] %s", userAccountData.Err.Error())
            close(this.c)
            goto ws
        }
        setdataf(&userAccountData.Datas, &this.balances)
    }
}

func (this *BalanceHandler) IsReady(currency string) bool {
    _, ok := this.balances.Load(currency)
    return ok
}

func (this *BalanceHandler) Get(currency string) api.WsUserAccountBalance {
    value, _ := this.balances.Load(currency)
    return value.(api.WsUserAccountBalance)
}

func (this *BalanceHandler) Stop() {
    if this.c != nil {
        close(this.c)
    }
}
