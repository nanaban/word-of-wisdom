package hashcash

import (
	"errors"
)

var (
	ErrInvalidChallenge = errors.New("invalid challenge")
	ErrInvalidSolution  = errors.New("invalid solution")
	ErrUnverified       = errors.New("unverified challenge")
)

// POW represents a proof of work algorithm implementation based on hashcash
type POW struct {
	difficulty uint64
}

// NewPOW creates a new POW
func NewPOW(difficulty uint64) *POW {
	return &POW{difficulty: difficulty}
}

// Challenge returns a challenge
func (p *POW) Challenge() []byte {
	return newToken(p.difficulty)
}

// Verify verifies a challenge and solution
func (p *POW) Verify(challenge, solution []byte) error {
	if len(challenge) != defaultTokenSize {
		return ErrInvalidChallenge
	}

	if len(solution) != defaultNonceSize {
		return ErrInvalidSolution
	}

	if !verify(challenge, solution) {
		return ErrUnverified
	}

	return nil
}

// Solve solves a challenge
func (p *POW) Solve(challenge []byte) []byte {
	if len(challenge) != defaultTokenSize {
		return nil
	}

	return solve(challenge)
}
