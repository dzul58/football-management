package models

import (
	"database/sql"
	"time"
)

// MatchStatus represents the status of a match
type MatchStatus string

const (
	StatusScheduled MatchStatus = "Scheduled"
	StatusCompleted MatchStatus = "Completed"
	StatusCancelled MatchStatus = "Cancelled"
)

// Match represents a football match entity
type Match struct {
	ID         int           `json:"id" db:"id"`
	MatchDate  string        `json:"match_date" db:"match_date"`
	MatchTime  string        `json:"match_time" db:"match_time"`
	HomeTeamID int           `json:"home_team_id" db:"home_team_id"`
	AwayTeamID int           `json:"away_team_id" db:"away_team_id"`
	HomeScore  sql.NullInt32 `json:"home_score" db:"home_score"`
	AwayScore  sql.NullInt32 `json:"away_score" db:"away_score"`
	Status     MatchStatus   `json:"status" db:"status"`
	DeletedAt  sql.NullTime  `json:"-" db:"deleted_at"`
	CreatedAt  time.Time     `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time     `json:"updated_at" db:"updated_at"`

	// Relations
	HomeTeam *Team  `json:"home_team,omitempty" db:"-"`
	AwayTeam *Team  `json:"away_team,omitempty" db:"-"`
	Goals    []Goal `json:"goals,omitempty" db:"-"`
}

// TableName returns the table name for Match model
func (Match) TableName() string {
	return "matches"
}

// IsValidStatus checks if a status is valid
func IsValidStatus(status string) bool {
	validStatuses := []MatchStatus{
		StatusScheduled,
		StatusCompleted,
		StatusCancelled,
	}

	for _, valid := range validStatuses {
		if MatchStatus(status) == valid {
			return true
		}
	}
	return false
}
