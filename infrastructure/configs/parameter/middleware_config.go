package parameter

import (
	"github.com/spf13/viper"
	"golang.org/x/time/rate"
	"time"
)

type MiddlewareConf struct {
	RateLimit struct {
		Rate      rate.Limit    `mapstructure:"rate"`
		Burst     int           `mapstructure:"burst"`
		ExpiresIn time.Duration `mapstructure:"expiresIn"`
	} `mapstructure:"rate_limit"`
	BodyLimit    string `mapstructure:"body_limit"`
	RecoverSize  int    `mapstructure:"recover_size"`
	LogSkipper   string `mapstructure:"log_skipper"`
	GzipSkipper  string `mapstructure:"gzip_skipper"`
	GzipLevel    int    `mapstructure:"gzip_level"`
	AllowOrigins string `mapstructure:"allow_origins"`
}

var _ parameterLoader = (*MiddlewareConf)(nil)

func (MiddlewareConf) loadDefault() {
	viper.SetDefault("middleware", map[string]interface{}{
		"rate_limit": map[string]interface{}{
			"rate":       10,
			"burst":      10,
			"expires_in": 180,
		},
		"body_limit":    "1M",
		"recover_size":  1,
		"log_skipper":   "localhost",
		"gzip_skipper":  "/auth, /docs",
		"gzip_level":    5,
		"allow_origins": "http://localhost:3000",
	})
}
