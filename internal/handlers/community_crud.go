package handlers

import (
	"errors"
	"log"
	request_model "symphony-api/internal/handlers/model"
	"symphony-api/internal/persistence/connectors/postgres"
	"symphony-api/internal/persistence/repository"
	"symphony-api/internal/server"
)

type CommunityCrud struct {
	repository repository.CommunityRepository
}

func NewCommunityCrud(connection postgres.PostgreConnection) *CommunityCrud {
	return &CommunityCrud{
		repository: *repository.NewCommunityRepository(connection),
	}
}

func (communityCrud *CommunityCrud) AddRoutes(server server.Server) {
	server.AddRoute("/api/create-community", createHandler(communityCrud.CreateCommunity))
	server.AddRoute("/api/get-community-by-name", createHandler(communityCrud.GetCommunityByName))
}

func (communityCrud *CommunityCrud) CreateCommunity(request request_model.CreateCommunityRequest) (*request_model.CommunityDataResponse, error) {
	createdCommunity, err := communityCrud.repository.Put(request.ToCommunity())

	if err != nil {
		log.Printf("Error creating community: %s", err)
		return nil, errors.New("error creating community")
	}

	return request_model.NewCommunityDataResponse(createdCommunity), nil
}

func (communityCrud *CommunityCrud) GetCommunityByName(request request_model.GetCommunityByNameRequest) (*request_model.CommunityDataResponse, error) {
	community, err := communityCrud.repository.GetByName(request.CommunityName)

	if err != nil {
		log.Printf("Error fetching community: %s", err)
		return nil, errors.New("error fetching community")
	}

	
	return request_model.NewCommunityDataResponse(community), nil
}