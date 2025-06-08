package repository_test

import (
	"testing"
	"time"

	"symphony-api/internal/persistence/model"
	"symphony-api/internal/persistence/repository"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUserRepository_Put(t *testing.T) {
	mockConn := new(MockPostgreConnection)
	repo := repository.NewUserRepository(mockConn)

	user := &model.User{
		UserId:       1,
		Username:     "john",
		Fullname:     "John Doe",
		Email:        "john@example.com",
		Register_date: time.Now(),
		Birth_date:    time.Date(1990, 5, 10, 0, 0, 0, 0, time.UTC),
		Telephone:    "123456789",
	}

	mockConn.On("Put", mock.Anything, "USER").Return(1)

	result, _ := repo.Put(user)

	assert.Equal(t, int32(1), result.UserId)

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
	repo := repository.NewUserRepository(mockConn)

	user, userMap := getFetchTestData()

	mockConn.On("Get", mock.Anything, "USER").Return([]map[string]any{userMap})

	result, _ := repo.GetById(1)

	assert.Equal(t, user, result)
	mockConn.AssertExpectations(t)
}

func TestUserRepository_GetByUsername(t *testing.T) {
	mockConn := new(MockPostgreConnection)
	repo := repository.NewUserRepository(mockConn)

	user, userMap := getFetchTestData()

	mockConn.On("Get",  mock.Anything, "USER").Return([]map[string]any{userMap})

	result, _ := repo.GetByUsername("john")

	assert.Equal(t, user, result)
	mockConn.AssertExpectations(t)
}
