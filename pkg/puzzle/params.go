package puzzle

import (
	"math/big"

	"github.com/rs/zerolog/log"
)

// PreParams define the necessary parameters for generating and solving puzzles. This includes the difficulty parameters.
type PreParams struct {
	t *big.Int
	n *big.Int
	g *big.Int
	h *big.Int
}

// NewPreParams generates a PreParams object that can be used for generating and solving puzzles.
func NewPreParams(lambda int, T int) *PreParams {
	p, err := generateStrongPrime(lambda / 2)
	if err != nil {
		log.Fatal().Msg(err.Error())
	}

	q, err := generateStrongPrime(lambda / 2)
	if err != nil {
		log.Fatal().Msg(err.Error())
	}

	N := new(big.Int)
	N.Mul(p, q)

	g, err := generateRandomGroupElement(N)
	if err != nil {
		log.Fatal().Msg(err.Error())
	}

	g.Exp(g, big.NewInt(2), N)
	g.ModInverse(g, N)

	t1 := totient(p, q)
	t2 := new(big.Int).Exp( // 2^t mod (t1/2)
		big.NewInt(2),
		big.NewInt(int64(T)),
		new(big.Int).Div(t1, big.NewInt(2)),
	)

	h := new(big.Int).Exp(g, t2, N)

	return &PreParams{
		t: big.NewInt(int64(T)),
		n: N,
		g: g,
		h: h,
	}
}
