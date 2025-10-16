package dto

// CreatePlayerRequest represents request to create a player
type CreatePlayerRequest struct {
	TeamID       int     `json:"team_id" binding:"required"`
	Name         string  `json:"name" binding:"required"`
	Height       float64 `json:"height" binding:"required,min=100,max=250"`
	Weight       float64 `json:"weight" binding:"required,min=30,max=200"`
	Position     string  `json:"position" binding:"required"`
	JerseyNumber int     `json:"jersey_number" binding:"required,min=1,max=99"`
}

// UpdatePlayerRequest represents request to update a player
type UpdatePlayerRequest struct {
	TeamID       int     `json:"team_id"`
	Name         string  `json:"name"`
	Height       float64 `json:"height" binding:"omitempty,min=100,max=250"`
	Weight       float64 `json:"weight" binding:"omitempty,min=30,max=200"`
	Position     string  `json:"position"`
	JerseyNumber int     `json:"jersey_number" binding:"omitempty,min=1,max=99"`
}

// PlayerResponse represents player data in response
type PlayerResponse struct {
	ID           int     `json:"id"`
	TeamID       int     `json:"team_id"`
	TeamName     string  `json:"team_name,omitempty"`
	Name         string  `json:"name"`
	Height       float64 `json:"height"`
	Weight       float64 `json:"weight"`
	Position     string  `json:"position"`
	JerseyNumber int     `json:"jersey_number"`
	CreatedAt    string  `json:"created_at"`
	UpdatedAt    string  `json:"updated_at"`
}
