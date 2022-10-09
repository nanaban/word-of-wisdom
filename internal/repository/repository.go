package repository

// Repository is an interface of repository of quotes.
type Repository interface {
	GetRandomQuote() (string, error)
}
