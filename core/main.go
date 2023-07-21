package core

import (
	"errors"

	"github.com/go-resty/resty/v2"
)

var (
	client = resty.New()
)

func init() {
	client.SetBaseURL("https://api.telegram.org")
}

// opts.offset? {int}
//
// opts.limit? {int=100}
//
// opts.timeout? {int=0}
//
// opts.allow_updates? {[]string}
func GetUpdates(token string, opts ...map[string]any) ([]*Update, error) {
	body := &Result[[]*Update]{}

	var data map[string]any

	if len(opts) > 0 {
		data = opts[0]
	}

	if _, err := client.R().
		SetPathParams(map[string]string{
			"token":  token,
			"method": "getUpdates",
		}).
		SetBody(data).
		SetResult(body).
		SetError(body).
		Post("/bot{token}/{method}"); err != nil {
		return nil, err
	} else if !body.Ok {
		return nil, errors.New(body.Description)
	}

	return body.Result, nil
}

// token: Telegrame Bot's Token
func GetMe(token string) (*User, error) {
	body := &Result[*User]{}

	if _, err := client.R().
		SetError(body).
		SetResult(body).
		SetPathParams(map[string]string{
			"token":  token,
			"method": "getMe",
		}).
		Get("/bot{token}/{method}"); err != nil {
		return nil, err
	} else if !body.Ok {
		return nil, errors.New(body.Description)
	}

	return body.Result, nil
}

// opts.commands {[]*BotCommand}
//
// opts.scope? {*BotCommandScope}
//
// opts.code? {string}
func SetMyCommands(token string, opts map[string]any) error {
	body := &Result[bool]{}

	if _, err := client.R().
		SetPathParams(map[string]string{
			"token":  token,
			"method": "setMyCommands",
		}).
		SetBody(opts).
		SetResult(body).
		SetError(body).
		Post("/bot{token}/{method}"); err != nil {
		return err
	} else if !body.Ok {
		return errors.New(body.Description)
	}

	return nil
}

// opts.chat_id? {string}
//
// opts.message_id? {string}
//
// opts.inline_message_id? {string}
//
// opts.text {string}
//
// opts.parse_mode? {MarkdownV2 | HTML | Markdown}
//
// opts.entities? {[]MessageEntity}
//
// opts.disable_web_page_preview? {bool}
//
// opts.reply_markup {InlineKeyboardMarkup | ReplyKeyboardMarkup | ReplyKeyboardRemove | ForceReply}
func EditMessageText(token string, opts map[string]any) error {
	body := &Result[bool]{}

	if _, err := client.R().
		SetPathParams(map[string]string{
			"token":  token,
			"method": "editMessageText",
		}).
		SetBody(opts).
		SetResult(body).
		SetError(body).
		Post("/bot{token}/{method}"); err != nil {
		return err
	} else if !body.Ok {
		return errors.New(body.Description)
	}

	return nil
}

func AnswerCallbackQuery() {

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
func SendMessage(token string, opts map[string]any) (*Message, error) {
	body := &Result[*Message]{}

	if _, err := client.R().
		SetPathParams(map[string]string{
			"token":  token,
			"method": "sendMessage",
		}).
		SetBody(opts).
		SetResult(body).
		SetError(body).
		Post("/bot{token}/{method}"); err != nil {
		return nil, err
	} else if !body.Ok {
		return nil, errors.New(body.Description)
	}

	return body.Result, nil
}
