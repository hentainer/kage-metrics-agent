package mocks

import (
	"github.com/msales/kage/kafka"
	"github.com/stretchr/testify/mock"
)

// MockMonitor represents a mock Kafka client.
type MockMonitor struct {
	mock.Mock
}

// Brokers returns a list of Kafka brokers.
func (m *MockMonito