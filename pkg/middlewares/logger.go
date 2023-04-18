package middlewares

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"log"
	"os"
	"time"
	"zoomer/pkg/constants"
)

func writeRequestLog(filePath string, logID, logMessage string) {
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal("Failed to open log file: ", err)
	}
	defer file.Close()

	// Create a new logger that writes to the log file
	logger := log.New(file, logID, log.Ldate|log.Ltime|log.Lshortfile)

	// Write some log messages
	logger.Println(logMessage)
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

		c.Logger().Infof("%s %s %s %s %d %s\n", protocol, host, address, remoteIP, method, path, status, size, latency)

		writeRequestLog(constants.RequestLogPath,
			"[Id: "+id+"] ", " Time: "+stop.Format(time.RFC3339)+" Remote_IP: "+remoteIP+" Agent: "+agent+" Host: "+host+" Method: "+method+" Uri: "+uri+" Status: "+fmt.Sprint(status)+" Latency: "+fmt.Sprint(latency)+" Bytes_in: "+fmt.Sprint(bytesIn)+" Bytes_out: "+fmt.Sprint(bytesOut)+"\n")
		return nil
	}
}
