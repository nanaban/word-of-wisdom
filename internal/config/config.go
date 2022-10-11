package config

import (
	"fmt"
	"time"

	"github.com/kelseyhightower/envconfig"
)

// todo timeouts
// ServerConfig represents the server configuration.
type ServerConfig struct {
	Addr      string        `envconfig:"SERVER_ADDR"       default:":3000"`
	KeepAlive time.Duration `envconfig:"SERVER_KEEP_ALIVE" default:"15s"`
	Deadline  time.Duration `envconfig:"SERVER_DEADLINE"   default:"5s"`
}

// NewServerConfig creates a new server config.
func NewServerConfig() (*ServerConfig, error) {
	var config ServerConfig
	if err := envconfig.Process("", &config); err != nil {
		return nil, fmt.Errorf("failed to process server config: %w", err)
	}

	return &config, nil
}

// ClientConfig represents the client configuration.
type ClientConfig struct {
	ServerAddr string        `envconfig:"SERVER_ADDR"       default:":3000"`
	KeepAlive  time.Duration `envconfig:"CLIENT_KEEP_ALIVE" default:"15s"`
}

// NewClientConfig creates a new client config.
func NewClientConfig() (*ClientConfig, error) {
	var config ClientConfig
	if err := envconfig.Process("", &config); err != nil {
		return nil, fmt.Errorf("failed to process client config: %w", err)
	}

	return &config, nil
}
