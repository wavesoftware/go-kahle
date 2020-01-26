package server

import (
	"fmt"
	"github.com/wavesoftware/go-ensure"
	"github.com/wavesoftware/go-kahle/internal/base"
	"net/http"
	"strings"
)

func (s *KahleServer) webhookHandler(writer http.ResponseWriter, req *http.Request) {
	log := base.Logger()
	log.Debugf("Received request: %v", PrettyRequest{req: req})

	s.respondWithStatus(writer, "received")
}

func (s *KahleServer) respondWithStatus(writer http.ResponseWriter, status string) {
	headers := writer.Header()
	headers.Add("X-Server", fmt.Sprintf("Kahle/%s", base.Version))
	headers.Add("Content-Type", "application/json")

	_, err := writer.Write([]byte(fmt.Sprintf("{\"status\": \"%s\"}", status)))
	ensure.NoError(err)
}

// PrettyRequest can display your request in easy readable form
type PrettyRequest struct {
	req *http.Request
}

// String presents a request as a easy readable ascii representation
func (p *PrettyRequest) String() string {
	// Create return string
	var request []string
	// Add the request string
	url := fmt.Sprintf("%v %v %v", p.req.Method, p.req.URL, p.req.Proto)
	request = append(request, url) // Add the host
	request = append(request, fmt.Sprintf("Host: %v", p.req.Host))
	// Loop through headers
	for name, headers := range p.req.Header {
		name = strings.ToLower(name)
		for _, h := range headers {
			request = append(request, fmt.Sprintf("%v: %v", name, h))
		}
	}

	// If this is a POST, add post data
	if p.req.Method == "POST" {
		ensure.NoError(p.req.ParseForm())
		request = append(request, "\n")
		request = append(request, p.req.Form.Encode())
	} // Return the request as a string
	return strings.Join(request, "\n")
}
