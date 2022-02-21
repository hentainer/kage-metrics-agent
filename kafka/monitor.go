package kafka

import (
	"fmt"
	"sync"
	"time"

	"github.com/Shopify/sarama"
	"github.com/msales/kage/store"
	"github.com/ryanuber/go-glob"
	"gopkg.in/inconshreveable/log15.v2"
)

// Broker represents a Kafka Broker.
t