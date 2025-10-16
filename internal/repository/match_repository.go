package repository

import (
	"database/sql"
	"errors"
	"football-management-api/internal/models"
	"time"
)

type MatchRepository interface {
	Create(match *models.Match) error
	FindByID(id int) (*models.Match, error)
	FindAll(limit, offset int) ([]models.Match, int64, error)
	FindByTeamID(teamID int) ([]models.Match, error)
	Update(id int, match *models.Match) error
	UpdateResult(id int, homeScore, awayScore int, status models.MatchStatus) error
	Delete(id int) error
	FindCompletedMatches() ([]models.Match, error)
}

type matchRepository struct {
	db *sql.DB
}

func NewMatchRepository(db *sql.DB) MatchRepository {
	return &matchRepository{db: db}
}

// Create creates a new match
func (r *matchRepository) Create(match *models.Match) error {
	query := `
		INSERT INTO matches (match_date, match_time, home_team_id, away_team_id, status, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id
	`

	err := r.db.QueryRow(query,
		match.MatchDate,
		match.MatchTime,
		match.HomeTeamID,
		match.AwayTeamID,
		match.Status,
		time.Now(),
		time.Now(),
	).Scan(&match.ID)

	if err != nil {
		return err
	}

	return nil
}

// FindByID finds a match by ID
func (r *matchRepository) FindByID(id int) (*models.Match, error) {
	query := `
		SELECT m.id, m.match_date, m.match_time, m.home_team_id, m.away_team_id, 
		       m.home_score, m.away_score, m.status, m.created_at, m.updated_at,
		       ht.id, ht.name, ht.logo_url, ht.home_city,
		       at.id, at.name, at.logo_url, at.home_city
		FROM matches m
		LEFT JOIN teams ht ON m.home_team_id = ht.id AND ht.deleted_at IS NULL
		LEFT JOIN teams at ON m.away_team_id = at.id AND at.deleted_at IS NULL
		WHERE m.id = $1 AND m.deleted_at IS NULL
	`

	var match models.Match
	var homeTeam, awayTeam models.Team
	var homeLogoURL, awayLogoURL sql.NullString

	err := r.db.QueryRow(query, id).Scan(
		&match.ID,
		&match.MatchDate,
		&match.MatchTime,
		&match.HomeTeamID,
		&match.AwayTeamID,
		&match.HomeScore,
		&match.AwayScore,
		&match.Status,
		&match.CreatedAt,
		&match.UpdatedAt,
		&homeTeam.ID,
		&homeTeam.Name,
		&homeLogoURL,
		&homeTeam.HomeCity,
		&awayTeam.ID,
		&awayTeam.Name,
		&awayLogoURL,
		&awayTeam.HomeCity,
	)

	if err == sql.ErrNoRows {
		return nil, errors.New("pertandingan tidak ditemukan")
	}

	if err != nil {
		return nil, err
	}

	homeTeam.LogoURL = homeLogoURL
	awayTeam.LogoURL = awayLogoURL
	match.HomeTeam = &homeTeam
	match.AwayTeam = &awayTeam

	return &match, nil
}

// FindAll finds all matches with pagination
func (r *matchRepository) FindAll(limit, offset int) ([]models.Match, int64, error) {
	// Get total count
	var total int64
	countQuery := "SELECT COUNT(*) FROM matches WHERE deleted_at IS NULL"
	err := r.db.QueryRow(countQuery).Scan(&total)
	if err != nil {
		return nil, 0, err
	}

	// Get matches
	query := `
		SELECT m.id, m.match_date, m.match_time, m.home_team_id, m.away_team_id,
		       m.home_score, m.away_score, m.status, m.created_at, m.updated_at,
		       ht.id, ht.name, ht.logo_url, ht.home_city,
		       at.id, at.name, at.logo_url, at.home_city
		FROM matches m
		LEFT JOIN teams ht ON m.home_team_id = ht.id AND ht.deleted_at IS NULL
		LEFT JOIN teams at ON m.away_team_id = at.id AND at.deleted_at IS NULL
		WHERE m.deleted_at IS NULL
		ORDER BY m.match_date DESC, m.match_time DESC
		LIMIT $1 OFFSET $2
	`

	rows, err := r.db.Query(query, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var matches []models.Match
	for rows.Next() {
		var match models.Match
		var homeTeam, awayTeam models.Team
		var homeLogoURL, awayLogoURL sql.NullString

		err := rows.Scan(
			&match.ID,
			&match.MatchDate,
			&match.MatchTime,
			&match.HomeTeamID,
			&match.AwayTeamID,
			&match.HomeScore,
			&match.AwayScore,
			&match.Status,
			&match.CreatedAt,
			&match.UpdatedAt,
			&homeTeam.ID,
			&homeTeam.Name,
			&homeLogoURL,
			&homeTeam.HomeCity,
			&awayTeam.ID,
			&awayTeam.Name,
			&awayLogoURL,
			&awayTeam.HomeCity,
		)
		if err != nil {
			return nil, 0, err
		}

		homeTeam.LogoURL = homeLogoURL
		awayTeam.LogoURL = awayLogoURL
		match.HomeTeam = &homeTeam
		match.AwayTeam = &awayTeam
		matches = append(matches, match)
	}

	return matches, total, nil
}

// FindByTeamID finds all matches for a team
func (r *matchRepository) FindByTeamID(teamID int) ([]models.Match, error) {
	query := `
		SELECT m.id, m.match_date, m.match_time, m.home_team_id, m.away_team_id,
		       m.home_score, m.away_score, m.status, m.created_at, m.updated_at
		FROM matches m
		WHERE (m.home_team_id = $1 OR m.away_team_id = $2) AND m.deleted_at IS NULL
		ORDER BY m.match_date DESC, m.match_time DESC
	`

	rows, err := r.db.Query(query, teamID, teamID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var matches []models.Match
	for rows.Next() {
		var match models.Match
		err := rows.Scan(
			&match.ID,
			&match.MatchDate,
			&match.MatchTime,
			&match.HomeTeamID,
			&match.AwayTeamID,
			&match.HomeScore,
			&match.AwayScore,
			&match.Status,
			&match.CreatedAt,
			&match.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		matches = append(matches, match)
	}

	return matches, nil
}

// Update updates a match
func (r *matchRepository) Update(id int, match *models.Match) error {
	query := `
		UPDATE matches
		SET match_date = $1, match_time = $2, home_team_id = $3, away_team_id = $4, status = $5, updated_at = $6
		WHERE id = $7 AND deleted_at IS NULL
	`

	result, err := r.db.Exec(query,
		match.MatchDate,
		match.MatchTime,
		match.HomeTeamID,
		match.AwayTeamID,
		match.Status,
		time.Now(),
		id,
	)

	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("pertandingan tidak ditemukan")
	}

	return nil
}

// UpdateResult updates match result
func (r *matchRepository) UpdateResult(id int, homeScore, awayScore int, status models.MatchStatus) error {
	query := `
		UPDATE matches
		SET home_score = $1, away_score = $2, status = $3, updated_at = $4
		WHERE id = $5 AND deleted_at IS NULL
	`

	result, err := r.db.Exec(query, homeScore, awayScore, status, time.Now(), id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("pertandingan tidak ditemukan")
	}

	return nil
}

// Delete soft deletes a match
func (r *matchRepository) Delete(id int) error {
	query := `
		UPDATE matches
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
		return errors.New("pertandingan tidak ditemukan")
	}

	return nil
}

// FindCompletedMatches finds all completed matches
func (r *matchRepository) FindCompletedMatches() ([]models.Match, error) {
	query := `
		SELECT m.id, m.match_date, m.match_time, m.home_team_id, m.away_team_id,
		       m.home_score, m.away_score, m.status, m.created_at, m.updated_at,
		       ht.id, ht.name, ht.logo_url, ht.home_city,
		       at.id, at.name, at.logo_url, at.home_city
		FROM matches m
		LEFT JOIN teams ht ON m.home_team_id = ht.id AND ht.deleted_at IS NULL
		LEFT JOIN teams at ON m.away_team_id = at.id AND at.deleted_at IS NULL
		WHERE m.status = 'Completed' AND m.deleted_at IS NULL
		ORDER BY m.match_date DESC, m.match_time DESC
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var matches []models.Match
	for rows.Next() {
		var match models.Match
		var homeTeam, awayTeam models.Team
		var homeLogoURL, awayLogoURL sql.NullString

		err := rows.Scan(
			&match.ID,
			&match.MatchDate,
			&match.MatchTime,
			&match.HomeTeamID,
			&match.AwayTeamID,
			&match.HomeScore,
			&match.AwayScore,
			&match.Status,
			&match.CreatedAt,
			&match.UpdatedAt,
			&homeTeam.ID,
			&homeTeam.Name,
			&homeLogoURL,
			&homeTeam.HomeCity,
			&awayTeam.ID,
			&awayTeam.Name,
			&awayLogoURL,
			&awayTeam.HomeCity,
		)
		if err != nil {
			return nil, err
		}

		homeTeam.LogoURL = homeLogoURL
		awayTeam.LogoURL = awayLogoURL
		match.HomeTeam = &homeTeam
		match.AwayTeam = &awayTeam
		matches = append(matches, match)
	}

	return matches, nil
}
