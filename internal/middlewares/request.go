package middlewares

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/thuongtruong1009/zoomer/pkg/constants"
	"github.com/thuongtruong1009/zoomer/pkg/helpers"
	"time"
)

func (mwm *MiddlewareManager) RequestMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		start := time.Now()

		defer func() {
			stop := time.Now()
			latency := stop.Sub(start)
			req := c.Request()
			res := c.Response()
			protocol := req.Proto
			host := req.Host
			address := req.RemoteAddr
			id := res.Header().Get(echo.HeaderXRequestID)
			remoteIP := c.RealIP()
			agent := req.UserAgent()
			method := req.Method
			path := req.URL.Path
			status := res.Status
			size := res.Size
			uri := req.RequestURI
			bytesIn := req.ContentLength
			bytesOut := res.Size

			c.Logger().Infof(fmt.Sprintf("%s %s %s %s %s %s %d %d %s",
				protocol, host, address, remoteIP, method, path, status, size, latency))

			helpers.WriteLog(constants.RequestLogPath,
				fmt.Sprintf("[Id: %s] Time: ", id), fmt.Sprintf("Start: %s Stop: %s Remote_IP: %s Agent: %s Host: %s Method: %s Uri: %s Status: %d Latency: %s Bytes_in: %d Bytes_out: %d\n", start, stop.Format(time.RFC3339), remoteIP, agent, host, method, uri, status, latency, bytesIn, bytesOut))
		}()

		if err := next(c); err != nil {
			c.Error(err)
		}
		return nil
	}
}
