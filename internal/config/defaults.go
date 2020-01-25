package config

import "github.com/wavesoftware/go-kahle/internal/base"

// NewRunConfig creates a new run configuration
func NewRunConfig() *RunConfig {
	return &RunConfig{
		Port: base.EnvI("PORT", 12987),
		Bind: "",
	}
}
