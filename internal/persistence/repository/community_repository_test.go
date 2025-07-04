package repository

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"symphony-api/internal/persistence/model"
)

func TestPut_Success(t *testing.T) {
	mockConn := new(MockPostgreConnection)
	repo := NewCommunityRepository(mockConn)

	input := &model.Community{
		CommunityName: "TestCommunity",
		Description:   "A test community",
	}

	// Expected table data map (simplified)
	tableData := input.ToTableData()
	mockConn.On("Put", tableData, "COMMUNITY").Return(nil)

	err := repo.Put(input)

	assert.NoError(t, err)

	mockConn.AssertExpectations(t)
}

func TestGetByName_Success(t *testing.T) {
	mockConn := new(MockPostgreConnection)
	repo := NewCommunityRepository(mockConn)

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
	repo := NewCommunityRepository(mockConn)

	mockConn.On("Get", mock.Anything, "COMMUNITY").Return([]map[string]any{}, nil)

	result, err := repo.GetByName("NonExistent")

	assert.Error(t, err)
	assert.Nil(t, result)
	mockConn.AssertExpectations(t)
}

func TestPut_Failure(t *testing.T) {
	mockConn := new(MockPostgreConnection)
	repo := NewCommunityRepository(mockConn)

	input := &model.Community{
		CommunityName: "ErrorCommunity",
		Description:   "This will fail",
	}

	mockConn.On("Put", input.ToTableData(), "COMMUNITY").Return(errors.New("db error"))

	err := repo.Put(input)

	assert.Error(t, err)
	mockConn.AssertExpectations(t)
}
