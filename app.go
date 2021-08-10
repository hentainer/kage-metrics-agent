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