package reporter_test

import (
	"testing"
	"time"

	"github.com/influxdata/influxdb/client/v2"
	"github.com/msales/kage/reporter"
	"github.com/msales/kage/store"
	"github.com/msales/kage/testutil"
	"github.com/msales/kage/testutil/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestInfluxReporter_ReportBrokerOffsets(t *testing.T) {
	c := new