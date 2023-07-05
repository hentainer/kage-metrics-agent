package server

import (
	"net/http"
)

type topicMetadata struct {
	Topic      string              `json:"topic"`
	Partitions 