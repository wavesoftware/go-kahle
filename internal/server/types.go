package server

import (
	"github.com/wavesoftware/go-kahle/internal/config"
)

// KahleServer is a server part of Kahle
type KahleServer struct {
}

// Server is a generic server implementation
type Server interface {
	Run(cfg *config.RunConfig)
}
