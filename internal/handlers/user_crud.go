package handlers

import (
	"errors"
	request_model "symphony-api/internal/handlers/model"
	"symphony-api/internal/persistence/connectors/postgres"
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

func (userCrud *UserCrud) AddRoutes(server server.Server) {
	server.AddRoute(
		"/api/create-user", 
		createHandler(userCrud.CreateUserHandler),
	)
	server.AddRoute(
		"/api/get-user-by-username", 
		createHandler(userCrud.GetUserByUsername),
	)
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
func (userCrud *UserCrud) CreateUserHandler(request request_model.CreateUserRequest) (*request_model.UserResponse, error) {
	
	createdUser, err := userCrud.repository.Put(request.ToUser())

	// Change this later if necessary. We should check why the creation failed and give a better
	// answer to the requester.
	if err != nil {
		return nil, errors.New("error creating user")
	}

	return request_model.NewUserResponse(createdUser), nil
}

func (userCrud *UserCrud) GetUserByUsername(request request_model.GetUserByUsernameRequest) (*request_model.UserResponse, error) {
	user, err := userCrud.repository.GetByUsername(request.Username)

	if err != nil {
        return nil, errors.New("error fetching user")
	}

	return request_model.NewUserResponse(user), nil
}