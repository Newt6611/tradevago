package tri

import (
	"context"

	"github.com/Newt6611/tradevago/pkg/notify"
)

type NotifyHandler struct {
    MsgChan chan string
    notifier notify.Notifier
}

func NewNotifyHandler(notifier notify.Notifier) *NotifyHandler {
    return &NotifyHandler{
        MsgChan: make(chan string, 1000),
        notifier: notifier,
    }
}

func (this *NotifyHandler) Handle(ctx context.Context) {
    premsg := ""
    for {
        msg := <-this.MsgChan
        if msg == premsg {
            continue
        }

        this.notifier.SendCodeMsg(ctx, msg)
    }
}

func (this *NotifyHandler) Stop() {
    this.notifier.Close()
}
