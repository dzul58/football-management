package utils

import (
	"database/sql"
	"math"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// NullStringToString converts sql.NullString to string
func NullStringToString(ns sql.NullString) string {
	if ns.Valid {
		return ns.String
	}
	return ""
}

// StringToNullString converts string to sql.NullString
func StringToNullString(s string) sql.NullString {
	if s == "" {
		return sql.NullString{Valid: false}
	}
	return sql.NullString{String: s, Valid: true}
}

// NullInt32ToIntPtr converts sql.NullInt32 to *int
func NullInt32ToIntPtr(ni sql.NullInt32) *int {
	if ni.Valid {
		val := int(ni.Int32)
		return &val
	}
	return nil
}

// IntToNullInt32 converts int to sql.NullInt32
func IntToNullInt32(i int) sql.NullInt32 {
	return sql.NullInt32{Int32: int32(i), Valid: true}
}

// IntPtrToNullInt32 converts *int to sql.NullInt32
func IntPtrToNullInt32(i *int) sql.NullInt32 {
	if i == nil {
		return sql.NullInt32{Valid: false}
	}
	return sql.NullInt32{Int32: int32(*i), Valid: true}
}

// FormatDateTime formats time.Time to string
func FormatDateTime(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

// FormatDate formats time.Time to date string
func FormatDate(t time.Time) string {
	return t.Format("2006-01-02")
}

// FormatTime formats time.Time to time string
func FormatTime(t time.Time) string {
	return t.Format("15:04:05")
}

// ParseDate parses date string to time.Time
func ParseDate(dateStr string) (time.Time, error) {
	return time.Parse("2006-01-02", dateStr)
}

// ParseTime parses time string to time.Time
func ParseTime(timeStr string) (time.Time, error) {
	return time.Parse("15:04:05", timeStr)
}

// GetPaginationParams extracts pagination parameters from request
func GetPaginationParams(c *gin.Context) (page int, limit int) {
	page = 1
	limit = 10

	if pageStr := c.Query("page"); pageStr != "" {
		if p, err := strconv.Atoi(pageStr); err == nil && p > 0 {
			page = p
		}
	}

	if limitStr := c.Query("limit"); limitStr != "" {
		if l, err := strconv.Atoi(limitStr); err == nil && l > 0 {
			limit = l
			// Maximum limit to prevent abuse
			if limit > 100 {
				limit = 100
			}
		}
	}

	return page, limit
}

// CalculateOffset calculates offset for pagination
func CalculateOffset(page, limit int) int {
	return (page - 1) * limit
}

// CalculateTotalPages calculates total pages for pagination
func CalculateTotalPages(total int64, limit int) int {
	return int(math.Ceil(float64(total) / float64(limit)))
}

// Contains checks if a string is in a slice
func Contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

// GetCurrentYear returns current year
func GetCurrentYear() int {
	return time.Now().Year()
}
