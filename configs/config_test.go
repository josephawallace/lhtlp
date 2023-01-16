package configs

import (
	"os"
	"testing"
)

func TestNewAppConfigEnvVar(t *testing.T) {
	_ = os.Setenv("LHTLP_PRE_PARAM_T", "1")
	_ = os.Setenv("LHTLP_PRE_PARAM_LAMBDA", "10")
	_ = os.Setenv("LHTLP_PROBABLY_PRIME_PRECISION", "100")
	cfg := newAppConfig()

	if cfg.PreParamT != 1 || cfg.PreParamLambda != 10 || cfg.ProbablyPrimePrecision != 100 {
		t.Errorf("failed to pull configuration values from environment")
	}
}
