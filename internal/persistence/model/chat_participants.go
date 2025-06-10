package model

type ChatParticipants struct {
	UserId int32
	ChatId int32
}

func NewChatParticipants(userId int32, chatId int32) *ChatParticipants {
	return &ChatParticipants{
		UserId: userId,
		ChatId: chatId,
	}
}

func (cp *ChatParticipants) ToMap() map[string]any {
	return map[string]any{
		"user_id": cp.UserId,
		"chat_id": cp.ChatId,
	}
}

func MapToChatParticipants(data map[string]any) *ChatParticipants {
	return &ChatParticipants{
		UserId: data["user_id"].(int32),
		ChatId: data["chat_id"].(int32),
	}
}