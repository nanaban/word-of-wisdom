package mock

import (
	"bytes"
	"errors"
)

var (
	defaultChallenge = []byte("challenge")
	defaultSolution  = []byte("solution")

	ErrNotEqual = errors.New("not equal")
)

// POW is a mock implementation of the proof of work
type POW struct{}

func New() *POW {
	return &POW{}
}

func (m *POW) Challenge() []byte {
	return defaultChallenge
}

func (m *POW) Verify(challenge, solution []byte) error {
	if bytes.Equal(challenge, defaultChallenge) && bytes.Equal(solution, defaultSolution) {
		return nil
	}

	return ErrNotEqual
}

func (m *POW) Solve(_ []byte) []byte {
	return defaultSolution
}
