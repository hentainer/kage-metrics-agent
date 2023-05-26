package server

import (
	"net/http"

	"github.com/go-zoo/bone"
	"github.com/msales/kage/store"
)

type consumerGroup struct {
	Group      string              `json:"group"`
	Topic      string            