package middlewares

import (
	"github.com/labstack/echo/v4"
	"log"
	"os"
	"time"
)

func writeRequestLog(filePath string) {
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal("Failed to open log file: ", err)
	}
	defer file.Close()

	// Create a new logger that writes to the log file
	logger := log.New(file, "example: ", log.Ldate|log.Ltime|log.Lshortfile)

	// Write some log messages
	logger.Println("This is a new log message.")
	logger.Println("This is another new log message.")
}

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
		protocol := req.Proto
		host := req.Host
		address := req.RemoteAddr
		remoteIP := c.RealIP()
		method := req.Method
		path := req.URL.Path
		status := res.Status
		size := res.Size

		c.Logger().Infof("%s %s %s %s %d %s", protocol, host, address, remoteIP, method, path, status, size, latency)

		writeRequestLog("logs/requests.log")
		return nil
	}
}

// {"time":"2017-11-13T20:26:28.6438003+01:00","id":"3","remote_ip":"::1","host":"localhost:1323","method":"GET","uri":"/?my=param","my":"param","status":200, "latency":0,"latency_human":"0s","bytes_in":0,"bytes_out":13}
