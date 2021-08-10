package kage

import (
	"gopkg.in/inconshreveable/log15.v2"
)

// Application represents the kage application.
type Application struct {
	Store     Store
	Reporters *Reporters
	Monitor   Monitor

	Logger log15.Logger
}

// NewApplication creates an instance of Application.
func NewApplication() *Application {
	return &Application{}
}

// Close gracefully shuts down the application
func (a *Application) Close() {
	if a.Store != nil {
		a.Store.Close()
	}

	if a.Monitor != nil {
		a.Monitor.Close()
	}
}

// Collect collects the current state of the Kafka cluster.
func (a *Application) Collect() {
	a.Monitor.Collect()
}

// Report reports the current state of the MemoryStore to the Reporters.
fu