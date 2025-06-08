package repository

import (
	"errors"
	"symphony-api/internal/persistence/connectors/postgres"
	"symphony-api/internal/persistence/model"
	"time"
)

const COMMUNITY_TABLE_NAME = "COMMUNITY"

type CommunityRepository struct {
	connection postgres.PostgreConnection
}

func NewCommunityRepository(connection postgres.PostgreConnection) *CommunityRepository {
	return &CommunityRepository{
		connection: connection,
	}
}

func (repository *CommunityRepository) Put(community *model.Community) (*model.Community, error) {
	id, err := repository.connection.Put(community.ToTableData(), COMMUNITY_TABLE_NAME)
	if err != nil {
		return nil, err
	}

	return model.NewCommunity(
		id,
		community.CommunityName,
		community.Description,
		time.Now(),
	), nil
}

func (repository *CommunityRepository) GetByName(communityName string) (*model.Community, error) {
	constraint := map[string]any {
		"community_name": communityName,
	}

	community, err := repository.connection.Get(constraint, COMMUNITY_TABLE_NAME)

	if len(community) == 0 {
		return nil, errors.New("community not found")
	}

	return model.NewCommunityFromMap(community[0]), err
}
