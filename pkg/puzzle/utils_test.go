package puzzle

import (
	"math/big"
	"testing"

	"github.com/milquellc/lhtlp/configs"
)

func TestGcd(t *testing.T) {
	var tests = []struct {
		name  string
		input []string
		want  string
	}{
		{"should have no non-trivial common denominators", []string{"5", "7"}, "1"},
		{"should have 2 as the greatest common denominator", []string{"144", "10"}, "2"},
		{"should have 2929 as the greatest common denominator", []string{"371983", "8787"}, "2929"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a, ok := new(big.Int).SetString(tt.input[0], 10)
			if !ok {
				t.Errorf("failed to set gcd input a")
			}
			b, ok := new(big.Int).SetString(tt.input[1], 10)
			if !ok {
				t.Errorf("failed to set gcd input b")
			}
			want, ok := new(big.Int).SetString(tt.want, 10)
			if !ok {
				t.Errorf("failed to set want value")
			}

			ans := gcd(a, b)
			if ans.Cmp(want) != 0 {
				t.Errorf("expected gcd %d, got %d", want, ans)
			}
		})
	}
}

func TestGenerateStrongPrime(t *testing.T) {
	p, err := generateStrongPrime(configs.App.PreParamLambda)
	if err != nil {
		t.Error(err)
	}

	if !p.ProbablyPrime(configs.App.ProbablyPrimePrecision) {
		t.Errorf("expected p to be a prime number, but is more likely a composite number")
	}

	pDash := new(big.Int)
	pDash.Sub(p, big.NewInt(1))
	pDash.Div(pDash, big.NewInt(2))
	if !pDash.ProbablyPrime(configs.App.ProbablyPrimePrecision) {
		t.Errorf("expected a pDash to be a prime number, but is more likely a composite number")
	}
}

func TestTotient(t *testing.T) {
	var tests = []struct {
		name  string
		input []string
		want  string
	}{
		{"totient test vector 1", []string{"29784356073242484683", "31488512254751271359"}, "937865061212171434056864722158340338156"},
		{"totient test vector 2", []string{"36048548036317893227", "27899008170650103899"}, "1005718736205305661425012164373370394948"},
		{"totient test vector 3", []string{"29687894650868717267", "30992286366278895443"}, "920095732631642694679664061941874101572"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p, ok := new(big.Int).SetString(tt.input[0], 10)
			if !ok {
				t.Errorf("failed to set totient input p")
			}
			q, ok := new(big.Int).SetString(tt.input[1], 10)
			if !ok {
				t.Errorf("failed to set totient input q")
			}
			want, ok := new(big.Int).SetString(tt.want, 10)
			if !ok {
				t.Errorf("failed to set want value")
			}

			ans := totient(p, q)
			if ans.Cmp(want) != 0 {
				t.Errorf("expected totient %d, got %d", want, ans)
			}
		})
	}
}
