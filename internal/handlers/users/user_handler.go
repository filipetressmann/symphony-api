package user_handlers

import (
	"errors"
	base_handlers "symphony-api/internal/handlers/base"
	request_model "symphony-api/internal/handlers/model"
	"symphony-api/internal/persistence/connectors/postgres"
	"symphony-api/internal/persistence/repository"
	"symphony-api/internal/persistence/service"
	"symphony-api/internal/server"
)

type UserHandler struct {
	repository *repository.UserRepository
	communityService *service.CommunityService
}

func NewUserHandler(connection postgres.PostgreConnection) *UserHandler {
	userRepository := repository.NewUserRepository(connection)
	return &UserHandler{
		repository: userRepository,
		communityService: service.NewCommunityService(
			repository.NewCommunityRepository(connection),
			userRepository,
		),
	}
}

func (handler *UserHandler) AddRoutes(server server.Server) {
	server.AddRoute(
		"/api/user/create", 
		base_handlers.CreateHandler(handler.CreateUserHandler),
	)
	server.AddRoute(
		"/api/user/get_by_username", 
		base_handlers.CreateHandler(handler.GetUserByUsername),
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
func (handler *UserHandler) CreateUserHandler(request request_model.CreateUserRequest) (*request_model.SuccessCreationResponse, error) {
	
	err := handler.repository.Put(request.ToUser())

	// Change this later if necessary. We should check why the creation failed and give a better
	// answer to the requester.
	if err != nil {
		return nil, errors.New("error creating user")
	}

	return request_model.NewSuccessCreationResponse("Successfully created user"), nil
}

func (handler *UserHandler) GetUserByUsername(request request_model.GetUserByUsernameRequest) (*request_model.UserResponse, error) {
	user, err := handler.repository.GetByUsername(request.Username)

	if err != nil {
        return nil, errors.New("error fetching user")
	}

	return request_model.NewUserResponse(user), nil
}

func (handler *UserHandler) ListUserCommunities(request request_model.ListUserCommunitiesRequest) (*request_model.ListUserCommunitiesResponse, error) {
	communities, err := handler.communityService.ListCommunitiesOfUser(request.Username)

	if err != nil {
		return nil, errors.New("error listing user communities")
	}

	communitiesResponseList := make([]*request_model.CommunityDataResponse, 0)

	for _, community := range communities {
		communitiesResponseList = append(communitiesResponseList, request_model.NewCommunityDataResponse(community))
	}

	return &request_model.ListUserCommunitiesResponse{
		Communities: communitiesResponseList,
	}, nil
}