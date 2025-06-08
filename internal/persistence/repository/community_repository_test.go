package repository_test

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"symphony-api/internal/persistence/model"
	"symphony-api/internal/persistence/repository"
)

// --- Mock for PostgreConnection ---
type MockPostgreConnection struct {
	mock.Mock
}

func (m *MockPostgreConnection) Put(data map[string]any, table string) (int32, error) {
	args := m.Called(data, table)
	return int32(args.Int(0)), args.Error(1)
}

func (m *MockPostgreConnection) Get(constraints map[string]any, table string) ([]map[string]any, error) {
	args := m.Called(constraints, table)
	return args.Get(0).([]map[string]any), args.Error(1)
}

// --- Tests ---

func TestPut_Success(t *testing.T) {
	mockConn := new(MockPostgreConnection)
	repo := repository.NewCommunityRepository(mockConn)

	input := &model.Community{
		CommunityName: "TestCommunity",
		Description:   "A test community",
	}

	// Expected table data map (simplified)
	tableData := input.ToTableData()
	mockConn.On("Put", tableData, "COMMUNITY").Return(123, nil)

	result, err := repo.Put(input)

	assert.NoError(t, err)
	assert.Equal(t, int32(123), result.Id)
	assert.Equal(t, input.CommunityName, result.CommunityName)
	assert.Equal(t, input.Description, result.Description)
	assert.WithinDuration(t, time.Now(), result.CreatedAt, time.Second)

	mockConn.AssertExpectations(t)
}

func TestGetByName_Success(t *testing.T) {
	mockConn := new(MockPostgreConnection)
	repo := repository.NewCommunityRepository(mockConn)

	constraint := map[string]any{
		"community_name": "TestCommunity",
	}

	dbResult := []map[string]any{
		{
			"id":             int32(123),
			"community_name": "TestCommunity",
			"description":    "Test description",
			"created_at":     time.Now(),
		},
	}

	mockConn.On("Get", constraint, "COMMUNITY").Return(dbResult, nil)

	result, err := repo.GetByName("TestCommunity")

	assert.NoError(t, err)
	assert.Equal(t, int32(123), result.Id)
	assert.Equal(t, "TestCommunity", result.CommunityName)
	assert.Equal(t, "Test description", result.Description)

	mockConn.AssertExpectations(t)
}

func TestGetByName_NotFound(t *testing.T) {
	mockConn := new(MockPostgreConnection)
	repo := repository.NewCommunityRepository(mockConn)

	mockConn.On("Get", mock.Anything, "COMMUNITY").Return([]map[string]any{}, nil)

	result, err := repo.GetByName("NonExistent")

	assert.Error(t, err)
	assert.Nil(t, result)
	mockConn.AssertExpectations(t)
}

func TestPut_Failure(t *testing.T) {
	mockConn := new(MockPostgreConnection)
	repo := repository.NewCommunityRepository(mockConn)

	input := &model.Community{
		CommunityName: "ErrorCommunity",
		Description:   "This will fail",
	}

	mockConn.On("Put", input.ToTableData(), "COMMUNITY").Return(0, errors.New("db error"))

	result, err := repo.Put(input)

	assert.Error(t, err)
	assert.Empty(t, result)
	mockConn.AssertExpectations(t)
}
