package models

// MatchReport represents a detailed match report
type MatchReport struct {
	MatchID      int            `json:"match_id"`
	MatchDate    string         `json:"match_date"`
	MatchTime    string         `json:"match_time"`
	HomeTeam     TeamInfo       `json:"home_team"`
	AwayTeam     TeamInfo       `json:"away_team"`
	FinalScore   ScoreInfo      `json:"final_score"`
	MatchResult  string         `json:"match_result"`
	TopScorer    *TopScorerInfo `json:"top_scorer"`
	HomeTeamWins int            `json:"home_team_total_wins"`
	AwayTeamWins int            `json:"away_team_total_wins"`
	GoalDetails  []GoalDetail   `json:"goal_details"`
}

// TeamInfo represents basic team information in a report
type TeamInfo struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	LogoURL  string `json:"logo_url,omitempty"`
	HomeCity string `json:"home_city"`
}

// ScoreInfo represents the final score of a match
type ScoreInfo struct {
	Home int `json:"home"`
	Away int `json:"away"`
}

// TopScorerInfo represents the top scorer in a match
type TopScorerInfo struct {
	PlayerID    int    `json:"player_id"`
	PlayerName  string `json:"player_name"`
	TeamName    string `json:"team_name"`
	GoalsScored int    `json:"goals_scored"`
}

// GoalDetail represents detailed information about a goal
type GoalDetail struct {
	ID         int    `json:"id"`
	PlayerID   int    `json:"player_id"`
	PlayerName string `json:"player_name"`
	TeamName   string `json:"team_name"`
	GoalTime   string `json:"goal_time"`
}

// TeamStatistics represents team statistics
type TeamStatistics struct {
	TeamID        int    `json:"team_id"`
	TeamName      string `json:"team_name"`
	TotalMatches  int    `json:"total_matches"`
	TotalWins     int    `json:"total_wins"`
	TotalDraws    int    `json:"total_draws"`
	TotalLosses   int    `json:"total_losses"`
	GoalsScored   int    `json:"goals_scored"`
	GoalsConceded int    `json:"goals_conceded"`
}

// PlayerStatistics represents player goal statistics
type PlayerStatistics struct {
	PlayerID   int    `json:"player_id"`
	PlayerName string `json:"player_name"`
	TeamName   string `json:"team_name"`
	Position   string `json:"position"`
	TotalGoals int    `json:"total_goals"`
}
