package models

import (
	"database/sql"
	"time"
)

// Goal represents a goal scored in a match
type Goal struct {
	ID        int          `json:"id" db:"id"`
	MatchID   int          `json:"match_id" db:"match_id"`
	PlayerID  int          `json:"player_id" db:"player_id"`
	GoalTime  string       `json:"goal_time" db:"goal_time"`
	DeletedAt sql.NullTime `json:"-" db:"deleted_at"`
	CreatedAt time.Time    `json:"created_at" db:"created_at"`

	// Relations
	Match  *Match  `json:"match,omitempty" db:"-"`
	Player *Player `json:"player,omitempty" db:"-"`
}

// TableName returns the table name for Goal model
func (Goal) TableName() string {
	return "goals"
}
