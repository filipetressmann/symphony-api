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
    server.AddRoute("/api/chat/create", base_handlers.CreatePostMethodHandler(handler.CreateChat))
    server.AddRoute("/api/chat/get_by_id", base_handlers.CreateGetMethodHandler(handler.GetChatById))
    server.AddRoute("/api/chat/list_users", base_handlers.CreateGetMethodHandler(handler.ListUsersFromChat))
    server.AddRoute("/api/chat/list_chats", base_handlers.CreateGetMethodHandler(handler.ListChatsFromUser))
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
func (handler *ChatHandler) CreateChat(request request_model.CreateChatRequest) (*request_model.BaseChatData, error) {
    chat, err := handler.chatService.CreateChat(request.Username1, request.Username2)
    if err != nil {
        log.Printf("Error creating chat: %s", err)
        return nil, err
    }

    return request_model.NewBaseChatData(chat.ChatId, chat.CreatedAt), nil
}

// GetChatById retrieves a chat by its ID
//	@Summary		Get chat by ID
//	@Description	Retrieves a chat by its ID.
//	@Tags			chat
//	@Accept			json
//	@Produce		json
//	@Param			chat_id	query		int32	true	"ID of the chat to retrieve"
//	@Success		200		{object}	request_model.BaseChatData
//	@Failure		400		{object}	map[string]string	"Invalid Input"
//	@Failure		404		{object}	map[string]string	"Chat Not Found"
//	@Failure		500		{object}	map[string]string	"Internal Server Error"
//  @Router			/api/chat/get_by_id [get]
func (handler *ChatHandler) GetChatById(request request_model.GetChatByIdRequest) (*request_model.BaseChatData, error) {
    
    chat, err := handler.chatService.GetChatById(request.ChatId)
    if err != nil {
        log.Printf("Error getting chat by id: %s", err)
        return nil, errors.New("chat does not exist")
    }

    return request_model.NewBaseChatData(chat.ChatId, chat.CreatedAt), nil
}

// ListUsersFromChat retrieves the usernames of the first two users in a chat.
//	@Summary		List users from chat
//	@Description	Retrieves the usernames of the first two users in a chat.
//	@Tags			chat
//	@Accept			json
//	@Produce		json
//	@Param			chat_id	query		int32	true	"ID of the chat to list users from"
//	@Success		200		{object}	request_model.ListUsersFromChatResponse
//	@Failure		400		{object}	map[string]string	"Invalid Input"
//	@Failure		404		{object}	map[string]string	"Chat Not Found"
//	@Failure		500		{object}	map[string]string	"Internal Server Error"
//  @Router			/api/chat/list_users [get]
func (handler *ChatHandler) ListUsersFromChat(request request_model.ListUsersFromChatRequest) (*request_model.ListUsersFromChatResponse, error) {
    
    users, err := handler.chatService.ListUsersFromChat(request.ChatId)
    if err != nil {
        log.Printf("Error listing users of chat: %s", err)
        return nil, errors.New("could not find any user")
    }
   
    return &request_model.ListUsersFromChatResponse{
        Username1: users[0].Username,
        Username2: users[1].Username,
    }, nil
}

// ListChatsFromUser retrieves all chat IDs associated with a specific user.
//	@Summary		List chats from user
//	@Description	Retrieves all chat IDs associated with a specific user.
//	@Tags			chat
//	@Accept			json
//	@Produce		json
//	@Param			username	query		string	true	"Username of the user to list chats from"
//	@Success		200		{object}	request_model.ListChatsFromUserResponse
//	@Failure		400		{object}	map[string]string	"Invalid Input"
//	@Failure		404		{object}	map[string]string	"User Not Found"
//	@Failure		500		{object}	map[string]string	"Internal Server Error"
//  @Router			/api/chat/list_chats [get]
func (handler *ChatHandler) ListChatsFromUser(request request_model.ListChatsFromUserRequest) (*request_model.ListChatsFromUserResponse, error) {
    
    chats, err := handler.chatService.ListChatsByUser(request.Username)
    if err != nil {
        log.Printf("Error listing chats of a user: %s", err)
        return nil, errors.New("error listing chats of user")
    }
    chatIds := make([]int32, len(chats))
    for i, chat := range chats {
        chatIds[i] = chat.ChatId
    }
    return &request_model.ListChatsFromUserResponse{
        ChatIds: chatIds,
    }, nil
}