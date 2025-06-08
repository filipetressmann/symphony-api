package repository_test

import "github.com/stretchr/testify/mock"

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