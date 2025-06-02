package repository

import (
	"testing"
	"time"

	"symphony-api/internal/persistence/model"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockPostgreConnection struct {
	mock.Mock
}

func (m *MockPostgreConnection) Put(data map[string]any, table string) (int64, error) {
	args := m.Called(data, table)
	return int64(args.Int(0)), nil
}

func (m *MockPostgreConnection) Get(constraint map[string]any, table string) []map[string]any {
	args := m.Called(constraint, table)
	return args.Get(0).([]map[string]any)
}

func TestUserRepository_Put(t *testing.T) {
	mockConn := &MockPostgreConnection{}
	repo := NewUserRepository(mockConn)

	user := &model.User{
		UserId:       1,
		Username:     "john",
		Fullname:     "John Doe",
		Email:        "john@example.com",
		Register_date: time.Now(),
		Birth_date:    time.Date(1990, 5, 10, 0, 0, 0, 0, time.UTC),
		Telephone:    "123456789",
	}

	mockConn.On("Put", mock.Anything, TABLE_NAME).Return(1)

	result, _ := repo.Put(user)

	assert.Equal(t, int64(1), result.UserId)

	mockConn.AssertExpectations(t)
}

func TestUserRepository_GetById(t *testing.T) {
	mockConn := new(MockPostgreConnection)
	repo := NewUserRepository(mockConn)

	user := &model.User{
		UserId:       1,
		Username:     "john",
		Fullname:     "John Doe",
		Email:        "john@example.com",
		Register_date: time.Now(),
		Birth_date:    time.Date(1990, 5, 10, 0, 0, 0, 0, time.UTC),
		Telephone:    "123456789",
	}

	mockConn.On("Get", mock.Anything, TABLE_NAME).Return([]map[string]any{user.ToMap()})

	result := repo.GetById(1)

	assert.Equal(t, user, result)
	mockConn.AssertExpectations(t)
}

func TestUserRepository_GetByUsername(t *testing.T) {
	mockConn := new(MockPostgreConnection)
	repo := NewUserRepository(mockConn)

	user := &model.User{
		UserId:       1,
		Username:     "john",
		Fullname:     "John Doe",
		Email:        "john@example.com",
		Register_date: time.Now(),
		Birth_date:    time.Date(1990, 5, 10, 0, 0, 0, 0, time.UTC),
		Telephone:    "123456789",
	}

	mockConn.On("Get",  mock.Anything, TABLE_NAME).Return([]map[string]any{user.ToMap()})

	result := repo.GetByUsername("john")

	assert.Equal(t, user, result)
	mockConn.AssertExpectations(t)
}
