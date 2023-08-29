package parameter

import (
	"github.com/spf13/viper"
	"log"
)

type parameterLoader interface {
	loadDefault()
}

type ParameterConfig struct {
	ServerConf     `mapstructure:"server"`
	MiddlewareConf `mapstructure:"middleware"`
	PostgresConf   `mapstructure:"postgres"`
	RedisConf      `mapstructure:"redis"`
	AuthConf       `mapstructure:"auth"`
	OtherConf      `mapstructure:"others"`
}

var pmt *ParameterConfig

func loadDefaultParameter(pmt *ParameterConfig) {
	pmt.MiddlewareConf.loadDefault()
	pmt.PostgresConf.loadDefault()
	pmt.RedisConf.loadDefault()
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
