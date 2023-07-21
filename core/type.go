package core

type Result[T any] struct {
	Result      T      `json:"result"`
	Ok          bool   `json:"ok"`
	ErrCode     int    `json:"err_code,omitempty"`
	Description string `json:"description,omitempty"`
}

type Update struct {
	UpdateId      int            `json:"update_id"`
	Message       *Message       `json:"message,omitempty"`
	CallbackQuery *CallbackQuery `json:"callback_query,omitempty"`
	InlineQuery   *InlineQuery   `json:"inline_query,omitempty"`
}

type Message struct {
	MessageId int `json:"message_id"`

	From *struct {
		Id        int    `json:"id"`
		IsBot     bool   `json:"is_bot"`
		FirstName string `json:"first_name"`
		Username  string `json:"username"`
		Language  string `json:"language_code"`
	} `json:"from"`

	Chat *struct {
		Id                          int    `json:"id"`
		FirstName                   string `json:"first_name"`
		Username                    string `json:"username"`
		Type                        string `json:"type"`
		AllMembersAreAdministrators bool   `json:"all_members_are_administrators"`
	} `json:"chat"`

	Date    int      `json:"date"`
	Text    string   `json:"text,omitempty"`
	Sticker *Sticker `json:"sticker,omitempty"`

	LeftChatParticipant *struct {
		Id        int    `json:"id"`
		IsBot     bool   `json:"is_bot"`
		FirstName string `json:"first_name"`
		Username  string `json:"username"`
	} `json:"left_chat_participant,omitempty"`

	LeftChatMember *struct {
		Id        int    `json:"id"`
		IsBot     bool   `json:"is_bot"`
		FirstName string `json:"first_name"`
		Username  string `json:"username"`
	} `json:"left_chat_member,omitempty"`

	// https://core.telegram.org/bots/api#messageentity
	Entities []*MessageEntity `json:"entities,omitempty"`

	NewChatParticipant *struct {
		Id        int    `json:"id"`
		IsBot     bool   `json:"is_bot"`
		FirstName string `json:"first_name"`
		Username  string `json:"username"`
	} `json:"new_chat_participant,omitempty"`

	NewChatMember *struct {
		Id        int    `json:"id"`
		IsBot     bool   `json:"is_bot"`
		FirstName string `json:"first_name"`
		Username  string `json:"username"`
	} `json:"new_chat_member,omitempty"`

	NewChatMembers []*struct {
		Id        int    `json:"id"`
		IsBot     bool   `json:"is_bot"`
		FirstName string `json:"first_name"`
		Username  string `json:"username"`
	} `json:"new_chat_members,omitempty"`
}

type Sticker struct {
	Thumbnail *struct {
		FileId       string `json:"file_id"`
		FileUniqueId string `json:"file_unique_id"`
		FileSize     int    `json:"file_size"`
		Width        int    `json:"width"`
		Height       int    `json:"height"`
	} `json:"thumbnail,omitempty"`
	Thumb *struct {
		FileId       string `json:"file_id"`
		FileUniqueId string `json:"file_unique_id"`
		FileSize     int    `json:"file_size"`
		Width        int    `json:"width"`
		Height       int    `json:"height"`
	} `json:"thumb,omitempty"`
	FileId       string `json:"file_id,omitempty"`
	FileUniqueId string `json:"file_unique_id,omitempty"`
	FileSize     int    `json:"file_size,omitempty"`
}

type BotCommand struct {
	Command     string `json:"command"`
	Description string `json:"description"`
}

type BotCommandScope struct {
	// type:
	//		chat
	// 		chat_administrators
	//  	chat_member
	// 		all_chat_administrators
	// 		all_group_chats
	//    all_private_chats
	// 		default
	Type string `json:"type"`
	// string or int
	ChatId any `json:"chat_id,omitempty"`
	UserId int `json:"user_id,omitempty"`
}

type User struct {
	Id       int    `json:"id"`
	IsBot    bool   `json:"is_bot"`
	Username string `json:"username"`
	Language string `json:"language_code"`
}

type InlineKeyboardButton struct {
	Text         string `json:"text"`
	Url          string `json:"url,omitempty"`
	CallbackData string `json:"callback_data,omitempty"`
}

type InlineKeyboardMarkup struct {
	InlineKeyboards [][]*InlineKeyboardButton `json:"inline_keyboard"`
}

type CallbackQuery struct {
	Id              string   `json:"id"`
	From            *User    `json:"from"`
	Message         *Message `json:"message"`
	InlineMessageId string   `json:"inline_message_id"`
	Data            string   `json:"data"`
}

type InlineQuery struct {
	Id       string `json:"id"`
	From     *User  `json:"from"`
	Query    string `json:"query"`
	Offset   string `json:"offset"`
	ChatType string `json:"chat_type"`
}

// https://core.telegram.org/bots/api#messageentity
type MessageEntity struct {
	Offset int `json:"offset"`
	Length int `json:"length"`
	/* mention | hashtag | cashtag | bot_command | url | email | phone_number | bold | italic | text_mention | spoiler | strikethrough | underline ...*/
	Type string `json:"type"`

	Url           string `json:"url"`
	User          *User  `json:"user"`
	Language      string `json:"language"`
	CustomEmojiId string `json:"custom_emoji_id"`
}
