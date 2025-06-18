package model

import (
	"time"
)

type ChatMessage struct {
	MessageId int32     `json:"message_id"`
	AuthorId  int32     `json:"author_id"`
	ChatId    int32     `json:"chat_id"`
	Message   string    `json:"message"`
	SentAt    time.Time `json:"sent_at"`
}

func NewChatMessage(
	messageId,
	authorId, 
	chatId int32, 
	message string) *ChatMessage {
	return &ChatMessage{
		MessageId: messageId,
		AuthorId:  authorId,
		ChatId:    chatId,
		Message:   message,
		SentAt:    time.Now(),
	}
}

func (message *ChatMessage) ToMap() map[string]any {
	return map[string]any{
		"message_id": message.MessageId,
		"author_id":  message.AuthorId,
		"chat_id":    message.ChatId,
		"message":    message.Message,
		"sent_at":    message.SentAt,
	}
}

func MapToChatMessage(data map[string]any) *ChatMessage {
	return &ChatMessage{
		MessageId: data["message_id"].(int32),
		AuthorId:  data["author_id"].(int32),
		ChatId:    data["chat_id"].(int32),
		Message:   data["message"].(string),
		SentAt:    data["sent_at"].(time.Time),
	}
}