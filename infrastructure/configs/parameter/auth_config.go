package parameter

import (
	"github.com/spf13/viper"
	"time"
)

type AuthConf struct {
	TokenTimeout   time.Duration `mapstructure:"token_timeout"`
	CookiePath     string        `mapstructure:"cookie_path"`
	CookieDomain   string        `mapstructure:"cookie_domain"`
	CookieSecure   bool          `mapstructure:"cookie_secure"`
	CookieHttpOnly bool          `mapstructure:"cookie_httpOnly"`
}

var _ parameterLoader = (*AuthConf)(nil)

func (AuthConf) loadDefault() {
	viper.SetDefault("auth", map[string]interface{}{
		"token_timeout":   86400,
		"cookie_path":     "/",
		"cookie_domain":   "localhost",
		"cookie_secure":   false,
		"cookie_httpOnly": true,
	})
}
