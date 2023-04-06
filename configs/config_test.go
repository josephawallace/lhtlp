package configs

import (
	"os"
	"testing"
)

func TestNewAppConfigEnvVar(t *testing.T) {
	_ = os.Setenv("LHTLP_TEST_PRE_PARAM_T", "1")
	_ = os.Setenv("LHTLP_TEST_PRE_PARAM_LAMBDA", "10")
	_ = os.Setenv("LHTLP_PROBABLY_PRIME_PRECISION", "100")

	appCfg := newAppConfig()
	testCfg := newTestConfig()

	if testCfg.PreParamT != 1 || testCfg.PreParamLambda != 10 || appCfg.ProbablyPrimePrecision != 100 {
		t.Errorf("failed to pull configuration values from environment")
	}
}
