package kage

import (
	"github.com/msales/kage/kafka"
	"github.com/msales/kage/store"
)

// Store represents an offset store.
type Store interface {
	// SetState adds a state into the store.
	SetState(interface{}) error

	// BrokerOffsets returns a snapshot of the current broker offsets.
	BrokerOffsets() store.BrokerOffsets

	// ConsumerOffsets returns a snapshot of the current consumer group offsets.
	ConsumerOffsets() store.ConsumerOffsets

	// BrokerMetadata returns a snapshot of the current broker metadata.
	BrokerMe