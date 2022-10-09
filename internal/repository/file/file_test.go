package file

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewEmbedRepository(t *testing.T) {
	r, err := NewEmbedRepository()
	require.NoError(t, err)

	q, err := r.GetRandomQuote()
	require.NoError(t, err)
	require.NotEmpty(t, q)
}

func TestNewRepositoryFromFile(t *testing.T) {
	r, err := NewRepositoryFromFile("./data/quotes.txt")
	require.NoError(t, err)

	q, err := r.GetRandomQuote()
	require.NoError(t, err)
	require.NotEmpty(t, q)
}
