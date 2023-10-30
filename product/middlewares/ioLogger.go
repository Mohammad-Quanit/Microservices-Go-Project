package middlewares

import (
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func IOLogger() gin.HandlerFunc {
	logFile, err := os.OpenFile("gin.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic("Failed to open log file: " + err.Error())
	}

	return func(c *gin.Context) {
		start := time.Now()
		c.Next()

		// Logging to file
		end := time.Now()
		latency := end.Sub(start)
		clientIP := c.ClientIP()
		method := c.Request.Method
		statusCode := c.Writer.Status()
		path := c.Request.URL.Path
		bodySize := c.Writer.Size()

		// errorMessage := c.Errors.ByType(gin.ErrorTypePrivate).String()
		if c.Request.URL.RawQuery != "" {
			path += "?" + c.Request.URL.RawQuery
		}

		// IO Log formatting
		logLine := time.Now().Format("2006/01/02 - 15:04:05") + " | Host: " + clientIP + " | Method: " + method + " | Route: " + path + " | Status: " + strconv.Itoa(statusCode) + " | Latency: " + latency.String() + " | Body Size: " + strconv.Itoa(bodySize) + "\n"

		if _, err := logFile.WriteString(logLine); err != nil {
			panic("Failed to write to log file: " + err.Error())
		}
	}
}
