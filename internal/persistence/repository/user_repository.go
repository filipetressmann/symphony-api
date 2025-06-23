package repository

import (
	"errors"
	"log"
	"symphony-api/internal/persistence/connectors/neo4j"
	"symphony-api/internal/persistence/connectors/postgres"
	"symphony-api/internal/persistence/model"

	neo4jDriver "github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

const USER_TABLE_NAME = "USERS"
const JOINED_COMMUNITY_AND_USER_COMMUNITY = "COMMUNITY c JOIN USER_COMMUNITY uc ON c.id = uc.community_id"

type UserRepository struct {
	connection postgres.PostgreConnection
	neo4jConn neo4j.Neo4jConnection
}

func NewUserRepository(connection postgres.PostgreConnection, neo4jConnection neo4j.Neo4jConnection) *UserRepository {
	return &UserRepository{
		connection: connection,
		neo4jConn: neo4jConnection,
	}
}

func (repository *UserRepository) Put(user *model.User) error {
	repository.neo4jConn.Execute(
		"CREATE (p:User {username:$username})",
		user.ToMap(),
	)
	return repository.connection.Put(user.ToMap(), USER_TABLE_NAME)
}

func (repository *UserRepository) AddFriendship(username1 string, username2 string) error {
	return repository.neo4jConn.Execute(
		`
		MATCH (u1:User {username:$username1}), (u2:User {username:$username2})
		MERGE (u1)-[:FRIENDS_WITH]-(u2)
		`,
		map[string]any{
			"username1": username1,
			"username2": username2,
		},
	)
}

func (repository *UserRepository) ListFriendshipsByUsername(username string) ([]*model.User, error) {
	result, err := repository.neo4jConn.ExecuteReturning(
		`
		MATCH (u:User {username:$username})-[:FRIENDS_WITH]-(friend:User)
		RETURN friend.username AS friend
		`,
		map[string]any{
			"username": username,
		},
	)

	if err != nil {
		return nil, err
	}

	friends := make([]*model.User, 1)

	for _, username := range getFriendUsernames(result) {
		friend, err := repository.GetByUsername(username)
		if err != nil {
			return nil, err
		}

		friends = append(friends, friend)
	}

	return friends, nil
}

func getFriendUsernames(records []*neo4jDriver.Record) []string {
	names := make([]string, 0)

	for _, record := range records {
		friendName, ok := record.Get("friend")

		if ok {
			names = append(names, friendName.(string))
		}
	}

	return names
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
		log.Println("Could not find user")
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