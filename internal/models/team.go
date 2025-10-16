package models

import (
	"database/sql"
	"time"
)

// Team represents a football team entity
type Team struct {
	ID          int            `json:"id" db:"id"`
	Name        string         `json:"name" db:"name"`
	LogoURL     sql.NullString `json:"logo_url" db:"logo_url"`
	FoundedYear int            `json:"founded_year" db:"founded_year"`
	HomeAddress string         `json:"home_address" db:"home_address"`
	HomeCity    string         `json:"home_city" db:"home_city"`
	DeletedAt   sql.NullTime   `json:"-" db:"deleted_at"`
	CreatedAt   time.Time      `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at" db:"updated_at"`
}

// TableName returns the table name for Team model
func (Team) TableName() string {
	return "teams"
}
