package mocks

import (
	"github.com/msales/kage/store"
	"github.com/stretchr/testify/mock"
)

// MockStore represents a mock offset store.
type MockStore struct {
	mock.Mock
}

// SetState adds a state into the store.
func (m *MockStore) SetState(v interface{}) error {
	args := m.Called(v)
	return args.Error(0)
}

// BrokerOffs