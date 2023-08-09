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
			m.CleanConsumerOffsets()
		}
	}()

	return m, nil
}

// SetState adds a state into the store.
func (m *MemoryStore) SetState(v interface{}) error {
	switch v.(type) {
	case *BrokerPartitionOffset:
		m.addBrokerOffset(v.(*BrokerPartitionOffset))

	case *ConsumerPartitionOffset:
		m.addConsumerOffset(v.(*ConsumerPartitionOffset))

	case *BrokerPartitionMetadata:
		m.addMetadata(v.(*BrokerPartitionMetadata))

	default:
		return errors.New("store: unknown state object")
	}

	return nil
}

// BrokerOffsets returns a snapshot of the current broker offsets.
func (m *MemoryStore) BrokerOffsets() BrokerOffsets {
	m.state.brokerLock.RLock()
	defer m.state.brokerLock.RUnlock()

	snapshot := make(BrokerOffsets)
	for topic, partitions := range m.state.broker {
		snapshot[topic] = make([]*BrokerOffset, len(partitions))

		for partition, offset := range partitions {
			if offset == nil {
				continue
			}

			snapshot[topic][partition] = &BrokerOffset{
				OldestOffset: offset.Old