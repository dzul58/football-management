package repository

import (
	"database/sql"
	"errors"
	"football-management-api/internal/models"
	"time"
)

type GoalRepository interface {
	Create(goal *models.Goal) error
	FindByID(id int) (*models.Goal, error)
	FindByMatchID(matchID int) ([]models.Goal, error)
	Delete(id int) error
	DeleteByMatchID(matchID int) error
	FindTopScorerInMatch(matchID int) (*models.TopScorerInfo, error)
}

type goalRepository struct {
	db *sql.DB
}

func NewGoalRepository(db *sql.DB) GoalRepository {
	return &goalRepository{db: db}
}

// Create creates a new goal
func (r *goalRepository) Create(goal *models.Goal) error {
	query := `
		INSERT INTO goals (match_id, player_id, goal_time, created_at)
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`

	err := r.db.QueryRow(query,
		goal.MatchID,
		goal.PlayerID,
		goal.GoalTime,
		time.Now(),
	).Scan(&goal.ID)

	if err != nil {
		return err
	}

	return nil
}

// FindByID finds a goal by ID
func (r *goalRepository) FindByID(id int) (*models.Goal, error) {
	query := `
		SELECT g.id, g.match_id, g.player_id, g.goal_time, g.created_at,
		       p.name, t.name
		FROM goals g
		LEFT JOIN players p ON g.player_id = p.id
		LEFT JOIN teams t ON p.team_id = t.id
		WHERE g.id = $1 AND g.deleted_at IS NULL
	`

	var goal models.Goal
	var player models.Player
	var team models.Team

	err := r.db.QueryRow(query, id).Scan(
		&goal.ID,
		&goal.MatchID,
		&goal.PlayerID,
		&goal.GoalTime,
		&goal.CreatedAt,
		&player.Name,
		&team.Name,
	)

	if err == sql.ErrNoRows {
		return nil, errors.New("gol tidak ditemukan")
	}

	if err != nil {
		return nil, err
	}

	player.Team = &team
	goal.Player = &player

	return &goal, nil
}

// FindByMatchID finds all goals in a match
func (r *goalRepository) FindByMatchID(matchID int) ([]models.Goal, error) {
	query := `
		SELECT g.id, g.match_id, g.player_id, g.goal_time, g.created_at,
		       p.name, t.name
		FROM goals g
		LEFT JOIN players p ON g.player_id = p.id
		LEFT JOIN teams t ON p.team_id = t.id
		WHERE g.match_id = $1 AND g.deleted_at IS NULL
		ORDER BY g.goal_time ASC
	`

	rows, err := r.db.Query(query, matchID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var goals []models.Goal
	for rows.Next() {
		var goal models.Goal
		var playerName, teamName string

		err := rows.Scan(
			&goal.ID,
			&goal.MatchID,
			&goal.PlayerID,
			&goal.GoalTime,
			&goal.CreatedAt,
			&playerName,
			&teamName,
		)
		if err != nil {
			return nil, err
		}

		goal.Player = &models.Player{Name: playerName, Team: &models.Team{Name: teamName}}
		goals = append(goals, goal)
	}

	return goals, nil
}

// Delete soft deletes a goal
func (r *goalRepository) Delete(id int) error {
	query := `
		UPDATE goals
		SET deleted_at = $1
		WHERE id = $2 AND deleted_at IS NULL
	`

	result, err := r.db.Exec(query, time.Now(), id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("gol tidak ditemukan")
	}

	return nil
}

// DeleteByMatchID deletes all goals for a match
func (r *goalRepository) DeleteByMatchID(matchID int) error {
	query := `
		UPDATE goals
		SET deleted_at = $1
		WHERE match_id = $2 AND deleted_at IS NULL
	`

	_, err := r.db.Exec(query, time.Now(), matchID)
	return err
}

// FindTopScorerInMatch finds the top scorer in a match
func (r *goalRepository) FindTopScorerInMatch(matchID int) (*models.TopScorerInfo, error) {
	query := `
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
	err := r.db.QueryRow(query, matchID).Scan(
		&topScorer.PlayerID,
		&topScorer.PlayerName,
		&topScorer.TeamName,
		&topScorer.GoalsScored,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &topScorer, nil
}
