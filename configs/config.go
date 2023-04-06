package configs

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

var App = newAppConfig()

type config struct {
	ProbablyPrimePrecision int `mapstructure:"probably_prime_precision"`
	PreParamLambda         int `mapstructure:"pre_param_lambda"`
	PreParamT              int `mapstructure:"pre_param_t"`
}

func newAppConfig() *config {
	envPrefix := "lhtlp"
	viper.SetEnvPrefix(envPrefix)
	viper.AutomaticEnv()

	viper.SetDefault("probably_prime_precision", 20)
	viper.SetDefault("pre_param_lambda", 1024)
	viper.SetDefault("pre_param_t", 10000)

	var cfg config
	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatal().Msg("failed to unmarshal configuration")
	}

	if cfg.PreParamLambda < 10 {
		log.Fatal().Msgf("lambda must be greater than 10, received %d", cfg.PreParamLambda)
	}

	return &cfg
}
