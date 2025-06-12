package community_handlers

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

type CommunityHandler struct {
	communityRepository *repository.CommunityRepository
	communityService *service.CommunityService
}

func NewCommunityHandler(connection postgres.PostgreConnection) *CommunityHandler {
	communityRepository := repository.NewCommunityRepository(connection)
	return &CommunityHandler{
		communityRepository: communityRepository,
		communityService: service.NewCommunityService(
			communityRepository,
			repository.NewUserRepository(connection),
		),
	}
}

func (handler *CommunityHandler) AddRoutes(server server.Server) {
	server.AddRoute("/api/community/create", base_handlers.CreatePostMethodHandler(handler.CreateCommunity))
	server.AddRoute("/api/community/get_by_name", base_handlers.CreateGetMethodHandler(handler.GetCommunityByName))
	server.AddRoute("/api/community/add_user", base_handlers.CreatePostMethodHandler(handler.AddUserToCommunity))
	server.AddRoute("/api/community/list_users", base_handlers.CreateGetMethodHandler(handler.ListUsersFromCommunity))
}

// CreateCommunity handles the creation of a new community.
//	@Summary		Create a new community
//	@Description	Creates a new community in the system.
//	@Tags			community
//	@Accept			json
//	@Produce		json
//	@Param			post	body		request_model.CreateCommunityRequest	true	"Community data"
//	@Success		200		{object}	request_model.SuccessCreationResponse
//	@Failure		400		{object}	map[string]string	"Invalid Input"
//	@Failure		500		{object}	map[string]string	"Internal Server Error"
//	@Router			/api/community/create [post]
func (handler *CommunityHandler) CreateCommunity(request request_model.CreateCommunityRequest) (*request_model.SuccessCreationResponse, error) {
	err := handler.communityRepository.Put(request.ToCommunity())

	if err != nil {
		log.Printf("Error creating community: %s", err)
		return nil, errors.New("error creating community")
	}

	return request_model.NewSuccessCreationResponse("Successfully created community"), nil
}

// GetCommunityByName returns a community by their name.
//	@Summary		Returns community data
//	@Description	Returns the coomunity data based on the name.
//	@Tags			community
//	@Accept			json
//	@Produce		json
//	@Param			post	body		request_model.GetCommunityByNameRequest	true	"Community data"
//	@Success		200		{object}	request_model.CommunityDataResponse
//	@Failure		400		{object}	map[string]string	"Invalid Input"
//	@Failure		500		{object}	map[string]string	"Internal Server Error"
//	@Router			/api/community/get_by_name [get]
func (handler *CommunityHandler) GetCommunityByName(request request_model.GetCommunityByNameRequest) (*request_model.CommunityDataResponse, error) {
	community, err := handler.communityRepository.GetByName(request.CommunityName)

	if err != nil {
		log.Printf("Error fetching community: %s", err)
		return nil, errors.New("error fetching community")
	}

	
	return request_model.NewCommunityDataResponse(community), nil
}

// AddUserToCommunity adds a user to a community
//	@Summary		Add a user to a community
//	@Description	Adds a user to a community based on the username and community name
//	@Tags			community
//	@Accept			json
//	@Produce		json
//	@Param			post	body		request_model.AddUserToCommunityRequest	true	"User and Community data"
//	@Success		200		{object}	request_model.SuccessCreationResponse
//	@Failure		400		{object}	map[string]string	"Invalid Input"
//	@Failure		500		{object}	map[string]string	"Internal Server Error"
//	@Router			/api/community/add_user [post]
func (handler *CommunityHandler) AddUserToCommunity(request request_model.AddUserToCommunityRequest) (*request_model.SuccessCreationResponse, error) {
	err := handler.communityService.AddUserToCommunity(request.Username, request.CommunityName)
	return request_model.NewSuccessCreationResponse("Successfully added user to community"), err
}

// List all users that belongs to a community
//	@Summary		List user of a community
//	@Description	List all user data of a community
//	@Tags			community
//	@Accept			json
//	@Produce		json
//	@Param			community_name	query		string	true	"Community Name"1'
//	@Success		200		{object}	request_model.ListUsersOfCommunityResponse
//	@Failure		400		{object}	map[string]string	"Invalid Input"
//	@Failure		500		{object}	map[string]string	"Internal Server Error"
//	@Router			/api/community/list_users [get]
func (handler *CommunityHandler) ListUsersFromCommunity(request request_model.ListUsersOfCommunityRequest) (*request_model.ListUsersOfCommunityResponse, error) {
	users, err := handler.communityService.ListUsersFromCommunity(request.CommunityName)

	usersResponse := make([]*request_model.UserResponse, 0)

	for _, user := range users {
		usersResponse = append(usersResponse, request_model.NewUserResponse(user))
	}

	return &request_model.ListUsersOfCommunityResponse{
		Users: usersResponse,
	}, err
}