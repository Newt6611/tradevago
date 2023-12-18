package tri

import (
	"context"

	"github.com/Newt6611/tradevago/pkg/notify"
)

type NotifyHandler struct {
	msgChan chan string
	notifier notify.Notifier
}

func NewNotifyHandler(notifier notify.Notifier) *NotifyHandler {
	return &NotifyHandler{
		msgChan:  make(chan string, 1000),
		notifier: notifier,
	}
}

func (this *NotifyHandler) Handle(ctx context.Context) {
	premsg := ""
	for {
		msg := <-this.msgChan
		if msg == premsg {
			continue
		}

		this.notifier.SendInlineCodeMsg(ctx, msg)
		premsg = msg
	}
}

func (this *NotifyHandler) SendMsg(msg string) {
    this.msgChan <- msg
}

func (this *NotifyHandler) HandleMessage(patterns map[string]func() string) {
	this.notifier.HandleMessage(patterns)
}

func (this *NotifyHandler) Stop() {
	this.notifier.Close()
}
