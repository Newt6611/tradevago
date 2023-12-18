package telegram

import (
	"context"
	"fmt"

	"github.com/Newt6611/tradevago/pkg/notify"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type TelegramClient struct {
	telegram  *tgbotapi.BotAPI
	channelId int64
}

func NewTelegramClient(token string, channelId int64) *TelegramClient {
	bot, _ := tgbotapi.NewBotAPI(token)
	return &TelegramClient{
		telegram:  bot,
		channelId: channelId,
	}
}

func (this *TelegramClient) SendMsg(ctx context.Context, msg string) {
	message := this.newMarkDownMessage()
	message.Text = msg
	this.telegram.Send(message)
}

func (this *TelegramClient) SendCodeMsg(ctx context.Context, msg string) {
	message := this.newMarkDownMessage()
	message.Text = "```" + msg + "```"
	this.telegram.Send(message)
}

func (this *TelegramClient) SendInlineCodeMsg(ctx context.Context, msg string) {
	message := this.newMarkDownMessage()
	message.Text = "`" + msg + "`"
	this.telegram.Send(message)
}

func (this *TelegramClient) SendBoldMsg(ctx context.Context, msg string) {
	message := this.newMarkDownMessage()
	message.Text = "*" + msg + "*"
	this.telegram.Send(message)
}

func (this *TelegramClient) SendItalicMsg(ctx context.Context, msg string) {
	message := this.newMarkDownMessage()
	message.Text = "_" + msg + "_"
	this.telegram.Send(message)
}

func (this *TelegramClient) HandleMessage(patterns map[string]func() string) {
	// this.telegram.Debug = true
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := this.telegram.GetUpdatesChan(u)

	for update := range updates {
		if update.Message.Sticker != nil {
			if f, ok := patterns[notify.Sticker]; ok {
				msg := this.newMarkDownMessage()
				msg.Text = f()
				this.telegram.Send(msg)
				continue
			}
		}
		if update.Message != nil { // If we got a message
            fmt.Println(update.Message.Chat.ID)
			if f, ok := patterns[update.Message.Text]; ok {
				msg := this.newMarkDownMessage()
				msg.Text = f()
				this.telegram.Send(msg)
			}
		}
	}
}

func (this *TelegramClient) Close() {}

func (this *TelegramClient) newMarkDownMessage() tgbotapi.MessageConfig {
	message := tgbotapi.NewMessage(this.channelId, "")
	message.ParseMode = "MarkDown"
	return message
}
