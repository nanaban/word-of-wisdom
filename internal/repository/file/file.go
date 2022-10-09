package file

import (
	"bufio"
	"bytes"
	_ "embed"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strings"
	"time"
)

//go:embed data/quotes.txt
var embedQuotes []byte

// Repository is a repository of quotes.
type Repository struct {
	rand   *rand.Rand
	quotes []string
}

// NewRepository creates a new repository with the given reader.
func NewRepository(r io.Reader) (*Repository, error) {
	var quotes []string

	s := bufio.NewScanner(r)
	for s.Scan() {
		if q := strings.TrimSpace(s.Text()); q != "" {
			quotes = append(quotes, q)
		}
	}
	if err := s.Err(); err != nil {
		return nil, fmt.Errorf("failed to read quotes: %w", err)
	}

	repo := &Repository{
		rand:   rand.New(rand.NewSource(time.Now().UnixNano())),
		quotes: quotes,
	}

	return repo, nil
}

// NewRepositoryFromFile creates a new repository with the given file.
func NewRepositoryFromFile(name string) (*Repository, error) {
	f, err := os.Open(name)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer func() { _ = f.Close() }()

	return NewRepository(f)
}

// NewEmbedRepository creates a new repository with embedded quotes.
func NewEmbedRepository() (*Repository, error) {
	r := bytes.NewReader(embedQuotes)

	return NewRepository(r)
}

// GetRandomQuote returns a random quote from the repository.
func (r *Repository) GetRandomQuote() (string, error) {
	n := r.rand.Intn(len(r.quotes))

	return r.quotes[n], nil
}
