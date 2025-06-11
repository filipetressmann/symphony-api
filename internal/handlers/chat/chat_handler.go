package chat_handlers

import (
	"log"
	"encoding/json"
    "net/http"
    "strconv"
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
    server.AddRoute("/api/chat/get_by_id", handler.GetChatById)
    server.AddRoute("/api/chat/list_users", handler.ListUsersFromChat)
    server.AddRoute("/api/chat/list_chats", handler.ListChatsFromUser)
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
func (handler *ChatHandler) GetChatById(w http.ResponseWriter, r *http.Request) {
    chatIdStr := r.URL.Query().Get("chat_id")
    chatId, err := strconv.Atoi(chatIdStr)
    if err != nil {
        http.Error(w, "invalid chat_id", http.StatusBadRequest)
        return
    }
    chat, err := handler.chatService.GetChatById(int32(chatId))
    if err != nil {
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    }
    resp := request_model.NewBaseChatData(chat.ChatId, chat.CreatedAt)
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(resp)
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
func (handler *ChatHandler) ListUsersFromChat(w http.ResponseWriter, r *http.Request) {
    chatIdStr := r.URL.Query().Get("chat_id")
    chatId, err := strconv.Atoi(chatIdStr)
    if err != nil {
        http.Error(w, "invalid chat_id", http.StatusBadRequest)
        return
    }
    users, err := handler.chatService.ListUsersFromChat(int32(chatId))
    if err != nil {
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    }
    if len(users) < 2 {
        http.Error(w, "not enough users in chat", http.StatusBadRequest)
        return
    }
    resp := request_model.ListUsersFromChatResponse{
        Username1: users[0].Username,
        Username2: users[1].Username,
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(resp)
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
func (handler *ChatHandler) ListChatsFromUser(w http.ResponseWriter, r *http.Request) {
    username := r.URL.Query().Get("username")
    if username == "" {
        http.Error(w, "username required", http.StatusBadRequest)
        return
    }
    chats, err := handler.chatService.ListChatsByUser(username)
    if err != nil {
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    }
    chatIds := make([]int32, len(chats))
    for i, chat := range chats {
        chatIds[i] = chat.ChatId
    }
    resp := request_model.ListChatsFromUserResponse{
        ChatIds: chatIds,
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(resp)
}