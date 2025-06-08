package repository

import (
	"errors"
	"log"
	"symphony-api/internal/persistence/connectors/postgres"
	"symphony-api/internal/persistence/model"
)

const USER_TABLE_NAME = "USERS"
const JOINED_COMMUNITY_AND_USER_COMMUNITY = "COMMUNITY c JOIN USER_COMMUNITY uc ON c.id = uc.community_id"

type UserRepository struct {
	connection postgres.PostgreConnection
}

func NewUserRepository(connection postgres.PostgreConnection) *UserRepository {
	return &UserRepository{
		connection: connection,
	}
}

func (repository *UserRepository) Put(user *model.User) error {
	return repository.connection.Put(user.ToMap(), USER_TABLE_NAME)
}

func (repository *UserRepository) get(constraint map[string]any) ([]*model.User, error) {
	data, err := repository.connection.Get(constraint, USER_TABLE_NAME)

	if err != nil {
		return nil, err
	}


	return model.MapArrayToUsers(data), nil
}

func (repository *UserRepository) GetById(userId int64) (*model.User, error) {
	constraint := map[string]any {
		"id": userId,
	}

	users, err := repository.get(constraint)
	return users[0], err
}

func (repository *UserRepository) GetByUsername(username string) (*model.User, error) {
	constraint := map[string]any {
		"username": username,
	}

	users, err := repository.get(constraint)

	if len(users) == 0 {
		log.Println("COuld not find user")
		return nil, errors.New("user not found")
	}

	return users[0], err
}

func (repository *UserRepository) ListUserCommunities(user *model.User) ([]*model.Community, error) {
	constraint := map[string]any {
		"uc.user_id": user.UserId,
	}

	communities, err := repository.connection.Get(
		constraint,
		JOINED_COMMUNITY_AND_USER_COMMUNITY,
	)

	if err != nil {
		return nil, err
	}

	return model.MapArrayToCommunity(communities), nil
}