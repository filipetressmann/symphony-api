package user_handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	base_handlers "symphony-api/internal/handlers/base"
	request_model "symphony-api/internal/handlers/model"
	"symphony-api/internal/persistence/repository"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)
type MockPostgreConnection struct {
    mock.Mock
}

func (m *MockPostgreConnection) Put(data map[string]any, table string) error {
    args := m.Called(data, table)
    return args.Error(0)
}

func (m *MockPostgreConnection) PutReturningId(data map[string]any, table string, idName string) (any, error) {
    args := m.Called(data, table)
    return args.Get(0), args.Error(1)
}

func (m *MockPostgreConnection) Get(constraints map[string]any, table string) ([]map[string]any, error) {
    args := m.Called(constraints, table)
    return args.Get(0).([]map[string]any), args.Error(1)
}

func TestCreateUserHandler(t *testing.T) {
	mockConn := &MockPostgreConnection{}
	userCrud := NewUserHandler(mockConn)

	birthDate := time.Date(1990, 5, 10, 0, 0, 0, 0, time.UTC)
	request := request_model.CreateUserRequest{
		BaseUserModel: &request_model.BaseUserModel{
			Username:   "guiwallace",
			Fullname:   "Guilherme Wallace",
			Email:      "guiwallace@example.com",
			Birth_date: birthDate,
			Telephone:  "123456789",
		},
	}

	mockConn.On("Put", mock.Anything, repository.USER_TABLE_NAME).Return(1, nil)

	requestBody, _ := json.Marshal(request)
	req := httptest.NewRequest("POST", "/api/create-user", bytes.NewBuffer(requestBody))
	w := httptest.NewRecorder()

	handler := base_handlers.CreateHandler(userCrud.CreateUserHandler)
	handler.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var response request_model.UserResponse
	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}
	assert.Equal(t, request.Username, response.Username)
	assert.Equal(t, request.Fullname, response.Fullname)
	assert.Equal(t, request.Email, response.Email)
	assert.Equal(t, request.Birth_date, response.Birth_date)
	assert.Equal(t, request.Telephone, response.Telephone)

	mockConn.AssertExpectations(t)
}

func TestGetUserByUsername(t *testing.T) {
	mockConn := &MockPostgreConnection{}
	userCrud := NewUserHandler(mockConn)

	birthDate := time.Date(1990, 5, 10, 0, 0, 0, 0, time.UTC)
	userMap := map[string]any{
		"id":            int32(1),
		"username":      "guiwallace",
		"fullname":      "Guilherme Wallace",
		"email":         "guiwallace@example.com",
		"register_date": time.Now(),
		"birth_date":    birthDate,
		"telephone":     "123456789",
	}

	mockConn.On("Get", mock.Anything, repository.USER_TABLE_NAME).Return([]map[string]any{userMap}, nil)

	request := request_model.GetUserByUsernameRequest{
		Username: "johndoe",
	}

	requestBody, _ := json.Marshal(request)
	req := httptest.NewRequest("POST", "/api/get-user-by-username", bytes.NewBuffer(requestBody))
	w := httptest.NewRecorder()

	handler := base_handlers.CreateHandler(userCrud.GetUserByUsername)
	handler.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var response request_model.UserResponse
	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}
	assert.Equal(t, "guiwallace", response.Username)
	assert.Equal(t, "Guilherme Wallace", response.Fullname)
	assert.Equal(t, "guiwallace@example.com", response.Email)
	assert.Equal(t, birthDate, response.Birth_date)
	assert.Equal(t, "123456789", response.Telephone)

	mockConn.AssertExpectations(t)
}
