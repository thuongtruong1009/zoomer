package parameter

import (
	"github.com/spf13/viper"
	"time"
)

type PostgresConf struct {
	MaxOpenConnection     int           `mapstructure:"max_open_conn"`
	MaxIdleConnection     int           `mapstructure:"max_idle_conn"`
	MaxLifetimeConnection time.Duration `mapstructure:"max_lifetime_conn"`
	MaxIdleTimeConnection time.Duration `mapstructure:"max_idle_time_conn"`
	AutoMigrate           bool          `mapstructure:"auto_migrate"`
}

var _ parameterLoader = (*PostgresConf)(nil)

func (PostgresConf) loadDefault() {
	viper.SetDefault("postgres", map[string]interface{}{
		"max_open_conn":      20,
		"max_idle_conn":      20,
		"max_lifetime_conn":  20,
		"max_idle_time_conn": 20,
		"auto_migrate":       true,
	})
}
