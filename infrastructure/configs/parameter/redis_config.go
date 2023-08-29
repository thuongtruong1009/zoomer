package parameter

import (
	"github.com/spf13/viper"
	"time"
)

type RedisConf struct {
	DB           int           `mapstructure:"db"`
	MinIdleConns int           `mapstructure:"min_idle_conns"`
	PoolSize     int           `mapstructure:"pool_size"`
	PoolTimeout  time.Duration `mapstructure:"pool_timeout"`
	IdleTimeout  time.Duration `mapstructure:"idle_timeout"`
}

var _ parameterLoader = (*RedisConf)(nil)

func (RedisConf) loadDefault() {
	viper.SetDefault("redis", map[string]interface{}{
		"db":             0,
		"min_idle_conns": 200,
		"pool_size":      12000,
		"pool_timeout":   240,
		"idle_timeout":   5,
	})
}
