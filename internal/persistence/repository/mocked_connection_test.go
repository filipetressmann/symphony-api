package repository

import "github.com/stretchr/testify/mock"

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