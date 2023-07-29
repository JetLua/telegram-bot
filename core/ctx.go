package core

type Ctx struct {
	token    string
	Msg      *Message
	Query    *CallbackQuery
	Entities []*MessageEntity
	Command  chan string
}

func NewCtx(token string, u *Update) *Ctx {
	ctx := ctxPool.Get().(*Ctx)
	ctx.token = token
	ctx.Msg = u.Message
	ctx.Query = u.CallbackQuery
	return ctx
}

// 注意：如果是回复指定消息，仍需指定 reply_to_message_id
//
//	c.Reply("ok", map[string]any{
//		"reply_markup": [][]string{
//			{"你好", "1"},
//			{"Hello", "2"},
//			{"Приве́т", "3"},
//		},
//	})
//
// 多组按钮
//
//	c.Reply("ok", map[string]any{
//		"reply_markup": [][][]string{
//			[][]string{
//				{"你好", "1"},
//				{"Hello", "2"},
//				{"Приве́т", "3"},
//			},
//			[][]string{
//				{"你好", "1"},
//				{"Hello", "2"},
//				{"Приве́т", "3"},
//			},
//		},
//	})
//
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
// opts.reply_markup {InlineKeyboardMarkup | ReplyKeyboardMarkup | ReplyKeyboardRemove | ForceReply | [][]string}
func (c *Ctx) Reply(text string, opts ...map[string]any) {
	var chatId any

	if c.Msg != nil {
		chatId = c.Msg.Chat.Id
	} else if c.Query != nil {
		chatId = c.Query.Message.Chat.Id
	}

	var markup *InlineKeyboardMarkup

	defer func() {
		if markup != nil {
			markup.Dispose()
		}
	}()

	if len(opts) > 0 {
		rm := opts[0]["reply_markup"]

		if rm != nil {
			if v, ok := rm.([][]string); ok {
				kb := make(InlineKeyboard, 0)

				for _, item := range v {
					kb.Text(item[0], item[1])
				}

				markup = &InlineKeyboardMarkup{
					InlineKeyboard: [][]*InlineKeyboardButton{kb},
				}
			} else if v, ok := rm.([][][]string); ok {
				markup = &InlineKeyboardMarkup{
					InlineKeyboard: [][]*InlineKeyboardButton{},
				}

				for _, rows := range v {
					kb := make(InlineKeyboard, 0)

					for _, item := range rows {
						kb.Text(item[0], item[1])
					}

					markup.InlineKeyboard = append(markup.InlineKeyboard, kb)
				}
			}
		}
	}

	if _, err := SendMessage(c.token, map[string]any{
		"chat_id":      chatId,
		"text":         text,
		"reply_markup": markup,
	}); err != nil {
		println(err.Error())
	}
}

func (c *Ctx) Dispose() {
	ctxPool.Put(c)
}
