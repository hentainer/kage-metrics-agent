package reporter

import (
	"fmt"
	"math"
	"time"

	"github.com/influxdata/influxdb/client/v2"
	"github.com/msales/kage/store"
	"gopkg.in/inconshreveable/log15.v2"
)

// InfluxReporterFunc represents a configuration function for InfluxReporter.
type InfluxReporterFunc func(c *InfluxReporter)

// Database configures the database on an InfluxReporter.
func Database(db string) InfluxReporterFunc {
	return func(c *InfluxReporter) {
		c.database = db
	}
}

// Log configures the logger on an InfluxReporter.
func Log(log log15.Logger) InfluxReporterFunc {
	return func(c *InfluxReporter) {
		c.log = log
	}
}

// Metric configures the metric name on an InfluxReporter.
func Metric(metric string) InfluxReporterFunc {
	return func(c *InfluxReporter) {
		c.metric = metric
	}
}

// Policy configures the retention policy name on an InfluxReporter.
func Policy(policy string) InfluxReporterFunc {
	return func(c *InfluxReporter) {
		c.policy = policy
	}
}

// Tags configures the additional tags on an InfluxReporter.
func Tags(tags map[string]string) InfluxReporterFunc {
	return func(c *InfluxReporter) {
		c.tags = tags
	}
}

// InfluxReporter represents an InfluxDB reporter.
type InfluxReporter struct {
	database string

	metric string
	policy string
	tags   map[string]string

	client client.Client

	log log15.Logger
}

// NewInfluxReporter creates and returns a new NewInfluxReporter.
func NewInfluxReporter(client client.Client, opts ...InfluxReporterFunc) *InfluxReporter {
	r := &InfluxReporter{
		client: client,
	}

	for _, o := range opts {
		o(r)
	}

	return r
}

// ReportBrokerOffsets reports a snapshot of the broker offsets.
func (r InfluxReporter) ReportBrokerOffsets(o *store.BrokerOffsets) {
	pts, _ := client.NewBatchPoints(client.BatchPointsConfig{
		Database:        r.database,
		Precision:       "s",
		RetentionPolicy: r.policy,
	})

	for topic, partitions := 