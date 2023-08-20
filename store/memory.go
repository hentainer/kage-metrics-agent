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
				OldestOffset: offset.OldestOffset,
				NewestOffset: offset.NewestOffset,
				Timestamp:    offset.Timestamp,
			}
		}
	}

	return snapshot
}

// ConsumerOffsets returns a snapshot of the current consumer group offsets.
func (m *MemoryStore) ConsumerOffsets() ConsumerOffsets {
	m.state.consumerLock.RLock()
	defer m.state.consumerLock.RUnlock()

	snapshot := make(ConsumerOffsets)
	for group, topics := range m.state.consumer {
		snapshot[group] = make(map[string][]*ConsumerOffset)

		for topic, partitions := range topics {
			snapshot[group][topic] = make([]*ConsumerOffset, len(partitions))

			for partition, offset := range partitions {
				if offset == nil {
					continue
				}

				snapshot[group][topic][partition] = &ConsumerOffset{
					Offset:    offset.Offset,
					Lag:       offset.Lag,
					Timestamp: offset.Timestamp,
				}
			}
		}
	}

	return snapshot
}

// BrokerMetadata returns a snapshot of the current broker metadata.
func (m *MemoryStore) BrokerMetadata() BrokerMetadata {
	m.state.metadataLock.RLock()
	defer m.state.metadataLock.RUnlock()

	snapshot := make(BrokerMetadata)
	for topic, partitions := range m.state.metadata {
		snapshot[topic] = make([]*Metadata, len(partitions))

		for partition, metadata := range partitions {
			if metadata == nil {
				continue
			}

			snapshot[topic][partition] = &Metadata{
				Leader:   metadata.Leader,
				Replicas: make([]int32, len(metadata.Replicas)),
				Isr:      make([]int32, len(metadata.Isr)),
			}
			copy(snapshot[topic][partition].Replicas, metadata.Replicas)
			copy(snapshot[topic][partition].Isr, metadata.Isr)
		}
	}

	return snapshot
}

// CleanConsumerOffsets cleans old offsets from the MemoryStore.
func (m *MemoryStore) CleanConsumerOffsets() {
	m.state.consumerLock.Lock()
	defer m.state.consumerLock.Unlock()

	ts := time.Now().Unix() * 1000
	for group, topics := range m.state.consumer {
		for topic, partitions := range topics {
			maxDuration := int64(0)

			for _, offset := range partitions {
				if offset == nil {
					continue
				}

				duration := ts - offset.Timestamp
				if duration > maxDuration {
					maxDuration = duration
				}
			}

			if maxDuration > (24 * int64(time.Hour.Seconds()) * 1000) {
				delete(m.state.consumer[group], topic)
			}
		}

		if len(m.state.consumer[group]) == 0 {
			delete(m.state.consumer, group)
		}
	}
}

// Channel get the offset channel.
func (m *MemoryStore) Channel() chan interface{} {
	return m.stateCh
}

// Close gracefully stops the MemoryStore.
func (m *MemoryStore) Close() {
	m.cleanupTicker.Stop()
	close(m.shutdown)
}

func (m *MemoryStore) addBrokerOffset(o *BrokerPartitionOffset) {
	m.state.brokerLock.Lock()
	defer m.state.brokerLock.Unlock()

	topic, ok := m.state.broker[o.Topic]
	if !ok {
		topic = make([]*BrokerOffset, o.TopicPartitionCount)
		m.state.broker[o.Topic] = topic
	}

	if o.TopicPartitionCount > len(topic) {
		for i := len(topic); i < o.TopicPartitionCount; i++ {
			topic = append(topic, nil)
		}
