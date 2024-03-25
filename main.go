package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/Evas1oN/inuecobot/commands"
	"github.com/Evas1oN/inuecobot/db"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func init() {
	db.Init()
}

func main() {
	log.Print("weeee ha")

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	opts := []bot.Option{
		bot.WithDefaultHandler(handler),
	}

	b, err := bot.New("6596136032:AAEix9MkpBOfCUX4OXLpL6csvLCOSJVmfy8", opts...)
	if err != nil {
		panic(err)
	}

	b.RegisterHandler(bot.HandlerTypeMessageText, "/create", bot.MatchTypePrefix, commands.CreateHandler)
	b.RegisterHandler(bot.HandlerTypeMessageText, "/list", bot.MatchTypePrefix, commands.List)

	b.Start(ctx)
}

func handler(ctx context.Context, b *bot.Bot, update *models.Update) {
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   update.Message.Text,
	})
}
