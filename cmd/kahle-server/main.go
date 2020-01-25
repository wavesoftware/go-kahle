package main

import (
	"github.com/wavesoftware/go-kahle/internal/config"
	"github.com/wavesoftware/go-kahle/internal/server"
)

func main() {
	srv := &server.KahleServer{}
	cfg := config.NewRunConfig()

	srv.Run(cfg)
}
