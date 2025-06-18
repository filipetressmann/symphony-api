package request_model

import (
	"time"
	"symphony-api/internal/persistence/model"
)

type BaseChatData struct {
	ChatId int32 `json:"chat_id" binding:"required"`
	CreatedAt time.Time `json:"created_at" binding:"required"`
}

type CreateChatRequest struct {
	Username1 string `json:"username1" binding:"required"`
	Username2 string `json:"username2" binding:"required"`
}

type GetChatByIdRequest struct {
	ChatId int32 `schema:"chat_id,required"`
}

type ListUsersFromChatRequest struct {
	ChatId int32 `schema:"chat_id,required"`
}

type ListUsersFromChatResponse struct {
	Username1 string `json:"username1" binding:"required"`
	Username2 string `json:"username2" binding:"required"`
}

type ListChatsFromUserRequest struct {
	Username string `schema:"username,required"`
}

type ListChatsFromUserResponse struct {
	ChatIds []int32 `json:"chat_ids" binding:"required"`
}

func NewBaseChatData(chatId int32, createdAt time.Time) *BaseChatData {
	return &BaseChatData{
		ChatId: chatId,
		CreatedAt: createdAt,
	}
}

type AddMessageToChatRequest struct {
	ChatId  int32  `json:"chat_id" binding:"required"`
	AuthorId int32  `json:"author_id" binding:"required"`
	Message string `json:"message" binding:"required"`
}

type AddMessageToChatResponse struct {
	MessageId int32     `json:"message_id" binding:"required"`
	AuthorId  int32     `json:"author_id" binding:"required"`
	ChatId    int32     `json:"chat_id" binding:"required"`
	SentAt    time.Time `json:"sent_at" binding:"required"`
}

func NewAddMessageToChatResponse(messageId, authorId, chatId int32, sentAt time.Time) *AddMessageToChatResponse {
	return &AddMessageToChatResponse{
		MessageId: messageId,
		AuthorId:  authorId,
		ChatId:    chatId,
		SentAt:    sentAt,
	}
}
	
type ListMessagesFromChatRequest struct {
	ChatId int32 `schema:"chat_id,required"`
	Limit int32 `schema:"limit,default=10"`
}

type MessagesFromChat struct {
	AuthorId int32     `json:"author_id" binding:"required"`
	SentAt	time.Time `json:"sent_at" binding:"required"`
	Message string    `json:"message" binding:"required"`
}

type ListMessagesFromChatResponse struct {
	ChatId  int32              `json:"chat_id" binding:"required"`
	Messages []MessagesFromChat `json:"messages" binding:"required"`
}

func MapsToMessagesFromChat(message []*model.ChatMessage) *ListMessagesFromChatResponse {
	messages := make([]MessagesFromChat, len(message))
	for i, msg := range message {
		messages[i] = MessagesFromChat{
			AuthorId: msg.AuthorId,
			SentAt:   msg.SentAt,
			Message:  msg.Message,
		}
	}
	return &ListMessagesFromChatResponse{
		ChatId:   message[0].ChatId,
		Messages: messages,
	}
}