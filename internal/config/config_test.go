package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewServerConfig(t *testing.T) {
	c, err := NewServerConfig()
	require.NoError(t, err)
	assert.NotEmpty(t, c.Addr)
	assert.NotEmpty(t, c.KeepAlive)
	assert.NotEmpty(t, c.Deadline)
}

func TestNewClientConfig(t *testing.T) {
	c, err := NewClientConfig()
	require.NoError(t, err)
	assert.NotEmpty(t, c.ServerAddr)
	assert.NotEmpty(t, c.KeepAlive)
}
