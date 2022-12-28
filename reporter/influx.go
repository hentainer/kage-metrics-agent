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

// D