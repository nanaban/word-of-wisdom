package tests

import (
	"context"
	"sync"
	"testing"
	"time"

	"word-of-wisdom/internal/config"
	"word-of-wisdom/internal/pow"
	"word-of-wisdom/internal/pow/hashcash"
	"word-of-wisdom/internal/repository"
	"word-of-wisdom/internal/repository/file"
	"word-of-wisdom/internal/service"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

var testEnv *env

type env struct {
	logger       *zap.Logger
	serverConfig *config.ServerConfig
	clientConfig *config.ClientConfig
	repo         repository.Repository
	pow          pow.POW
}

func initTestEnv(t *testing.T) {
	t.Helper()

	logger, err := zap.NewDevelopment()
	require.NoError(t, err)

	serverConfig, err := config.NewServerConfig()
	require.NoError(t, err)

	clientConfig, err := config.NewClientConfig()
	require.NoError(t, err)

	repo, err := file.NewEmbedRepository()
	require.NoError(t, err)

	powHC := hashcash.MustNewPOW(20)

	testEnv = &env{
		logger:       logger,
		serverConfig: serverConfig,
		clientConfig: clientConfig,
		repo:         repo,
		pow:          powHC,
	}
}

func newTestServer(t *testing.T) *service.Server {
	t.Helper()

	return service.NewServer(testEnv.serverConfig, testEnv.logger, testEnv.repo, testEnv.pow)
}

func newTestClient(t *testing.T) *service.Client {
	t.Helper()

	return service.NewClient(testEnv.clientConfig, testEnv.logger, testEnv.pow)
}

// Integration test for the service.
func TestServer(t *testing.T) {
	initTestEnv(t)

	var (
		ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
		server      = newTestServer(t)
		wg          = sync.WaitGroup{}
		n           = 10
	)

	wg.Add(1)
	go func() {
		defer wg.Done()
		err := server.Run(ctx)
		require.NoError(t, err)
	}()

	time.Sleep(500 * time.Millisecond)

	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()

			client := newTestClient(t)
			q, err := client.GetQuote(ctx)
			require.NoError(t, err)
			require.NotEmpty(t, q)
		}()
	}

	<-ctx.Done()
	cancel()
	server.Stop()

	wg.Wait()
}
