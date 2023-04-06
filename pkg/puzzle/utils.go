package puzzle

import (
	"crypto/rand"
	"math/big"

	"github.com/milquellc/lhtlp/configs"
)

func generateStrongPrime(size int) (*big.Int, error) {
	p := new(big.Int)
	for !p.ProbablyPrime(configs.App.ProbablyPrimePrecision) {
		pDash, err := rand.Prime(rand.Reader, size)
		if err != nil {
			return nil, err
		}

		p = new(big.Int).Mul(pDash, big.NewInt(2))
		p = new(big.Int).Add(p, big.NewInt(1))
	}

	return p, nil
}

func generateRandomGroupElement(N *big.Int) (*big.Int, error) {
	randomElement, err := rand.Int(rand.Reader, N)
	if err != nil {
		return nil, err
	}

	for gcd(randomElement, N).Cmp(big.NewInt(1)) != 0 {
		randomElement, err = rand.Int(rand.Reader, N)
		if err != nil {
			return nil, err
		}
	}

	return randomElement, nil
}

func gcd(a *big.Int, b *big.Int) *big.Int {
	for b.Cmp(big.NewInt(0)) != 0 {
		t := b
		b = new(big.Int).Mod(a, b)
		a = t
	}

	return a
}

func totient(p *big.Int, q *big.Int) *big.Int {
	pMinusOne := new(big.Int).Sub(p, big.NewInt(1))
	qMinusOne := new(big.Int).Sub(q, big.NewInt(1))
	return new(big.Int).Mul(pMinusOne, qMinusOne)
}
