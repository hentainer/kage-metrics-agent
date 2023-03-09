package reporter

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/inconshreveable/log15.v2"
)

func TestDatabase(t *testing.T) {
	r := &InfluxReporter{}

	Database("test")(r)

	assert.Equal(