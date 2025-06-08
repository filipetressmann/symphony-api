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
	server.AddRoute("/api/community/create", base_handlers.CreateHandler(handler.CreateCommunity))
	server.AddRoute("/api/community/get_by_name", base_handlers.CreateHandler(handler.GetCommunityByName))
	server.AddRoute("/api/community/add_user", base_handlers.CreateHandler(handler.AddUserToCommunity))
	server.AddRoute("/api/community/list_users", base_handlers.CreateHandler(handler.ListUsersFromCommunity))
}

func (handler *CommunityHandler) CreateCommunity(request request_model.CreateCommunityRequest) (*request_model.SuccessCreationResponse, error) {
	err := handler.communityRepository.Put(request.ToCommunity())

	if err != nil {
		log.Printf("Error creating community: %s", err)
		return nil, errors.New("error creating community")
	}

	return request_model.NewSuccessCreationResponse("Successfully created community"), nil
}

func (handler *CommunityHandler) GetCommunityByName(request request_model.GetCommunityByNameRequest) (*request_model.CommunityDataResponse, error) {
	community, err := handler.communityRepository.GetByName(request.CommunityName)

	if err != nil {
		log.Printf("Error fetching community: %s", err)
		return nil, errors.New("error fetching community")
	}

	
	return request_model.NewCommunityDataResponse(community), nil
}

func (handler *CommunityHandler) AddUserToCommunity(request request_model.AddUserToCommunityRequest) (*request_model.SuccessCreationResponse, error) {
	err := handler.communityService.AddUserToCommunity(request.Username, request.CommunityName)
	return request_model.NewSuccessCreationResponse("Successfully added user to community"), err
}

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