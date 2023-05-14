package kage

import "github.com/msales/kage/store"

// Reporter represents a offset reporter.
type Reporter interface {
	// ReportBrokerOffsets reports a snapshot of the broker offsets.
	ReportBrokerOffsets(o *store.BrokerOffsets)

	// ReportBrokerMetadata reports a snapshot of the broker metadata.
	ReportBrokerMetadata(o *store.BrokerMetadata)

	// ReportConsumerOffsets reports a snapshot of the consumer group offsets.
	ReportConsumerOffsets(o *store.ConsumerOffsets)
}

// Reporters represents a set of reporters.
type Reporters map[string]Reporter

// Add adds a Reporter to the set.
func (rs *Reporters) Add(key string, r Reporter) 