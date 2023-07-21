package tgbot

import (
	"log"
	"testing"

	"github.com/jetlua/telegram-bot/core"
)

func TestMain(m *testing.M) {
	bot := New(&BotConfig{
		Interval: 1,
		Token:    "1486089197:AAFPiHlhGFWDqOwOT32LNGDo_YcGDxPgwSo",
	})

	if err := bot.SetMyCommands(map[string]any{
		"commands": []*core.BotCommand{
			{Command: "/start", Description: "🛫️"},
		},
		"scope": &core.BotCommandScope{
			Type:   "chat",
			ChatId: "-493798401",
		},
	}); err != nil {
		log.Println(err)
	}

	for u := range bot.Channel {
		handle(u)
	}
}

func handle(u *core.Update) {
	m := u.Message

	if m != nil && m.Entities != nil && m.Entities[0].Type == "bot_command" {
		println(m.Text)
	}
}

func init() {
	log.SetFlags(log.Lshortfile)
}
