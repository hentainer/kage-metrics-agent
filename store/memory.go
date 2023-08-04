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
			case v := <-m.stateCh:
				go m.SetState(v)

			case <-m.shutdown:
				return
			}
		}
	}()

	// Start cleanup task
	m.cleanupTicker = time.NewTicker(1 * time.Hour)
	go func() {
		for range m.cleanupTicker.C {
			m.Clean