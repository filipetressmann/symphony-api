package service

import (
	"errors"
	"symphony-api/internal/persistence/model"
	"symphony-api/internal/persistence/repository"
)

type ChatService struct {
	chatRepository *repository.ChatRepository
	userRepository *repository.UserRepository
}

func NewChatService(
	chatRepository *repository.ChatRepository,
	userRepository *repository.UserRepository,
) *ChatService {
	return &ChatService{
		chatRepository: chatRepository,
		userRepository: userRepository,
	}
}

func (service *ChatService) GetChatById(chatId int32) (*model.Chat, error) {
	chat, err := service.chatRepository.GetByChatId(chatId)
	if err != nil {
		return nil, errors.New("chat does not exist")
	}
	if chat == nil {
		return nil, errors.New("chat not found")
	}
	return chat, nil
}

func (service *ChatService) CreateChat(username1, username2 string) (*model.Chat, error) {
    if username1 == "" || username2 == "" {
        return nil, errors.New("both usernames must be provided")
    }
    if username1 == username2 {
        return nil, errors.New("users must be different")
    }

    user1, err := service.userRepository.GetByUsername(username1)
    if err != nil {
        return nil, errors.New("user does not exist: " + username1)
    }
    user2, err := service.userRepository.GetByUsername(username2)
    if err != nil {
        return nil, errors.New("user does not exist: " + username2)
    }

    existingChat, err := service.chatRepository.FindChatByUsers(user1.UserId, user2.UserId)
    if err != nil {
        return nil, err
    }
    if existingChat != nil {
        return existingChat, nil
    }

    chat := &model.Chat{}
    if err := service.chatRepository.Put(chat); err != nil {
        return nil, err
    }

    createdChat, err := service.chatRepository.GetByChatId(chat.ChatId)
    if err != nil {
        return nil, errors.New("could not retrieve created chat")
    }

    if err := service.chatRepository.AddUserToChat(user1, createdChat); err != nil {
        return nil, err
    }
    if err := service.chatRepository.AddUserToChat(user2, createdChat); err != nil {
        return nil, err
    }

    return createdChat, nil
}

func (service *ChatService) ListUsersFromChat(chatId int32) ([]*model.User, error) {
	chat, err := service.chatRepository.GetByChatId(chatId)
	if err != nil {
		return nil, errors.New("chat does not exist")
	}

	users, err := service.chatRepository.ListUsersFromChat(chat)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (service *ChatService) ListChatsByUser(username string) ([]*model.Chat, error) {
	user, err := service.userRepository.GetByUsername(username)
	if err != nil {
		return nil, errors.New("user does not exist")
	}

	chats, err := service.chatRepository.ListChatsByUser(user)
	if err != nil {
		return nil, err
	}

	return chats, nil
}