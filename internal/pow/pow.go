package pow

//go:generate mockgen -source=pow.go -destination=pow_mock.go -package=pow POW

// Verifier is an interface of proof of work verifier.
type Verifier interface {
	Challenge() []byte
	Verify(challenge, solution []byte) error
}

// Solver is an interface of proof of work solver.
type Solver interface {
	Solve(challenge []byte) []byte
}

// POW is a combined interface of proof of work verifier and solver.
type POW interface {
	Verifier
	Solver
}
