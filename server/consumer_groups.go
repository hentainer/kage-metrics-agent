package server

import (
	"net/http"

	"github.com/go-zoo/bone"
	"github.com/msales/kage/store"
)

type consumerGroup struct {
	Group      string              `json:"group"`
	Topic      string              `json:"topic"`
	TotalLag   int64               `json:"total_lag"`
	Partitions []consumerPartition `json:"partitions"`
}

type consumerPartition struct {
	Partition int   `json:"partition"`
	Offset    int64 `json:"offset"`
	Lag       int6