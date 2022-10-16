package kafka

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/inconshreveable/log15.v2"
)

func TestBrokers(t *testing.T) {
	brokers := []string{"127.0.0.1"}
	c :