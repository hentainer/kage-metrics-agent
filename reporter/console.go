
package reporter

import (
	"fmt"
	"io"
	"strings"

	"github.com/msales/kage/store"
)

// ConsoleReporter represents a console reporter.
type ConsoleReporter struct {
	w io.Writer
}

// NewConsoleReporter creates and returns a new ConsoleReporter.
func NewConsoleReporter(w io.Writer) *ConsoleReporter {
	return &ConsoleReporter{
		w: w,
	}
}

// ReportBrokerOffsets reports a snapshot of the broker offsets.
func (r ConsoleReporter) ReportBrokerOffsets(o *store.BrokerOffsets) {
	for topic, partitions := range *o {
		for partition, offset := range partitions {
			if offset == nil {
				continue
			}
