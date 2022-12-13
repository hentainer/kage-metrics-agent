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
				OldestOffset: 0,
				NewestOffset: 1000,
				Timestamp:    time.Now().Unix() * 1000,
			},
		},
	}
	r.ReportBrokerOffsets(offsets)

	assert.Equal(t, "test:0 oldest:0 newest:1000 available:1000 \n", buf.String())
}

func TestConsoleReporter_ReportBro