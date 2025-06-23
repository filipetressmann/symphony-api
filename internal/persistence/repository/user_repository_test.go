package repository

import (
	"testing"
	"time"

	"symphony-api/internal/persistence/model"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUserRepository_Put(t *testing.T) {
	mockConn := new(MockPostgreConnection)
	mockNeo4j := &MockNeo4jConn{}
	repo := NewUserRepository(mockConn, mockNeo4j)

	user := &model.User{
		Username:     "john",
		Fullname:     "John Doe",
		Email:        "john@example.com",
		Register_date: time.Now(),
		Birth_date:    time.Date(1990, 5, 10, 0, 0, 0, 0, time.UTC),
		Telephone:    "123456789",
	}

	mockConn.On("Put", mock.Anything, USER_TABLE_NAME).Return(nil)

	err := repo.Put(user)

	assert.NoError(t, err)

	mockConn.AssertExpectations(t)
}

func getFetchTestData() (*model.User, map[string]any) {
	user := &model.User{
		UserId:       1,
		Username:     "john",
		Fullname:     "John Doe",
		Email:        "john@example.com",
		Register_date: time.Now(),
		Birth_date:    time.Date(1990, 5, 10, 0, 0, 0, 0, time.UTC),
		Telephone:    "123456789",
	}
	userMap := user.ToMap()
	userMap["id"] = int32(1)
	userMap["register_date"] =  time.Now()

	return user, userMap
}

func TestUserRepository_GetById(t *testing.T) {
	mockConn := new(MockPostgreConnection)
	mockNeo4j := &MockNeo4jConn{}
	repo := NewUserRepository(mockConn, mockNeo4j)

	user, userMap := getFetchTestData()

	mockConn.On("Get", mock.Anything, USER_TABLE_NAME).Return([]map[string]any{userMap}, nil)

	result, _ := repo.GetById(1)

	assert.Equal(t, user, result)
	mockConn.AssertExpectations(t)
}

func TestUserRepository_GetByUsername(t *testing.T) {
	mockConn := new(MockPostgreConnection)
	mockNeo4j := &MockNeo4jConn{}
	repo := NewUserRepository(mockConn, mockNeo4j)

	user, userMap := getFetchTestData()

	mockConn.On("Get",  mock.Anything, USER_TABLE_NAME).Return([]map[string]any{userMap}, nil)

	result, _ := repo.GetByUsername("john")

	assert.Equal(t, user, result)
	mockConn.AssertExpectations(t)
}