package repository

//go:generate mockgen -source=repository.go -destination=repository_mock.go -package=repository Repository

// Repository is an interface of repository of quotes.
type Repository interface {
	GetRandomQuote() (string, error)
}
