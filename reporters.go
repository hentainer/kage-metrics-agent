package kage

import "github.com/msales/kage/store"

// Reporter represents a offset reporter.
type Reporter interface {
	// ReportBrokerOffsets reports a snapshot of the