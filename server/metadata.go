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

// MetadataHandler handles requests for topic metadata.
func (s *Server) MetadataHandler(w http.ResponseWriter, r *http.Request) {
	metadata := s.Store.BrokerMetadata()

	topics := []topicMetadata{}
	for topic, partitions := range metadata {
		bt := t