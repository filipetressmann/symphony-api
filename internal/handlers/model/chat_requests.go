package request_model

import (
	"time"
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



