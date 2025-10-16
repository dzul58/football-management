package utils

import (
	"net/http"

	"football-management-api/internal/dto"

	"github.com/gin-gonic/gin"
)

// SendSuccess sends a successful JSON response
func SendSuccess(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusOK, dto.SuccessResponse(message, data))
}

// SendCreated sends a created response
func SendCreated(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusCreated, dto.SuccessResponse(message, data))
}

// SendBadRequest sends a bad request error response
func SendBadRequest(c *gin.Context, message string, err string) {
	c.JSON(http.StatusBadRequest, dto.ErrorResponse(message, err))
}

// SendNotFound sends a not found error response
func SendNotFound(c *gin.Context, message string, err string) {
	c.JSON(http.StatusNotFound, dto.ErrorResponse(message, err))
}

// SendInternalError sends an internal server error response
func SendInternalError(c *gin.Context, message string, err string) {
	c.JSON(http.StatusInternalServerError, dto.ErrorResponse(message, err))
}

// SendUnauthorized sends an unauthorized error response
func SendUnauthorized(c *gin.Context, message string, err string) {
	c.JSON(http.StatusUnauthorized, dto.ErrorResponse(message, err))
}

// SendConflict sends a conflict error response
func SendConflict(c *gin.Context, message string, err string) {
	c.JSON(http.StatusConflict, dto.ErrorResponse(message, err))
}

// SendPaginated sends a paginated response
func SendPaginated(c *gin.Context, message string, data interface{}, meta dto.PaginationMeta) {
	c.JSON(http.StatusOK, dto.PaginatedSuccessResponse(message, data, meta))
}
