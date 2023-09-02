
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
	memStore.SetState(&store.BrokerPartitionOffset{
		Topic:               "test",
		Partition:           1,
		Oldest:              false,
		Offset:              1000,
		Timestamp:           time.Now().Unix(),
		TopicPartitionCount: 2,
	})

	offsets := memStore.BrokerOffsets()

	assert.Contains(t, offsets, "test")
	assert.Len(t, offsets["test"], 2)
	assert.Nil(t, offsets["test"][0])
	assert.Equal(t, int64(0), offsets["test"][1].OldestOffset)
	assert.Equal(t, int64(1000), offsets["test"][1].NewestOffset)
}

func TestMemoryStore_BrokerOffsetsIncreasePartitions(t *testing.T) {
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
	memStore.SetState(&store.BrokerPartitionOffset{
		Topic:               "test",
		Partition:           1,
		Oldest:              true,
		Offset:              0,
		Timestamp:           time.Now().Unix(),
		TopicPartitionCount: 2,
	})
	memStore.SetState(&store.BrokerPartitionOffset{
		Topic:               "test",
		Partition:           1,
		Oldest:              false,
		Offset:              1000,
		Timestamp:           time.Now().Unix(),
		TopicPartitionCount: 2,
	})

	offsets := memStore.BrokerOffsets()

	assert.Contains(t, offsets, "test")
	assert.Len(t, offsets["test"], 2)
	assert.Equal(t, int64(0), offsets["test"][1].OldestOffset)
	assert.Equal(t, int64(1000), offsets["test"][1].NewestOffset)
}

func TestMemoryStore_ConsumerOffsets(t *testing.T) {
	memStore, err := store.New()
	assert.NoError(t, err)

	defer memStore.Close()

	memStore.SetState(&store.BrokerPartitionOffset{
		Topic:               "test",
		Partition:           0,
		Oldest:              false,
		Offset:              1000,
		Timestamp:           time.Now().Unix(),
		TopicPartitionCount: 1,
	})
	memStore.SetState(&store.ConsumerPartitionOffset{
		Group:     "foo",
		Topic:     "test",
		Partition: 0,
		Offset:    500,
		Timestamp: time.Now().Unix(),
	})

	offsets := memStore.ConsumerOffsets()

	assert.Contains(t, offsets, "foo")
	assert.Contains(t, offsets["foo"], "test")
	assert.Len(t, offsets["foo"]["test"], 1)
	assert.Equal(t, int64(500), offsets["foo"]["test"][0].Lag)
}

func TestMemoryStore_ConsumerOffsetsZeroOffset(t *testing.T) {
	memStore, err := store.New()
	assert.NoError(t, err)

	defer memStore.Close()

	memStore.SetState(&store.BrokerPartitionOffset{
		Topic:               "test",
		Partition:           0,
		Oldest:              false,
		Offset:              1000,
		Timestamp:           time.Now().Unix(),
		TopicPartitionCount: 1,
	})
	memStore.SetState(&store.ConsumerPartitionOffset{
		Group:     "foo",
		Topic:     "test",
		Partition: 0,
		Offset:    0,
		Timestamp: time.Now().Unix(),
	})

	offsets := memStore.ConsumerOffsets()

	assert.Contains(t, offsets, "foo")
	assert.Contains(t, offsets["foo"], "test")
	assert.Len(t, offsets["foo"]["test"], 1)
	assert.Equal(t, int64(0), offsets["foo"]["test"][0].Lag)
}

func TestMemoryStore_ConsumerOffsetsMissingPartition(t *testing.T) {
	memStore, err := store.New()
	assert.NoError(t, err)

	defer memStore.Close()

	memStore.SetState(&store.BrokerPartitionOffset{
		Topic:               "test",
		Partition:           1,
		Oldest:              false,
		Offset:              1000,
		Timestamp:           time.Now().Unix(),
		TopicPartitionCount: 2,
	})
	memStore.SetState(&store.ConsumerPartitionOffset{
		Group:     "foo",
		Topic:     "test",
		Partition: 1,
		Offset:    0,
		Timestamp: time.Now().Unix(),
	})

	offsets := memStore.ConsumerOffsets()

	assert.Contains(t, offsets, "foo")
	assert.Contains(t, offsets["foo"], "test")
	assert.Len(t, offsets["foo"]["test"], 2)
	assert.Nil(t, offsets["foo"]["test"][0])
	assert.Equal(t, int64(0), offsets["foo"]["test"][1].Lag)
}

func TestMemoryStore_ConsumerOffsetsNoBrokerPartition(t *testing.T) {
	memStore, err := store.New()
	assert.NoError(t, err)

	defer memStore.Close()

	memStore.SetState(&store.ConsumerPartitionOffset{
		Group:     "foo",
		Topic:     "test",
		Partition: 0,
		Offset:    500,
		Timestamp: time.Now().Unix(),
	})

	offsets := memStore.ConsumerOffsets()

	assert.Len(t, offsets, 0)
}

func TestMemoryStore_ConsumerOffsetsBrokerPartitionNil(t *testing.T) {
	memStore, err := store.New()
	assert.NoError(t, err)

	defer memStore.Close()

	memStore.SetState(&store.BrokerPartitionOffset{
		Topic:               "test",
		Partition:           0,
		Oldest:              false,
		Offset:              1000,
		Timestamp:           time.Now().Unix(),
		TopicPartitionCount: 2,
	})
	memStore.SetState(&store.ConsumerPartitionOffset{
		Group:     "foo",
		Topic:     "test",
		Partition: 1,
		Offset:    500,
		Timestamp: time.Now().Unix(),
	})

	offsets := memStore.ConsumerOffsets()

	assert.Len(t, offsets, 0)
}

func TestMemoryStore_ConsumerOffsetsIncreasePartitions(t *testing.T) {
	memStore, err := store.New()
	assert.NoError(t, err)

	defer memStore.Close()

	memStore.SetState(&store.BrokerPartitionOffset{
		Topic:               "test",
		Partition:           0,
		Oldest:              false,
		Offset:              1000,
		Timestamp:           time.Now().Unix(),
		TopicPartitionCount: 1,
	})
	memStore.SetState(&store.ConsumerPartitionOffset{
		Group:     "foo",
		Topic:     "test",
		Partition: 0,
		Offset:    500,
		Timestamp: time.Now().Unix(),
	})
	memStore.SetState(&store.BrokerPartitionOffset{
		Topic:               "test",
		Partition:           1,
		Oldest:              false,
		Offset:              1000,
		Timestamp:           time.Now().Unix(),
		TopicPartitionCount: 2,
	})
	memStore.SetState(&store.ConsumerPartitionOffset{
		Group:     "foo",
		Topic:     "test",
		Partition: 1,
		Offset:    500,
		Timestamp: time.Now().Unix(),
	})

	offsets := memStore.ConsumerOffsets()

	assert.Contains(t, offsets, "foo")
	assert.Contains(t, offsets["foo"], "test")
	assert.Len(t, offsets["foo"]["test"], 2)
	assert.Equal(t, int64(500), offsets["foo"]["test"][1].Lag)
}

func TestMemoryStore_ConsumerOffsetsIncreasePartitionsBeforeBroker(t *testing.T) {
	memStore, err := store.New()
	assert.NoError(t, err)

	defer memStore.Close()

	memStore.SetState(&store.BrokerPartitionOffset{
		Topic:               "test",
		Partition:           0,
		Oldest:              false,
		Offset:              1000,
		Timestamp:           time.Now().Unix(),