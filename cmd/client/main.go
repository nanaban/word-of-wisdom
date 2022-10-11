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
	"word-of-wisdom/internal/service"

	"go.uber.org/zap"
)

var (
	flagAddr       = flag.String("a", ":3000", "TCP address of server")
	flagKeepAlive  = flag.Duration("K", 15*time.Second, "keep-alive period for network connections")
	flagComplexity = flag.Uint64("c", 20, "complexity of PoW algorithm (1-24)")
	flagCount      = flag.Int("n", 100, "number of quotes to get")
	flagDebug      = flag.Bool("d", false, "debug log level")
)

func initConfig() *config.ClientConfig {
	return &config.ClientConfig{
		ServerAddr: *flagAddr,
		KeepAlive:  *flagKeepAlive,
	}
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
		powHC  = initPoW(logger)
	)

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		<-ctx.Done()
		cancel()
	}()

	client := service.NewClient(conf, logger, powHC)
	for i := 0; i < *flagCount; i++ {
		if ctx.Err() != nil {
			break
		}

		q, err := client.GetQuote(ctx)
		if err != nil {
			logger.Error("failed to get quote", zap.Error(err))
		} else {
			logger.Info("quote", zap.ByteString("quote", q))
		}
	}

	logger.Info("done")
}
