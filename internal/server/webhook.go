package server

import (
	"fmt"
	"github.com/wavesoftware/go-ensure"
	"github.com/wavesoftware/go-kahle/internal/base"
	"net/http"
)

func (s *KahleServer) webhookHandler(writer http.ResponseWriter, req *http.Request) {
	log := base.Logger()
	log.Debugf("Received webhook: %v", req)

	writer.Header().Add("X-Server", fmt.Sprintf("Kahle/%s", base.Version))
	writer.Header().Add("Content-Type", "application/json")

	_, err := writer.Write([]byte("{\"status\": \"Received!\"}"))
	ensure.NoError(err)
}
