package reporter

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/inconshreveable/log15.v2"
)

func TestDatabase(t *testing.T) {
	r := &InfluxReporter{}

	Database("test")(r)

	assert.Equal(t, r.database, "test")
}

func TestMetric(t *testing.T) {
	r := &InfluxReporter{}

	Metric("kafka")(r)

	assert.Equal(t, r.metric, "kafka")
}

func TestPolicy(t *testing.T) {
	r := &InfluxReporter{}

	Policy("foobar")(r)

	assert.Eq