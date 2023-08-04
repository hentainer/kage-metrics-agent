package store

import (
	"errors"
	"sync"
	"time"
)

// State represents the state of the store.
type State struct {
	broker     BrokerOffsets
	brokerLock sync.RWMutex

	consumer     ConsumerOffsets
	consumerLock sync.RWMutex

	metadata     BrokerMetadata
	metadataLock sync.RWMutex
}

// MemoryStore represents an in memory data store.
type MemoryStore struct {
	state         *State
	cleanupTicker *time.Ticker
	shutdown      chan struct{}

	stateCh chan interface{}
}

// New creates and returns a new MemoryStore.
func New() (*MemoryStore, error) {
	m := &MemoryStore{
		shutdown: make(chan struct{}),
		stateCh:  make(chan interface{}, 10000),
	}

	// Initialise the cluster offsets
	m.state = &State{
		broker:   make(BrokerOffsets),
		consumer: make(ConsumerOffsets),
		metadata: make(BrokerMetadata),
	}

	// Start the offset reader
	go func() {
		for {
			select {
			c