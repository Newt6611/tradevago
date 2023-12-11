package discord

import (
	"context"

	_ "github.com/Newt6611/tradevago/pkg/notify"
	"github.com/bwmarrin/discordgo"
)

type DiscordClient struct {
    discord *discordgo.Session
}

func NewDiscordClient(token string) *DiscordClient {
    d, _ := discordgo.New(token)
    return &DiscordClient {
        discord: d,
    }
}

func (this *DiscordClient) SendMsg(ctx context.Context, channelId string, msg string) {
    this.discord.ChannelMessageSend(channelId, msg)
}
