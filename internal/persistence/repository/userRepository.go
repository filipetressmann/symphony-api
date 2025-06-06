package repository

import (
	"errors"
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

func (repository *UserRepository) Put(user *model.User) (*model.User, error) {
	id, err := repository.connection.Put(user.ToMap(), TABLE_NAME)
	return model.NewUser(
		id,
		user.Username,
		user.Fullname,
		user.Email,
		user.Birth_date,
		user.Telephone,
	), err
}

func (repository *UserRepository) get(constraint map[string]any) ([]*model.User, error) {
	data, err := repository.connection.Get(constraint, TABLE_NAME)

	if err != nil {
		return nil, err
	}


	users := make([]*model.User, 0)

	for _, user := range data {
		users = append(users, model.MapToUser(user))
	}

	return users, nil
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
		return nil, errors.New("user not found")
	}

	return users[0], err
}
