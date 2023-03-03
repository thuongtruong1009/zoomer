package middlewares

import (
	"github.com/labstack/echo/v4"
	"time"
)

func LoggerMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		start := time.Now()
		if err := next(c); err != nil {
			c.Error(err)
		}
		stop := time.Now()
		latency := stop.Sub(start)
		req := c.Request()
		res := c.Response()
		remoteIP := c.RealIP()
		method := req.Method
		path := req.URL.Path
		status := res.Status
		size := res.Size

		c.Logger().Infof("%s %s %s %s %d %s", remoteIP, method, path, status, size, latency)
		return nil
	}
}
