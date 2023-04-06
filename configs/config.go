package configs

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

var App = newAppConfig()
var Test = newTestConfig()

type appConfig struct {
	ProbablyPrimePrecision int `mapstructure:"probably_prime_precision"`
}

type testConfig struct {
	PreParamLambda int `mapstructure:"test_pre_param_lambda"`
	PreParamT      int `mapstructure:"test_pre_param_t"`
}

func newAppConfig() *appConfig {
	envPrefix := "lhtlp"
	viper.SetEnvPrefix(envPrefix)
	viper.AutomaticEnv()

	viper.SetDefault("probably_prime_precision", 20)

	var cfg appConfig
	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatal().Msg("failed to unmarshal configuration")
	}

	return &cfg
}

func newTestConfig() *testConfig {
	envPrefix := "lhtlp"
	viper.SetEnvPrefix(envPrefix)
	viper.AutomaticEnv()

	viper.SetDefault("test_pre_param_lambda", 512)
	viper.SetDefault("test_pre_param_t", 100000)

	var cfg testConfig
	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatal().Msg("failed to unmarshal configuration")
	}

	return &cfg
}
