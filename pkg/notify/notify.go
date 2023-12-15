package notify

import "context"

const (
	Sticker string = "STICKER"
)

type Notifier interface {
	SendMsg(ctx context.Context, msg string)
	SendCodeMsg(ctx context.Context, msg string)
	SendInlineCodeMsg(ctx context.Context, msg string)
	SendBoldMsg(ctx context.Context, msg string)
	SendItalicMsg(ctx context.Context, msg string)
	HandleMessage(map[string]func() string)
	Close()
}
