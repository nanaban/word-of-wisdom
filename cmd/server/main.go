package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"syscall"
	"time"

	"word-of-wisdom/internal/common"
	"word-of-wisdom/internal/config"
	"word-of-wisdom/internal/pow"
	"word-of-wisdom/internal/pow/hashcash"
	"word-of-wisdom/internal/repository"
	"word-of-wisdom/internal/repository/file"
	"word-of-wisdom/internal/service"

	"go.uber.org/zap"
)

var (
	flagAddr       = flag.String("a", ":3000", "TCP address to listen")
	flagKeepAlive  = flag.Duration("K", 15*time.Second, "keep-alive period for network connections")
	flagDeadline   = flag.Duration("D", 10*time.Second, "deadline duration for connections I/O")
	flagComplexity = flag.Uint64("c", 20, "complexity of PoW algorithm (1-24)")
	flagDebug      = flag.Bool("d", false, "debug log level")
)

func initConfig() *config.ServerConfig {
	return &config.ServerConfig{
		Addr:      *flagAddr,
		KeepAlive: *flagKeepAlive,
		Deadline:  *flagDeadline,
	}
}

func initRepository(logger *zap.Logger) repository.Repository {
	repo, err := file.NewEmbedRepository()
	if err != nil {
		logger.Fatal("failed to initialize repository", zap.Error(err))
	}

	return repo
}

func initPoW(logger *zap.Logger) pow.POW {
	p, err := hashcash.NewPOW(*flagComplexity)
	if err != nil {
		logger.Fatal("failed to initialize PoW", zap.Error(err))
	}

	return p
}

func main() {
	flag.Parse()

	var (
		logger = common.MustNewLogger(*flagDebug)
		conf   = initConfig()
		repo   = initRepository(logger)
		powHC  = initPoW(logger)
	)

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		<-ctx.Done()
		cancel()
	}()

	server := service.NewServer(conf, logger, repo, powHC)
	if err := server.Run(ctx); err != nil {
		logger.Fatal("server run error", zap.Error(err))
	}

	logger.Info("done")
}
