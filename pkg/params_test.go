package pkg

import (
	"math/big"
	"testing"

	"github.com/milquellc/lhtlp/configs"
)

func TestNewPreParams(t *testing.T) {
	preParams := NewPreParams(configs.App.PreParamLambda, configs.App.PreParamT)

	if preParams.t.Int64() != int64(configs.App.PreParamT) {
		t.Errorf("expected t to be assigned to value specified in config, but received %d", preParams.t.Int64())
	}
	if preParams.n.ProbablyPrime(configs.App.ProbablyPrimePrecision) {
		t.Errorf("expected n to be a composite value, but received prime number %d", preParams.t.Int64())
	}
	if gcd(preParams.g, preParams.n).Cmp(big.NewInt(1)) != 0 {
		t.Errorf("expected g to be an element in ring Z/nZ, but g is not coprime with n")
	}
	if gcd(preParams.h, preParams.n).Cmp(big.NewInt(1)) != 0 {
		t.Errorf("expected h to be an element in ring Z/nZ, but h is not coprime with n")
	}
}
