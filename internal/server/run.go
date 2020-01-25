package server

import (
	"fmt"
	"github.com/wavesoftware/go-ensure"
	"github.com/wavesoftware/go-kahle/internal/base"
	"github.com/wavesoftware/go-kahle/internal/config"
	"net/http"
)

// Run is responsible for running a Kahle server with given configuration
func (s *KahleServer) Run(cfg *config.RunConfig) {
	log := base.Logger()
	http.HandleFunc("/webhook", s.webhookHandler)
	addr := fmt.Sprintf("%s:%d", cfg.Bind, cfg.Port)
	log.Infof("Kahle (%v) is listenning on %s", base.Version, addr)
	ensure.NoError(http.ListenAndServe(addr, nil))
}
