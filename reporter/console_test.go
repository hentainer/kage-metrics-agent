package reporter_test

import (
	"bytes"
	"testing"
	"time"

	"github.com/msales/kage/reporter"
	"github.com/msales/kage/store"
	"github.com/stretchr/testify/assert"
)

func TestConsoleReporter_ReportBrokerOffsets(t *testing.T) {
	buf := bytes.NewBuffer([]byte{})
	r := reporter.NewConsoleReporter(buf)

	offsets := &store.BrokerOffsets{
		"test": []*store.BrokerOffset{
			{
				OldestOffs