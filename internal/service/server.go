package service

import (
	"context"
	"errors"
	"fmt"
	"net"
	"sync"
	"time"

	"word-of-wisdom/internal/config"
	"word-of-wisdom/internal/pow"
	"word-of-wisdom/internal/repository"

	"go.uber.org/zap"
)

// Server represents the server.
type Server struct {
	conf     *config.ServerConfig
	log      *zap.Logger
	verifier pow.Verifier
	repo     repository.Repository
	listener net.Listener
	wg       sync.WaitGroup
	cancel   context.CancelFunc
}

// NewServer creates a new server instance.
func NewServer(conf *config.ServerConfig, log *zap.Logger, repo repository.Repository, verifier pow.Verifier) *Server {
	return &Server{
		conf:     conf,
		log:      log,
		verifier: verifier,
		repo:     repo,
	}
}

func (s *Server) handleConnection(conn net.Conn) error {
	defer func() {
		if err := conn.Close(); err != nil {
			s.log.Error("failed to close connection", zap.Error(err))
		}
	}()

	// Set deadline timeout
	if err := conn.SetDeadline(time.Now().Add(s.conf.Deadline)); err != nil {
		return fmt.Errorf("failed to set deadline timeout: %w", err)
	}

	// 1. Receive challenge request
	if _, err := ReadMessage(conn); err != nil {
		return fmt.Errorf("")
	}

	// 2. Send challenge
	challenge := s.verifier.Challenge()
	if err := WriteMessage(conn, challenge); err != nil {
		return fmt.Errorf("send challenge err: %w", err)
	}

	// 3. Receive solution
	solution, err := ReadMessage(conn)
	if err != nil {
		return fmt.Errorf("receive proof err: %w", err)
	}

	// 4. Verify solution
	if err = s.verifier.Verify(challenge, solution); err != nil {
		return fmt.Errorf("invalid solution: %w", err)
	}

	// 5. Send result
	quote, err := s.repo.GetRandomQuote()
	if err != nil {
		return fmt.Errorf("get random quote err: %w", err)
	}

	if err = WriteMessage(conn, []byte(quote)); err != nil {
		return fmt.Errorf("send quote err: %w", err)
	}

	return nil
}

func (s *Server) serve(ctx context.Context) {
	defer s.wg.Done()

	go func() {
		<-ctx.Done()
		err := s.listener.Close()
		if err != nil && !errors.Is(err, net.ErrClosed) {
			s.log.Error("failed to close listener", zap.Error(err))
		}
	}()

	for {
		conn, err := s.listener.Accept()
		if errors.Is(err, net.ErrClosed) {
			s.log.Debug("listener closed")
			return
		} else if err != nil {
			s.log.Error("failed to accept connection", zap.Error(err))
			continue
		}

		s.wg.Add(1)
		go func(conn net.Conn) {
			defer s.wg.Done()

			if err := s.handleConnection(conn); err != nil {
				s.log.Error("handle connection error", zap.Error(err))
			}
		}(conn)
	}
}

// Run starts the server.
func (s *Server) Run(ctx context.Context) (err error) {
	ctx, s.cancel = context.WithCancel(ctx)
	defer s.cancel()

	lc := net.ListenConfig{
		KeepAlive: s.conf.KeepAlive,
	}
	if s.listener, err = lc.Listen(ctx, "tcp", s.conf.Addr); err != nil {
		return fmt.Errorf("failed to listen: %w", err)
	}

	s.log.Info("starting server", zap.String("addr", s.listener.Addr().String()))

	s.wg.Add(1)
	go s.serve(ctx)
	s.wg.Wait()

	s.log.Info("server stopped")

	return nil
}

// Stop stops the server.
func (s *Server) Stop() {
	s.cancel()
}
