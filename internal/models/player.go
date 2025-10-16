package models

import (
	"database/sql"
	"time"
)

// PlayerPosition represents available player positions
type PlayerPosition string

const (
	PositionPenyerang     PlayerPosition = "Penyerang"
	PositionGelandang     PlayerPosition = "Gelandang"
	PositionBertahan      PlayerPosition = "Bertahan"
	PositionPenjagaGawang PlayerPosition = "Penjaga Gawang"
)

// Player represents a football player entity
type Player struct {
	ID           int            `json:"id" db:"id"`
	TeamID       int            `json:"team_id" db:"team_id"`
	Name         string         `json:"name" db:"name"`
	Height       float64        `json:"height" db:"height"`
	Weight       float64        `json:"weight" db:"weight"`
	Position     PlayerPosition `json:"position" db:"position"`
	JerseyNumber int            `json:"jersey_number" db:"jersey_number"`
	DeletedAt    sql.NullTime   `json:"-" db:"deleted_at"`
	CreatedAt    time.Time      `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at" db:"updated_at"`

	// Relations
	Team *Team `json:"team,omitempty" db:"-"`
}

// TableName returns the table name for Player model
func (Player) TableName() string {
	return "players"
}

// IsValidPosition checks if a position is valid
func IsValidPosition(position string) bool {
	validPositions := []PlayerPosition{
		PositionPenyerang,
		PositionGelandang,
		PositionBertahan,
		PositionPenjagaGawang,
	}

	for _, valid := range validPositions {
		if PlayerPosition(position) == valid {
			return true
		}
	}
	return false
}
