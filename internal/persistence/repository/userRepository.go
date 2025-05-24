package repository

import (
	"symphony-api/internal/persistence/connectors/postgres"
	"symphony-api/internal/persistence/model"
)

var TABLE_NAME = "USERS"

type UserRepository struct {
	connection postgres.PostgreConnection
}

func NewUserRepository(connection postgres.PostgreConnection) *UserRepository {
	
	return &UserRepository{
		connection: connection,
	}
}

func (repository *UserRepository) Put(user *model.User) int {
	return repository.connection.Put(user.ToMap(), TABLE_NAME)
}

func (repository *UserRepository) get(constraint map[string]any) []*model.User {
	data := repository.connection.Get(constraint, TABLE_NAME)

	users := make([]*model.User, 0)

	for _, user := range data {
		users = append(users, model.MapToUser(user))
	}

	return users
}

func (repository *UserRepository) GetById(userId int64) *model.User {
	constraint := map[string]any {
		"userId": userId,
	}

	users := repository.get(constraint)
	return users[0]
}

func (repository *UserRepository) GetByUsername(username string) *model.User {
	constraint := map[string]any {
		"username": username,
	}

	users := repository.get(constraint)
	return users[0]
}
