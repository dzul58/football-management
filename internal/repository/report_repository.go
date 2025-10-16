package repository

import (
	"database/sql"
	"football-management-api/internal/models"
)

type ReportRepository interface {
	GetMatchReport(matchID int) (*models.MatchReport, error)
	GetTeamWins(teamID int, upToMatchID int) (int, error)
	GetTeamStatistics(teamID int) (*models.TeamStatistics, error)
	GetPlayerStatistics(playerID int) (*models.PlayerStatistics, error)
	GetTopScorers(limit int) ([]models.PlayerStatistics, error)
}

type reportRepository struct {
	db *sql.DB
}

func NewReportRepository(db *sql.DB) ReportRepository {
	return &reportRepository{db: db}
}

// GetMatchReport gets detailed match report
func (r *reportRepository) GetMatchReport(matchID int) (*models.MatchReport, error) {
	query := `
		SELECT m.id, m.match_date, m.match_time,
		       ht.id, ht.name, COALESCE(ht.logo_url, ''), ht.home_city,
		       at.id, at.name, COALESCE(at.logo_url, ''), at.home_city,
		       COALESCE(m.home_score, 0), COALESCE(m.away_score, 0)
		FROM matches m
		LEFT JOIN teams ht ON m.home_team_id = ht.id
		LEFT JOIN teams at ON m.away_team_id = at.id
		WHERE m.id = $1 AND m.deleted_at IS NULL AND m.status = 'Completed'
	`

	var report models.MatchReport
	var homeScore, awayScore int

	err := r.db.QueryRow(query, matchID).Scan(
		&report.MatchID,
		&report.MatchDate,
		&report.MatchTime,
		&report.HomeTeam.ID,
		&report.HomeTeam.Name,
		&report.HomeTeam.LogoURL,
		&report.HomeTeam.HomeCity,
		&report.AwayTeam.ID,
		&report.AwayTeam.Name,
		&report.AwayTeam.LogoURL,
		&report.AwayTeam.HomeCity,
		&homeScore,
		&awayScore,
	)

	if err != nil {
		return nil, err
	}

	report.FinalScore = models.ScoreInfo{
		Home: homeScore,
		Away: awayScore,
	}

	// Determine match result
	if homeScore > awayScore {
		report.MatchResult = "Tim Home Menang"
	} else if awayScore > homeScore {
		report.MatchResult = "Tim Away Menang"
	} else {
		report.MatchResult = "Draw"
	}

	// Get goal details
	goalsQuery := `
		SELECT g.id, g.player_id, p.name, t.name, g.goal_time
		FROM goals g
		JOIN players p ON g.player_id = p.id
		JOIN teams t ON p.team_id = t.id
		WHERE g.match_id = $1 AND g.deleted_at IS NULL
		ORDER BY g.goal_time ASC
	`

	rows, err := r.db.Query(goalsQuery, matchID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var goals []models.GoalDetail
	for rows.Next() {
		var goal models.GoalDetail
		err := rows.Scan(
			&goal.ID,
			&goal.PlayerID,
			&goal.PlayerName,
			&goal.TeamName,
			&goal.GoalTime,
		)
		if err != nil {
			return nil, err
		}
		goals = append(goals, goal)
	}
	report.GoalDetails = goals

	// Get top scorer
	topScorerQuery := `
		SELECT p.id, p.name, t.name, COUNT(*) as goals_count
		FROM goals g
		JOIN players p ON g.player_id = p.id
		JOIN teams t ON p.team_id = t.id
		WHERE g.match_id = $1 AND g.deleted_at IS NULL
		GROUP BY p.id, p.name, t.name
		ORDER BY goals_count DESC
		LIMIT 1
	`

	var topScorer models.TopScorerInfo
	err = r.db.QueryRow(topScorerQuery, matchID).Scan(
		&topScorer.PlayerID,
		&topScorer.PlayerName,
		&topScorer.TeamName,
		&topScorer.GoalsScored,
	)
	if err == nil {
		report.TopScorer = &topScorer
	}

	// Get team wins up to this match
	homeWins, _ := r.GetTeamWins(report.HomeTeam.ID, matchID)
	awayWins, _ := r.GetTeamWins(report.AwayTeam.ID, matchID)
	report.HomeTeamWins = homeWins
	report.AwayTeamWins = awayWins

	return &report, nil
}

// GetTeamWins gets total wins for a team up to a specific match
func (r *reportRepository) GetTeamWins(teamID int, upToMatchID int) (int, error) {
	query := `
		SELECT COUNT(*)
		FROM matches m
		WHERE m.id <= $1 
		AND m.status = 'Completed'
		AND m.deleted_at IS NULL
		AND (
			(m.home_team_id = $2 AND m.home_score > m.away_score)
			OR
			(m.away_team_id = $3 AND m.away_score > m.home_score)
		)
	`

	var wins int
	err := r.db.QueryRow(query, upToMatchID, teamID, teamID).Scan(&wins)
	if err != nil {
		return 0, err
	}

	return wins, nil
}

// GetTeamStatistics gets team statistics
func (r *reportRepository) GetTeamStatistics(teamID int) (*models.TeamStatistics, error) {
	var stats models.TeamStatistics
	stats.TeamID = teamID

	// Get team name
	teamQuery := "SELECT name FROM teams WHERE id = $1 AND deleted_at IS NULL"
	err := r.db.QueryRow(teamQuery, teamID).Scan(&stats.TeamName)
	if err != nil {
		return nil, err
	}

	// Get match statistics
	statsQuery := `
		SELECT 
			COUNT(*) as total_matches,
			SUM(CASE 
				WHEN (home_team_id = $1 AND home_score > away_score) OR (away_team_id = $2 AND away_score > home_score)
				THEN 1 ELSE 0 
			END) as wins,
			SUM(CASE 
				WHEN home_score = away_score 
				THEN 1 ELSE 0 
			END) as draws,
			SUM(CASE 
				WHEN (home_team_id = $3 AND home_score < away_score) OR (away_team_id = $4 AND away_score < home_score)
				THEN 1 ELSE 0 
			END) as losses,
			COALESCE(SUM(CASE 
				WHEN home_team_id = $5 THEN home_score 
				WHEN away_team_id = $6 THEN away_score 
			END), 0) as goals_scored,
			COALESCE(SUM(CASE 
				WHEN home_team_id = $7 THEN away_score 
				WHEN away_team_id = $8 THEN home_score 
			END), 0) as goals_conceded
		FROM matches
		WHERE (home_team_id = $9 OR away_team_id = $10)
		AND status = 'Completed'
		AND deleted_at IS NULL
	`

	err = r.db.QueryRow(statsQuery,
		teamID, teamID, teamID, teamID, teamID, teamID, teamID, teamID, teamID, teamID,
	).Scan(
		&stats.TotalMatches,
		&stats.TotalWins,
		&stats.TotalDraws,
		&stats.TotalLosses,
		&stats.GoalsScored,
		&stats.GoalsConceded,
	)

	if err != nil {
		return nil, err
	}

	return &stats, nil
}

// GetPlayerStatistics gets player goal statistics
func (r *reportRepository) GetPlayerStatistics(playerID int) (*models.PlayerStatistics, error) {
	query := `
		SELECT p.id, p.name, t.name, p.position, COUNT(g.id) as total_goals
		FROM players p
		LEFT JOIN teams t ON p.team_id = t.id
		LEFT JOIN goals g ON p.id = g.player_id AND g.deleted_at IS NULL
		WHERE p.id = $1 AND p.deleted_at IS NULL
		GROUP BY p.id, p.name, t.name, p.position
	`

	var stats models.PlayerStatistics
	err := r.db.QueryRow(query, playerID).Scan(
		&stats.PlayerID,
		&stats.PlayerName,
		&stats.TeamName,
		&stats.Position,
		&stats.TotalGoals,
	)

	if err != nil {
		return nil, err
	}

	return &stats, nil
}

// GetTopScorers gets top scorers
func (r *reportRepository) GetTopScorers(limit int) ([]models.PlayerStatistics, error) {
	query := `
		SELECT p.id, p.name, t.name, p.position, COUNT(g.id) as total_goals
		FROM players p
		LEFT JOIN teams t ON p.team_id = t.id
		LEFT JOIN goals g ON p.id = g.player_id AND g.deleted_at IS NULL
		WHERE p.deleted_at IS NULL
		GROUP BY p.id, p.name, t.name, p.position
		HAVING total_goals > 0
		ORDER BY total_goals DESC
		LIMIT $1
	`

	rows, err := r.db.Query(query, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var scorers []models.PlayerStatistics
	for rows.Next() {
		var stats models.PlayerStatistics
		err := rows.Scan(
			&stats.PlayerID,
			&stats.PlayerName,
			&stats.TeamName,
			&stats.Position,
			&stats.TotalGoals,
		)
		if err != nil {
			return nil, err
		}
		scorers = append(scorers, stats)
	}

	return scorers, nil
}
