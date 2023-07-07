package server

import (
	"net/http"
)

type topicMetadata struct {
	Topic      string              `json:"topic"`
	Partitions []partitionMetadata `json:"partitions"`
}

type partitionMetadata struct {
	Partition int     `json:"partition"`
	Leader    int32   `json:"leader"`
	Replicas  []int32 `json:"replicas"`
	Isr       []int32 `json:"isr"`
}

// MetadataHandler handl