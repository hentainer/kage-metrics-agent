package mocks

import (
	"github.com/msales/kage/store"
	"github.com/stretchr/testify/mock"
)

// MockReporter represents a mock Reporter.
type MockReporter struct {
	mock.Mock
}

// ReportBrokerOffsets reports a snapshot of the broker offsets.
func (m *MockReporter) ReportBrokerOffsets(v *sto