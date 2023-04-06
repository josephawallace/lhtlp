package puzzle

import (
	"crypto/rand"
	"math/big"

	"github.com/rs/zerolog/log"
)

// Puzzle represents an encoded value that can be decoded (or "opened") within a pre-specified amount of time, via
// repeated squaring.
type Puzzle struct {
	u *big.Int
	v *big.Int
}

// GeneratePuzzle encodes a secret value in a time-lock puzzle.
func GeneratePuzzle(secret *big.Int, preParams *PreParams) *Puzzle {
	N2 := new(big.Int).Exp(preParams.n, big.NewInt(2), nil)

	r, err := rand.Int(rand.Reader, N2)
	if err != nil {
		log.Fatal().Err(err)
	}

	u := new(big.Int).Exp(preParams.g, r, preParams.n)

	rN := new(big.Int).Mul(r, preParams.n)

	v := new(big.Int).Exp(preParams.h, rN, N2)
	onePlusN := new(big.Int).Add(big.NewInt(1), preParams.n)
	onePlusNS := new(big.Int).Exp(onePlusN, secret, N2)
	v.Mul(v, onePlusNS)
	v.Mod(v, N2)

	return &Puzzle{
		u: u,
		v: v,
	}
}

// SolvePuzzle "force opens" a given puzzle, using repeated squaring.
func SolvePuzzle(puzzle *Puzzle, preParams *PreParams) *big.Int {
	N2 := new(big.Int).Exp(preParams.n, big.NewInt(2), nil)

	s := puzzle.u

	for i := big.NewInt(1); i.Cmp(preParams.t) <= 0; i.Add(i, big.NewInt(1)) {
		s.Mul(s, s)
		s.Mod(s, preParams.n)
	}

	s.Exp(s, preParams.n, N2)
	s.ModInverse(s, N2)
	s.Mul(s, puzzle.v)
	s.Mod(s, N2)
	s.Sub(s, big.NewInt(1))
	s.Div(s, preParams.n)

	return s
}

// EvaluatePuzzles collapses many puzzles to just one by leveraging the homomorphic properties of LHTLPs.
func EvaluatePuzzles(puzzles []*Puzzle, preParams *PreParams) *Puzzle {
	N2 := new(big.Int).Exp(preParams.n, big.NewInt(2), nil)

	u := big.NewInt(1)
	v := big.NewInt(1)

	for _, puzzle := range puzzles {
		u.Mul(u, puzzle.u)
		u.Mod(u, preParams.n)

		v.Mul(v, puzzle.v)
		v.Mod(v, N2)
	}

	return &Puzzle{
		u: u,
		v: v,
	}
}
