package middleware

import (
	"fmt"
	"football-management-api/pkg/logger"
	"time"

	"github.com/gin-gonic/gin"
)

// Logger is a middleware for logging HTTP requests
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer
		startTime := time.Now()

		// Process request
		c.Next()

		// Calculate latency
		latency := time.Since(startTime)

		// Get request details
		statusCode := c.Writer.Status()
		method := c.Request.Method
		path := c.Request.URL.Path
		clientIP := c.ClientIP()

		// Log the request
		logMessage := fmt.Sprintf(
			"%s | %d | %v | %s | %s",
			clientIP,
			statusCode,
			latency,
			method,
			path,
		)

		if statusCode >= 500 {
			logger.Error(logMessage)
		} else if statusCode >= 400 {
			logger.Warn(logMessage)
		} else {
			logger.Info(logMessage)
		}
	}
}
