package hashcash

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestChallenge(t *testing.T) {
	p := NewPOW(16)
	token := p.Challenge()
	require.Len(t, token, defaultTokenSize)

	nonce := p.Solve(token)
	require.Len(t, nonce, defaultNonceSize)

	err := p.Verify(token, nonce)
	require.NoError(t, err)
}
