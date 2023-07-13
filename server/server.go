package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-zoo/bone"
	"github.com/msales/kage"
)

// Server represents an http server.
type Server struct {
	*kage.Application

	mux *bone.Mux
}

// New creates a new instance of Server.
func New(app *kage.Application) *Server {
	s := &Server{
		Application: app,
		mux:         bone.New(),
	}

	s.mux.GetFunc("/brokers", s.BrokersHandler)
	s.mux.GetFunc("/brokers/health", s.BrokersHealthHandler)
	s.mux.GetFunc("/metadata", s.