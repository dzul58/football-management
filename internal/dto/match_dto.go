package dto

// CreateMatchRequest represents request to create a match
type CreateMatchRequest struct {
	MatchDate  string `json:"match_date" binding:"required"`
	MatchTime  string `json:"match_time" binding:"required"`
	HomeTeamID int    `json:"home_team_id" binding:"required"`
	AwayTeamID int    `json:"away_team_id" binding:"required"`
}

// UpdateMatchRequest represents request to update a match
type UpdateMatchRequest struct {
	MatchDate  string `json:"match_date"`
	MatchTime  string `json:"match_time"`
	HomeTeamID int    `json:"home_team_id"`
	AwayTeamID int    `json:"away_team_id"`
	Status     string `json:"status"`
}

// UpdateMatchResultRequest represents request to update match result
type UpdateMatchResultRequest struct {
	HomeScore int               `json:"home_score" binding:"required,min=0"`
	AwayScore int               `json:"away_score" binding:"required,min=0"`
	Goals     []GoalInputDetail `json:"goals" binding:"required"`
}

// GoalInputDetail represents goal details in match result
type GoalInputDetail struct {
	PlayerID int    `json:"player_id" binding:"required"`
	GoalTime string `json:"goal_time" binding:"required"`
}

// MatchResponse represents match data in response
type MatchResponse struct {
	ID           int    `json:"id"`
	MatchDate    string `json:"match_date"`
	MatchTime    string `json:"match_time"`
	HomeTeamID   int    `json:"home_team_id"`
	HomeTeamName string `json:"home_team_name,omitempty"`
	AwayTeamID   int    `json:"away_team_id"`
	AwayTeamName string `json:"away_team_name,omitempty"`
	HomeScore    *int   `json:"home_score"`
	AwayScore    *int   `json:"away_score"`
	Status       string `json:"status"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}
