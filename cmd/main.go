package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"pingPongBot/pkg/bot"
)

func main() {
	token := os.Getenv("token")
	b := bot.NewBot(token)

	quitCtx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	b.Run(quitCtx)
}
