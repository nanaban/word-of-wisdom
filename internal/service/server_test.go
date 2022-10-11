package service

import (
	"context"
	"testing"
	"time"

	"word-of-wisdom/internal/config"
	"word-of-wisdom/internal/pow"
	"word-of-wisdom/internal/repository"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

func TestServerRun(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	server := NewServer(&config.ServerConfig{}, zap.NewNop(), repository.NewMockRepository(ctrl), pow.NewMockPOW(ctrl))
	defer server.Stop()

	err := server.Run(ctx)
	require.NoError(t, err)
}
