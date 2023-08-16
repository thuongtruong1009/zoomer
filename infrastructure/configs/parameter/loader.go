package parameter

import (
	"github.com/spf13/viper"
	"log"
)

type parameterLoader interface {
	loadDefault()
}

type ParameterConfig struct {
	RateLimitConf `mapstructure:"rate_limit"`
	PostgresConf  `mapstructure:"postgres"`
	ServerConf    `mapstructure:"server"`
	AuthConf      `mapstructure:"auth"`
	OtherConf     `mapstructure:"others"`
}

var pmt *ParameterConfig

func loadDefaultParameter(pmt *ParameterConfig) {
	pmt.RateLimitConf.loadDefault()
	pmt.PostgresConf.loadDefault()
	pmt.ServerConf.loadDefault()
	pmt.AuthConf.loadDefault()
	pmt.OtherConf.loadDefault()
}

func LoadParameterConfigs(path string) *ParameterConfig {
	if pmt == nil {
		pmt = &ParameterConfig{}
		viper.AddConfigPath(path)
		viper.SetConfigName("parameter")
		viper.SetConfigType("yml")

		if err := viper.ReadInConfig(); err != nil {
			log.Println("Error reading config file: ", err)
			log.Println("Using default config")
		}

		loadDefaultParameter(pmt)
		if err := viper.Unmarshal(pmt); err != nil {
			log.Panic(err)
		}
	}

	return pmt
}
