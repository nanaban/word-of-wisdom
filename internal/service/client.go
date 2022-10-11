package service

import (
	"context"
	"fmt"
	"net"

	"word-of-wisdom/internal/config"
	"word-of-wisdom/internal/pow"

	"go.uber.org/zap"
)

var (
	challengeRequest = []byte("challenge")
)

// Client represents a client
type Client struct {
	conf   *config.ClientConfig
	log    *zap.Logger
	solver pow.Solver
}

// NewClient creates a new client
func NewClient(conf *config.ClientConfig, log *zap.Logger, solver pow.Solver) *Client {
	return &Client{
		conf:   conf,
		log:    log,
		solver: solver,
	}
}

// GetQuote returns a quote from the server
func (c *Client) GetQuote(ctx context.Context) ([]byte, error) {
	var dialer net.Dialer
	conn, err := dialer.DialContext(ctx, "tcp", c.conf.ServerAddr)
	if err != nil {
		return nil, fmt.Errorf("failed to dial: %w", err)
	}
	defer func() {
		if err := conn.Close(); err != nil {
			c.log.Error("failed to close connection", zap.Error(err))
		}
	}()

	// 1. Send challenge request
	if err := WriteMessage(conn, challengeRequest); err != nil {
		return nil, fmt.Errorf("send challenge request err: %w", err)
	}

	// 2. Receive challenge
	challenge, err := ReadMessage(conn)
	if err != nil {
		return nil, fmt.Errorf("receive challenge err: %w", err)
	}

	// 3. Send solution
	solution := c.solver.Solve(challenge)
	if err := WriteMessage(conn, solution); err != nil {
		return nil, fmt.Errorf("send solution err: %w", err)
	}

	// 4. Receive quote
	quote, err := ReadMessage(conn)
	if err != nil {
		return nil, fmt.Errorf("receive quote err: %w", err)
	}

	c.log.Debug("quote", zap.ByteString("quote", quote))

	return quote, nil
}
