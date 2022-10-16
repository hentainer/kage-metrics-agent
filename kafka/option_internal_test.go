package kafka

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/inconshreveable/log15.v2"
)

func TestBrokers(t *testing.T) {
	brokers := []string{"127.0.0.1"}
	c := &Monitor{}

	Brokers(brokers)(c)

	assert.Equal(t, brokers, c.brokers)
}

func TestIgnoreGroups(t *testing.T) {
	i := []string{"test"}
	c := &Monitor{}

	IgnoreGroups(i)(c)

	assert.Equal(t, i, c.ignoreGroups)
}

func TestIgnoreTopics