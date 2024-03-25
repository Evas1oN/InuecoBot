package commands

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/Evas1oN/inuecobot/db"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func CreateHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	log.Print(update.Message.Text)
	args := strings.Split(update.Message.Text, " ")

	switch args[1] {
	case "group":
		res := db.Database.Create(&db.Group{Code: args[2]})
		if res.Error != nil {
			log.Fatal(res.Error)
		}

		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   args[2],
		})

	case "subject":
		var sb strings.Builder

		for i := 2; i < len(args); i++ {
			sb.WriteString(fmt.Sprintf("%s ", args[i]))
		}

		db.Database.Create(&db.Subject{Name: sb.String()})
		return

	case "pair":
		var sb strings.Builder
		const layout = "02.01.2006 15:04 -07"
		time, err := time.Parse(layout, args[3]+" "+args[4]+" +05")

		if err != nil {
			b.SendMessage(ctx, &bot.SendMessageParams{
				ChatID: update.Message.Chat.ID,
				Text:   err.Error(),
			})
			return
		}

		var group db.Group

		result := db.Database.Where(&db.Group{Code: args[2]}).First(&group)

		if result.Error != nil {
			log.Fatal(result.Error)
		}

		for i := 5; i < len(args); i++ {
			sb.WriteString(fmt.Sprintf("%s ", args[i]))
		}

		db.Database.Create(&db.Pair{
			Group:     group,
			StartTime: time,
			Subject:   sb.String(),
		})

		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   sb.String(),
		})
	}
}
