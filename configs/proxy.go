package configs

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/url"
)

func ProxyConfig(echo *echo.Echo) {
	url1, err := url.Parse("http://localhost:8082")
	if err != nil {
		panic(err)
	}

	url2, err := url.Parse("http://localhost:8083")
	if err != nil {
		panic(err)
	}

	targets := []*middleware.ProxyTarget{
		{
			URL: url1,
		},
		{
			URL: url2,
		},
	}

	echo.Use(middleware.Proxy(middleware.NewRoundRobinBalancer(targets)))
}
