package repository

import (
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/stretchr/testify/mock"
)

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

type MockNeo4jConn struct {
    mock.Mock
}

func (m *MockNeo4jConn) Execute(query string, data map[string]any) (error) {
	return nil
}

func (m *MockNeo4jConn) ExecuteReturning(query string, data map[string]any) ([]*neo4j.Record, error) {
	return nil, nil
}