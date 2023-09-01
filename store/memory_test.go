
package store_test

import (
	"testing"
	"time"

	"github.com/msales/kage/store"
	"github.com/stretchr/testify/assert"
)

func TestMemoryStore_SetState(t *testing.T) {
	memStore, err := store.New()
	assert.NoError(t, err)

	defer memStore.Close()

	err = memStore.SetState(1)
	assert.Error(t, err)

	err = memStore.SetState(&store.BrokerPartitionOffset{
		Topic:               "test",
		Partition:           0,
		Oldest:              true,
		Offset:              0,
		Timestamp:           time.Now().Unix(),
		TopicPartitionCount: 1,
	})
	assert.NoError(t, err)
}

func TestMemoryStore_BrokerOffsets(t *testing.T) {
	memStore, err := store.New()
	assert.NoError(t, err)

	defer memStore.Close()

	memStore.SetState(&store.BrokerPartitionOffset{
		Topic:               "test",
		Partition:           0,
		Oldest:              true,
		Offset:              0,
		Timestamp:           time.Now().Unix(),
		TopicPartitionCount: 1,
	})
	memStore.SetState(&store.BrokerPartitionOffset{
		Topic:               "test",
		Partition:           0,
		Oldest:              false,
		Offset:              1000,
		Timestamp:           time.Now().Unix(),
		TopicPartitionCount: 1,
	})

	offsets := memStore.BrokerOffsets()

	assert.Contains(t, offsets, "test")
	assert.Len(t, offsets["test"], 1)
	assert.Equal(t, int64(0), offsets["test"][0].OldestOffset)
	assert.Equal(t, int64(1000), offsets["test"][0].NewestOffset)
}

func TestMemoryStore_BrokerOffsetsMissingParition(t *testing.T) {
	memStore, err := store.New()
	assert.NoError(t, err)

	defer memStore.Close()

	memStore.SetState(&store.BrokerPartitionOffset{
		Topic:               "test",
		Partition:           1,
		Oldest:              true,
		Offset:              0,
		Timestamp:           time.Now().Unix(),
		TopicPartitionCount: 2,
	})