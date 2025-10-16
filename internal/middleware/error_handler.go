package middleware

import (
	"football-management-api/internal/dto"
	"football-management-api/pkg/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ErrorHandler is a middleware for handling errors
func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		// Check if there are any errors
		if len(c.Errors) > 0 {
			err := c.Errors.Last()

			logger.Error(err.Error())

			// Return error response
			c.JSON(http.StatusInternalServerError, dto.ErrorResponse(
				"Terjadi kesalahan pada server",
				err.Error(),
			))
		}
	}
}

// Recovery recovers from panics
func Recovery() gin.HandlerFunc {
	return gin.CustomRecovery(func(c *gin.Context, recovered interface{}) {
		logger.Error("Panic recovered: " + recovered.(error).Error())

		c.JSON(http.StatusInternalServerError, dto.ErrorResponse(
			"Terjadi kesalahan yang tidak terduga",
			"Internal server error",
		))

		c.AbortWithStatus(http.StatusInternalServerError)
	})
}
