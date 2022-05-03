package kafka

import (
	"fmt"
	"sync"
	"time"

	"github.com/Shopify/sarama"
	"github.com/msales/kage/store"
	"github.com/ryanuber/go-glob"
	"gopkg.in/inconshreveable/log15.v2"
)

// Broker represents a Kafka Broker.
type Broker struct {
	ID        int32
	Connected bool
}

// Monitor represents a Kafka cluster connection.
type Monitor struct {
	brokers []string

	client        sarama.Client
	refreshTicker *time.Ticker
	stateCh       chan interface{}

	ignoreTopics []string
	ignoreGroups []string

	log log15.Logger
}

// New creates and returns a new Monitor for a Kafka cluster.
func New(opts ...MonitorFunc) (*Monitor, error) {
	monitor := &Monitor{}

	for _, o := range opts {
		o(monitor)
	}

	config := sarama.NewConfig()
	config.Version = sarama.V0_10_1_0

	kafka, err := sarama.NewClient(monitor.brokers, config)
	if err != nil {
		return nil, err
	}
	monitor.client = kafka

	monitor.refreshTicker = time.NewTicker(2 * time.Minute)
	go func() {
		for range monitor.refreshTicker.C {
			monitor.refreshMetadata()
		}
	}()

	// Collect initial information
	go monitor.Collect()

	return monitor, nil
}

// Brokers returns a list of Kafka brokers.
func (m *Monitor) Brokers() []Broker {
	brokers := []Broker{}
	for _, b := range m.client.Brokers() {
		connected, _ := b.Connected()
		brokers = append(brokers, Broker{
			ID:        b.ID(),
			Connected: connected,
		})
	}
	return brokers
}

// Collect collects the state of Kafka.
func (m *Monitor) Collect() {
	m.getBrokerOffsets()
	m.getBrokerMetadata()
	m.getConsumerOffsets()
}

// IsHealthy checks the health of the Kafka cluster.
func (m *Monitor) IsHealthy() bool {
	for _, b := range m.client.Brokers() {
		if ok, _ := b.Connected(); ok {
			return true
		}
	}

	return false
}

// Close gracefully stops the Monitor.
func (m *Monitor) Close() {
	// Stop the offset ticker
	m.refreshTicker.Stop()
}

// getTopics gets the topics for the Kafka cluster.
func (m *Monitor) getTopics() map[string]int {
	// If auto create topics is on, trying to fetch metadata for a missing
	// topic will recreate it. To get around this we refresh the metadata
	// before getting topics and partitions.
	m.client.RefreshMetadata()

	topics, _ := m.client.Topics()

	topicMap := make(map[string]int)
	for _, topic := range topics {
		partitions, _ := m.client.Partitions(topic)

		topicMap[topic] = len(partitions)
	}

	return topicMap
}

// refreshMetadata refreshes the broker metadata
func (m *Monitor) refreshMetadata(topics ...string) {
	if err := m.client.RefreshMetadata(topics...); err != nil {
		m.log.Error(fmt.Sprintf("could not refresh topic metadata: %v", err))
	}
}

// getBrokerOffsets gets all broker topic offsets and sends them to the store.
func (m *Monitor) getBrokerOffsets() {
	topicMap := m.getTopics()

	requests := make(map[int32]map[int64]*sarama.OffsetRequest)
	brokers := make(map[int32]*sarama.Broker)

	for topic, partitions := range topicMap {
		if containsString(m.ignoreTopics,