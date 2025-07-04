package repository

import (
	"errors"
	"symphony-api/internal/persistence/connectors/postgres"
	"symphony-api/internal/persistence/model"
)

const CHAT_TABLE_NAME = "CHAT"
const USER_TO_CHAT_TABLE = "CHAT_PARTICIPANTS"
const CHAT_MESSAGE_TABLE = "CHAT_MESSAGE"
const JOINED_USERS_AND_CHAT_PARTICIPANTS = "USERS u JOIN CHAT_PARTICIPANTS cp ON u.id = cp.user_id"
const JOINED_CHATS_AND_PARTICIPANTS = "CHAT c JOIN CHAT_PARTICIPANTS cp ON c.chat_id = cp.chat_id"

type ChatRepository struct {
	connection postgres.PostgreConnection
}

func NewChatRepository(connection postgres.PostgreConnection) *ChatRepository {
	return &ChatRepository{
		connection: connection,
	}
}

func (repository *ChatRepository) Put(chat *model.Chat) error {
    id, err := repository.connection.PutReturningId(chat.ToMap(), CHAT_TABLE_NAME, "chat_id")
    if err != nil {
        return err
    }
    chat.ChatId = id.(int32)
    return nil
}

func (repository *ChatRepository) GetByChatId(chatId int32) (*model.Chat, error) {
	constraint := map[string]any{
		"chat_id": chatId,
	}

	chat, err := repository.connection.Get(constraint, CHAT_TABLE_NAME)

	if len(chat) == 0 {
		return nil, errors.New("chat not found")
	}

	return model.MapToChat(chat[0]), err
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

    chatsData, err := repository.connection.Get(
        constraint,
        JOINED_CHATS_AND_PARTICIPANTS,
    )

    if err != nil {
        return nil, err
    }

    var chatList []*model.Chat
    for _, chatData := range chatsData {
        chatList = append(chatList, model.MapToChat(chatData))
    }

    return chatList, nil
}

func (repository *ChatRepository) FindChatByUsers(userId1, userId2 int32) (*model.Chat, error) {
    user1 := &model.User{UserId: userId1}
    chats, err := repository.ListChatsByUser(user1)
    if err != nil {
        return nil, err
    }
    for _, chat := range chats {
        users, err := repository.ListUsersFromChat(chat)
        if err != nil {
            continue
        }
        found1, found2 := false, false
        for _, u := range users {
            if u.UserId == userId1 {
                found1 = true
            }
            if u.UserId == userId2 {
                found2 = true
            }
        }
        if found1 && found2 && len(users) == 2 {
            return chat, nil
        }
    }
    return nil, nil
}

func (repository *ChatRepository) AddMessageToChatAndReturn(chatId int32, authorId int32, message string) (*model.ChatMessage, error) {
    data := map[string]any{
        "chat_id":  chatId,
        "author_id": authorId,
        "message":   message,
    }
    id, err := repository.connection.PutReturningId(data, CHAT_MESSAGE_TABLE, "message_id")
    if err != nil {
        return nil, err
    }

    constraint := map[string]any{
        "message_id": id,
    }
    msgs, err := repository.connection.Get(constraint, CHAT_MESSAGE_TABLE)
    if err != nil || len(msgs) == 0 {
        return nil, errors.New("could not retrieve inserted message")
    }
    return model.MapToChatMessage(msgs[0]), nil
}

func (repository *ChatRepository) ListMessagesFromChat(chatId int32, limit int32) ([]*model.ChatMessage, error) {
    messagesData, err := repository.connection.GetChatWithLimit(
        chatId,
        limit,
        CHAT_MESSAGE_TABLE,
    )

    if err != nil {
        return nil, err
    }

    var messages []*model.ChatMessage
    for _, messageData := range messagesData {
        messages = append(messages, model.MapToChatMessage(messageData))
    }

    return messages, nil
}