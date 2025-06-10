package repository

import (
	"errors"
	"symphony-api/internal/persistence/connectors/postgres"
	"symphony-api/internal/persistence/model"
)

const COMMUNITY_TABLE_NAME = "COMMUNITY"
const USER_TO_COMMUNITY_RELATIONSHIP_TABLE = "USER_COMMUNITY"
const JOINED_USERS_AND_USER_COMMUNITY = "USERS u JOIN USER_COMMUNITY uc ON u.id = uc.user_id"

type CommunityRepository struct {
	connection postgres.PostgreConnection
}

func NewCommunityRepository(connection postgres.PostgreConnection) *CommunityRepository {
	return &CommunityRepository{
		connection: connection,
	}
}

func (repository *CommunityRepository) Put(community *model.Community) error {
	return repository.connection.Put(community.ToTableData(), COMMUNITY_TABLE_NAME)
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

func (repository *CommunityRepository) AddUserToCommunity(user *model.User, community *model.Community) error {
	return repository.connection.Put(
		map[string]any {
			"community_id": community.Id,
			"user_id": user.UserId,
		}, 
		USER_TO_COMMUNITY_RELATIONSHIP_TABLE,
	)
}

func (repository *CommunityRepository) ListUsersFromCommunity(community *model.Community) ([]*model.User, error) {
	constraint := map[string]any {
		"uc.community_id": community.Id,
	}

	users, err := repository.connection.Get(
		constraint,
		JOINED_USERS_AND_USER_COMMUNITY,
	)

	if err != nil {
		return nil, err
	}

	return model.MapArrayToUsers(users), nil
}
