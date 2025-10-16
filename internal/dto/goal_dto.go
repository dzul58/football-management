package dto

// CreateGoalRequest represents request to create a goal
type CreateGoalRequest struct {
	MatchID  int    `json:"match_id" binding:"required"`
	PlayerID int    `json:"player_id" binding:"required"`
	GoalTime string `json:"goal_time" binding:"required"`
}

// GoalResponse represents goal data in response
type GoalResponse struct {
	ID         int    `json:"id"`
	MatchID    int    `json:"match_id"`
	PlayerID   int    `json:"player_id"`
	PlayerName string `json:"player_name,omitempty"`
	TeamName   string `json:"team_name,omitempty"`
	GoalTime   string `json:"goal_time"`
	CreatedAt  string `json:"created_at"`
}
