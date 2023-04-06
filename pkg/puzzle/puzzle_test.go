package puzzle

import (
	"math/big"
	"testing"

	"github.com/milquellc/lhtlp/configs"
)

func TestGeneratePuzzle(t *testing.T) {
	preParams := NewPreParams(configs.Test.PreParamLambda, configs.Test.PreParamT)
	puzzle := GeneratePuzzle(big.NewInt(574938), preParams)

	if puzzle.u == nil {
		t.Errorf("failed to set puzzle u value")
	}
	if puzzle.v == nil {
		t.Errorf("failed to set puzzle v value")
	}
}

func TestSolvePuzzle(t *testing.T) {
	var tests = []struct {
		name  string
		input string
	}{
		{"solve test vector 1", "35483049850"},
		{"solve test vector 2", "58309453435"},
		{"solve test vector 3", "10000000000"},
	}

	for _, tt := range tests {
		secret, ok := new(big.Int).SetString(tt.input, 10)
		if !ok {
			t.Errorf("failed to set secret value")
		}
		preParams := NewPreParams(configs.Test.PreParamLambda, configs.Test.PreParamT)

		puzzle := GeneratePuzzle(secret, preParams)
		ans := SolvePuzzle(puzzle, preParams)
		if ans.Cmp(secret) != 0 {
			t.Errorf("expected to uncover secret %d from puzzle, got %d", secret, ans)
		}
	}
}

func TestEvaluatePuzzles(t *testing.T) {
	var tests = []struct {
		name  string
		input string
	}{
		{"eval test vector 1", "35483049850"},
		{"eval test vector 2", "58309453435"},
		{"eval test vector 3", "10000000000"},
	}

	preParams := NewPreParams(configs.Test.PreParamLambda, configs.Test.PreParamT)

	want := big.NewInt(0)

	var puzzles []*Puzzle
	for _, tt := range tests {
		secret, ok := new(big.Int).SetString(tt.input, 10)
		if !ok {
			t.Errorf("failed to set secret value")
		}
		want.Add(want, secret)

		puzzles = append(puzzles, GeneratePuzzle(secret, preParams))
	}

	evaluatedPuzzle := EvaluatePuzzles(puzzles, preParams)

	ans := SolvePuzzle(evaluatedPuzzle, preParams)
	if ans.Cmp(want) != 0 {
		t.Errorf("expected to uncover secret %d from evaluated puzzle, got %d", want, ans)
	}
}
