package parameter

import (
	"github.com/spf13/viper"
	"golang.org/x/time/rate"
	"time"
)

type RateLimitConf struct {
	Rate      rate.Limit    `mapstructure:"rate"`
	Burst     int           `mapstructure:"burst"`
	ExpiresIn time.Duration `mapstructure:"expiresIn"`
}

var _ parameterLoader = (*RateLimitConf)(nil)

func (RateLimitConf) loadDefault() {
	viper.SetDefault("rate_limit", map[string]interface{}{
		"rate":       10,
		"burst":      10,
		"expires_in": 180,
	})
}
