package model

import (
	"time"
)

type Chat struct {
	ChatId int32
	CreatedAt time.Time
}

func NewChat(chatId int32) *Chat {
	return &Chat{
		ChatId: chatId,
		CreatedAt: time.Now(),
	}
}

func (chat *Chat) ToMap() map[string]any {
    return map[string]any{}
}

func MapToChat(data map[string]any) *Chat {
	return &Chat{
		ChatId: data["chat_id"].(int32),
		CreatedAt: data["created_at"].(time.Time),
	}
}