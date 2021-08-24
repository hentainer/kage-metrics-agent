
package kage_test

import (
	"testing"

	"github.com/msales/kage"
	"github.com/msales/kage/store"
	"github.com/msales/kage/testutil/mocks"
	"github.com/stretchr/testify/assert"
)

func TestNewApplication(t *testing.T) {
	app := kage.NewApplication()

	assert.IsType(t, &kage.Application{}, app)
}

func TestApplication_Close(t *testing.T) {
	store := new(mocks.MockStore)
	store.On("Close").Return()

	monitor := new(mocks.MockMonitor)
	monitor.On("Close").Return()

	app := &kage.Application{
		Store:   store,
		Monitor: monitor,
	}

	app.Close()

	store.AssertExpectations(t)
	monitor.AssertExpectations(t)
}

func TestApplication_IsHealthy(t *testing.T) {
	reporters := &kage.Reporters{}

	monitor := new(mocks.MockMonitor)
	monitor.On("IsHealthy").Return(true).Once()
	monitor.On("IsHealthy").Return(false).Once()

	app := &kage.Application{
		Reporters: reporters,
		Monitor:   monitor,
	}

	assert.True(t, app.IsHealthy())
	assert.False(t, app.IsHealthy())

	monitor.AssertExpectations(t)
}

func TestApplication_IsHealthyNoServices(t *testing.T) {
	app := &kage.Application{}

	assert.False(t, app.IsHealthy())
}

func TestApplication_Report(t *testing.T) {