package repository

import "github.com/stretchr/testify/mock"

type MockPostgreConnection struct {
    mock.Mock
}

func (m *MockPostgreConnection) Put(data map[string]any, table string) error {
    args := m.Called(data, table)
    return args.Error(0)
}

func (m *MockPostgreConnection) PutReturningId(data map[string]any, tableName string, idName string) (any, error) {
    args := m.Called(data, tableName, idName)
    return args.Get(0), args.Error(1)
}

func (m *MockPostgreConnection) Get(constraints map[string]any, table string) ([]map[string]any, error) {
    args := m.Called(constraints, table)
    return args.Get(0).([]map[string]any), args.Error(1)
}

func (m *MockPostgreConnection) GetChatWithLimit(chat_id int32, limit int32, tableName string) ([]map[string]any, error) {
    args := m.Called(chat_id, limit, tableName)
    return args.Get(0).([]map[string]any), args.Error(1)
}