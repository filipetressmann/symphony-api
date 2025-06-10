package chat_handlers

import (
	"errors"
	"log"
	base_handlers "symphony-api/internal/handlers/base"
	request_model "symphony-api/internal/handlers/model"
	"symphony-api/internal/persistence/connectors/postgres"
	"symphony-api/internal/persistence/repository"
	"symphony-api/internal/persistence/service"
	"symphony-api/internal/server"
)

type ChatHandler struct {
	chatRepository   *repository.ChatRepository
	chatService      *service.ChatService
}

func NewChatHandler(connection postgres.PostgreConnection) *ChatHandler {
	chatRepository := repository.NewChatRepository(connection)
	chatService := service.NewChatService(chatRepository, repository.NewUserRepository(connection))

	return &ChatHandler{
		chatRepository: chatRepository,
		chatService:    chatService,
	}
}

func (handler *ChatHandler) AddRoutes(server server.Server) {
	server.AddRoute("/api/chat/create", base_handlers.CreateHandler(handler.CreateChat))
	server.AddRoute("/api/chat/get_by_id", base_handlers.CreateHandler(handler.GetChatById))
	server.AddRoute("/api/chat/list_users", base_handlers.CreateHandler(handler.ListUsersFromChat))
	server.AddRoute("/api/chat/list_chats", base_handlers.CreateHandler(handler.ListChatsFromUser))
}

// CreateChat handles the creation of a new chat between two users.
//	@Summary		Create a new chat
//	@Description	Creates a new chat between two users in the system.
//	@Tags			chat
//	@Accept			json
//	@Produce		json
//	@Param			usernames	body		request_model.CreateChatRequest	true	"Usernames of the users to create a chat between"
//	@Success		200		{object}	request_model.SuccessCreationResponse
//	@Failure		400		{object}	map[string]string	"Invalid Input"
//	@Failure		500		{object}	map[string]string	"Internal Server Error"
//  @Router			/api/chat/create [post]
func (handler *ChatHandler) CreateChat(request request_model.CreateChatRequest) (*request_model.SuccessCreationResponse, error) {
	err := handler.chatService.CreateChat(request.Username1, request.Username2)
	if err != nil {
		log.Printf("Error creating chat: %s", err)
		return nil, errors.New("error creating chat")
	}

	return request_model.NewSuccessCreationResponse("Successfully created chat"), nil
}

// GetChatById returns a chat by its ID.
//	@Summary		Returns chat data
//	@Description	Returns chat data by its ID.
//	@Tags			chat
//	@Accept			json
//	@Produce		json
//	@Param			chat_id	body		request_model.GetChatByIdRequest	true	"Chat ID"
//	@Success		200		{object}	request_model.BaseChatData
//	@Failure		400		{object}	map[string]string	"Invalid Input"
//	@Failure		500		{object}	map[string]string	"Internal Server Error"
//	@Router			/api/chat/get_by_id [get]
func (handler *ChatHandler) GetChatById(request request_model.GetChatByIdRequest) (*request_model.BaseChatData, error) {
	chat, err := handler.chatService.GetChatById(request.ChatId)
	if err != nil {
		log.Printf("Error getting chat by ID: %s", err)
		return nil, errors.New("error getting chat by ID")
	}

	return request_model.NewChatDataResponse(chat.ChatId, chat.CreatedAt), nil
}


// ListUsersFromChat returns the users in a chat by its ID.
//	@Summary		List users from a chat
//	@Description	Returns the users in a chat by its ID.
//	@Tags			chat
//	@Accept			json
//	@Produce		json
//	@Param			chat_id	body		request_model.ListUsersFromChatRequest	true	"Chat ID"
//	@Success		200		{object}	request_model.ListUsersFromChatResponse
//	@Failure		400		{object}	map[string]string	"Invalid Input"
//	@Failure		500		{object}	map[string]string	"Internal Server Error"
//  @Router			/api/chat/list_users [get]
func (handler *ChatHandler) ListUsersFromChat(request request_model.ListUsersFromChatRequest) (*request_model.ListUsersFromChatResponse, error) {
	users, err := handler.chatService.ListUsersFromChat(request.ChatId)
	if err != nil {
		log.Printf("Error listing users from chat: %s", err)
		return nil, errors.New("error listing users from chat")
	}

	if len(users) < 2 {
		return nil, errors.New("not enough users in chat")
	}

	return &request_model.ListUsersFromChatResponse{
		Username1: users[0].Username,
		Username2: users[1].Username,
	}, nil
}

// ListChatsFromUser returns the chat IDs for a user.
//	@Summary		List chats from a user
//	@Description	Returns the chat IDs for a user.
//	@Tags			chat
//	@Accept			json
//	@Produce		json
//	@Param			user_id	body		request_model.ListChatsFromUserRequest	true	"User data"
//	@Success		200		{object}	request_model.ListChatsFromUserResponse
//	@Failure		400		{object}	map[string]string	"Invalid Input"
//	@Failure		500		{object}	map[string]string	"Internal Server Error"
//  @Router			/api/chat/list_chats [get]
func (handler *ChatHandler) ListChatsFromUser(request request_model.ListChatsFromUserRequest) (*request_model.ListChatsFromUserResponse, error) {
	chats, err := handler.chatService.ListChatsByUser(request.Username)
	if err != nil {
		log.Printf("Error listing chats from user: %s", err)
		return nil, errors.New("error listing chats from user")
	}

	chatIds := make([]int32, len(chats))
	for i, chat := range chats {
		chatIds[i] = chat.ChatId
	}

	return &request_model.ListChatsFromUserResponse{
		ChatIds: chatIds,
	}, nil
}