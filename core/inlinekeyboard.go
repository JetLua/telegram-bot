package core

import "sync"

type InlineKeyboard []*InlineKeyboardButton

var btnPool = &sync.Pool{
	New: func() any {
		return &InlineKeyboardButton{}
	},
}

func (kb *InlineKeyboard) Text(label, data string) *InlineKeyboard {
	btn := btnPool.Get().(*InlineKeyboardButton)
	btn.Text = label
	btn.CallbackData = data
	*kb = append(*kb, btn)
	return kb
}
