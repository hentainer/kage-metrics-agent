package kage

import "github.com/msales/kage/store"

// Reporter represents a offset reporter.
type Reporter interface {
	// ReportBrokerOffsets reports a snapshot of the broker offsets.
	ReportBrokerOffsets(o *store.BrokerOffsets)

	// ReportBrokerMetadata reports a snapshot of the broker metadata.
	ReportBrokerMetadata(o *store.BrokerMetadata)

	// ReportConsumerOffsets reports a sn