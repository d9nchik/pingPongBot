package bot

import (
	"context"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Bot struct {
	bot *tgbotapi.BotAPI
}

func NewBot(token string) *Bot {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	log.Printf("Authorized on account %s", bot.Self.UserName)

	return &Bot{bot: bot}
}

func (b *Bot) Run(ctx context.Context) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := b.bot.GetUpdatesChan(u)
	for {
		select {
		case update := <-updates:
			if update.Message != nil { // If we got a message
				if update.Message.Command() == "ping" {
					log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "pong")
					b.bot.Send(msg)
				}
			}
		case <-ctx.Done():
			return
		}
	}
}
