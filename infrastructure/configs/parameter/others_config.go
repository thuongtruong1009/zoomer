package parameter

import (
	"github.com/spf13/viper"
	"time"
)

type OtherConf struct {
	CtxTimeout time.Duration `mapstructure:"ctx_timeout"`
	TokenTimeout time.Duration `mapstructure:"token_timeout"`
}

var _ parameterLoader = (*OtherConf)(nil)

func (OtherConf) loadDefault() {
	viper.SetDefault("others", map[string]interface{}{
		"ctx_timeout": 3,
		"token_timeout": 86400,
	})
}
