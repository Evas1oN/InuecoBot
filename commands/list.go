package commands

import (
	"context"
	"log"
	"os"
	"strings"

	"github.com/Evas1oN/inuecobot/db"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/jedib0t/go-pretty/table"
)

func List(ctx context.Context, b *bot.Bot, update *models.Update) {
	log.Print(update.Message.Text)
	args := strings.Split(update.Message.Text, " ")
	tbl := table.NewWriter()
	tbl.SetOutputMirror(os.Stdout)

	switch args[1] {
	case "group":
		tbl.AppendHeader(table.Row{"#", "Группа"})
		var groups []db.Group
		db.Database.Find(&groups)

		for _, group := range groups {
			tbl.AppendRow(table.Row{group.ID, group.Code})
		}
		msg := tbl.Render()
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID:    update.Message.Chat.ID,
			Text:      "<pre>" + msg + "</pre>",
			ParseMode: models.ParseModeHTML,
		})

	case "subject":
	case "pair":
		tbl.AppendHeader(table.Row{"#", "Группа", "Предмет", "Время"})
		var pairs []db.Pair
		db.Database.Joins("Group").Find(&pairs)

		for _, pair := range pairs {
			tbl.AppendRow(table.Row{pair.ID, pair.Group.Code, pair.Subject, pair.StartTime.UTC().Local()})
		}

		msg := tbl.Render()
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID:    update.Message.Chat.ID,
			Text:      "<pre>" + msg + "</pre>",
			ParseMode: models.ParseModeHTML,
		})
	}
}
