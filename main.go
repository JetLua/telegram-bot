package main

import (
	"log"
	"time"

	"github.com/jetlua/telegram-bot/core"
)

type Bot struct {
	token   string
	Self    *core.User
	Channel chan *core.Update
}

type BotConfig struct {
	Token string `json:"token"`
	// 轮询 GetUpdates 的时间间隔 单位：秒 默认1s
	Interval int `json:"interval"`
}

func New(config *BotConfig) *Bot {
	if config.Token == "" {
		log.Fatalln("token: 不能为空")
	}

	bot := &Bot{
		token:   config.Token,
		Channel: make(chan *core.Update),
	}

	if u, err := core.GetMe(config.Token); err != nil {
		log.Fatalln(err)
	} else {
		bot.Self = u
	}

	go func() {
		params := map[string]any{
			"offset": 0,
		}

		for {
			if updates, err := core.GetUpdates(bot.token, params); err != nil {
				log.Println(err)
			} else {
				for _, v := range updates {
					bot.Channel <- v
					params["offset"] = v.UpdateId + 1
				}
			}
			time.Sleep(time.Second * time.Duration(config.Interval))
		}
	}()

	return bot
}

func (bot *Bot) GetMe() (*core.User, error) {
	return core.GetMe(bot.token)
}

// opts.chat_id {string | int}
//
// opts.message_thread_id? {int}
//
// opts.text {string}
//
// opts.parse_mode? {MarkdownV2 | HTML | Markdown}
//
// opts.entities? {[]MessageEntity}
//
// opts.disable_web_page_preview? {bool}
//
// opts.disable_notification {bool}
//
// opts.protect_content {bool}
//
// opts.reply_to_message_id {int}
//
// opts.allow_sending_without_reply {bool}
//
// opts.reply_markup {InlineKeyboardMarkup | ReplyKeyboardMarkup | ReplyKeyboardRemove | ForceReply}
func (bot *Bot) SendMessage(opts map[string]any) (*core.Message, error) {
	return core.SendMessage(bot.token, opts)
}

// opts.commands {[]*BotCommand}
//
// opts.scope? {*BotCommandScope}
//
// opts.code? {string}
func (bot *Bot) SetMyCommands(opts map[string]any) error {
	return core.SetMyCommands(bot.token, opts)
}
