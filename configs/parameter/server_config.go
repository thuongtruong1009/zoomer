package parameter

import (
	"github.com/spf13/viper"
	"time"
)

type ServerConf struct {
	ShutdownTimeout    time.Duration  `mapstructure:"shutdown_timeout"`
	WriteTimeout	   time.Duration  `mapstructure:"write_timeout"`
	ReadTimeout		   time.Duration  `mapstructure:"read_timeout"`
	StreamMaxConnection int            `mapstructure:"stream_max_connection"`
}

var _ parameterLoader = (*ServerConf)(nil)

func (ServerConf) loadDefault() {
	viper.SetDefault("server", map[string]interface{}{
		"shutdown_timeout": 3,
		"write_timeout": 5,
		"read_timeout": 5,
		"stream_max_connection": 100,
	})
}
