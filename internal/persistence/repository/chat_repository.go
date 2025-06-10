package repository

import (
	"errors"
	"symphony-api/internal/persistence/connectors/postgres"
	"symphony-api/internal/persistence/model"
)

const CHAT_TABLE_NAME = "CHAT"
const USER_TO_CHAT_TABLE = "CHAT_PARTICIPANTS"
const JOINED_USERS_AND_CHAT_PARTICIPANTS = "USERS u JOIN CHAT_PARTICIPANTS cp ON u.id = cp.user_id"

type ChatRepository struct {
	connection postgres.PostgreConnection
}

func NewChatRepository(connection postgres.PostgreConnection) *ChatRepository {
	return &ChatRepository{
		connection: connection,
	}
}

func (repository *ChatRepository) Put(chat *model.Chat) error {
	return repository.connection.Put(chat.ToMap(), CHAT_TABLE_NAME)
}

func (repository *ChatRepository) GetByChatId(chatId int32) (*model.Chat, error) {
	constraint := map[string]any{
		"chat_id": chatId,
	}

	chats, err := repository.connection.Get(constraint, CHAT_TABLE_NAME)

	if len(chats) == 0 {
		return nil, errors.New("chat not found")
	}

	return model.MapToChat(chats[0]), err
}

func (repository *ChatRepository) AddUserToChat(user *model.User, chat *model.Chat) error {
	return repository.connection.Put(
		map[string]any{
			"chat_id": chat.ChatId,
			"user_id": user.UserId,
		},
		USER_TO_CHAT_TABLE,
	)
}

func (repository *ChatRepository) ListUsersFromChat(chat *model.Chat) ([]*model.User, error) {
	constraint := map[string]any{
		"cp.chat_id": chat.ChatId,
	}

	users, err := repository.connection.Get(
		constraint,
		JOINED_USERS_AND_CHAT_PARTICIPANTS,
	)

	if err != nil {
		return nil, err
	}

	var userList []*model.User
	for _, userData := range users {
		userList = append(userList, model.MapToUser(userData))
	}

	return userList, nil
}

func (repository *ChatRepository) ListChatsByUser(user *model.User) ([]*model.Chat, error) {
	constraint := map[string]any{
		"cp.user_id": user.UserId,
	}

	chats, err := repository.connection.Get(
		constraint,
		JOINED_USERS_AND_CHAT_PARTICIPANTS,
	)

	if err != nil {
		return nil, err
	}

	var chatList []*model.Chat
	for _, chatData := range chats {
		chatList = append(chatList, model.MapToChat(chatData))
	}

	return chatList, nil
}

