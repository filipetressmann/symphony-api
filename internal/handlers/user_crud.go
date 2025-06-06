package handlers

import (
	"log"
	"net/http"
	request_model "symphony-api/internal/handlers/model"
	"symphony-api/internal/persistence/connectors/postgres"
	"symphony-api/internal/persistence/model"
	"symphony-api/internal/persistence/repository"
	"symphony-api/internal/server"
)

type UserCrud struct {
	repository repository.UserRepository
}

func NewUserCrud(connection postgres.PostgreConnection) *UserCrud {
	return &UserCrud{
		repository: *repository.NewUserRepository(connection),
	}
}

// CreateUserHandler handles the creation of a new user.
// @Summary Create a new user
// @Description Creates a new user in the system.
// @Tags User
// @Accept json
// @Produce json
// @Param user body model.User true "User data"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string "Invalid Input"
// @Failure 500 {object} map[string]string "Internal Server Error"
// @Router /api/create-user [post]
func (userCrud *UserCrud) AddRoutes(server server.Server) {
	server.AddRoute("/api/create-user", userCrud.CreateUserHandler)
	server.AddRoute("/api/get-user-by-username", userCrud.GetUserByUsername)
}

func (userCrud *UserCrud) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	user, err := model.UserFromRequest(r)

	if err != nil {
		http.Error(w, "Invalid Input", http.StatusBadRequest)
		return
	}

	createdUser, err := userCrud.repository.Put(user)

	// Change this later if necessary. We should check why the creation failed and give a better
	// answer to the requester.
	if err != nil {
		log.Printf("Error creating user: %s", err)
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		return
	}

	mustEncodeAnswer(
		map[string]any{"user": createdUser,}, 
		w, 
		"Error creating user",
	)
}

func (userCrud *UserCrud) GetUserByUsername(w http.ResponseWriter, r *http.Request) {
	request, err := request_model.NewGetUserByUsernameRequest(r)

	if err != nil {
        http.Error(w, "Invalid Input", http.StatusBadRequest)
        return
    }

	user, err := userCrud.repository.GetByUsername(request.Username)

	if err != nil {
		log.Printf("Error fetching user user: %s", err)
		http.Error(w, "Error fetching user", http.StatusInternalServerError)
        return
	}

	mustEncodeAnswer(
		map[string]any{"user": user,}, 
		w, 
		"Error fetching user",
	)
}