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

// Log configures the logger on an InfluxRepor