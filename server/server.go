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
func New(app *kage.Application) *Server