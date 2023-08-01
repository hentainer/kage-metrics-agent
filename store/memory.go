package store

import (
	"errors"
	"sync"
	"time"
)

// State represents the state of the store.
type State struct {
	broker     BrokerOffsets
	brokerLock sync.RWMutex

	consumer     ConsumerOffsets
	cons