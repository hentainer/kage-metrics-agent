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

// IsHealthy checks the health of t