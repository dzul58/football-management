package repository

import (
	"database/sql"
	"errors"
	"football-management-api/internal/models"
	"time"
)

type TeamRepository interface {
	Create(team *models.Team) error
	FindByID(id int) (*models.Team, error)
	FindAll(limit, offset int) ([]models.Team, int64, error)
	Update(id int, team *models.Team) error
	Delete(id int) error
	FindByName(name string) (*models.Team, error)
}

type teamRepository struct {
	db *sql.DB
}

func NewTeamRepository(db *sql.DB) TeamRepository {
	return &teamRepository{db: db}
}

// Create creates a new team
func (r *teamRepository) Create(team *models.Team) error {
	query := `
		INSERT INTO teams (name, logo_url, founded_year, home_address, home_city, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id
	`

	err := r.db.QueryRow(query,
		team.Name,
		team.LogoURL,
		team.FoundedYear,
		team.HomeAddress,
		team.HomeCity,
		time.Now(),
		time.Now(),
	).Scan(&team.ID)

	if err != nil {
		return err
	}

	return nil
}

// FindByID finds a team by ID
func (r *teamRepository) FindByID(id int) (*models.Team, error) {
	query := `
		SELECT id, name, logo_url, founded_year, home_address, home_city, created_at, updated_at
		FROM teams
		WHERE id = $1 AND deleted_at IS NULL
	`

	var team models.Team
	err := r.db.QueryRow(query, id).Scan(
		&team.ID,
		&team.Name,
		&team.LogoURL,
		&team.FoundedYear,
		&team.HomeAddress,
		&team.HomeCity,
		&team.CreatedAt,
		&team.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, errors.New("tim tidak ditemukan")
	}

	if err != nil {
		return nil, err
	}

	return &team, nil
}

// FindAll finds all teams with pagination
func (r *teamRepository) FindAll(limit, offset int) ([]models.Team, int64, error) {
	// Get total count
	var total int64
	countQuery := "SELECT COUNT(*) FROM teams WHERE deleted_at IS NULL"
	err := r.db.QueryRow(countQuery).Scan(&total)
	if err != nil {
		return nil, 0, err
	}

	// Get teams
	query := `
		SELECT id, name, logo_url, founded_year, home_address, home_city, created_at, updated_at
		FROM teams
		WHERE deleted_at IS NULL
		ORDER BY created_at DESC
		LIMIT $1 OFFSET $2
	`

	rows, err := r.db.Query(query, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var teams []models.Team
	for rows.Next() {
		var team models.Team
		err := rows.Scan(
			&team.ID,
			&team.Name,
			&team.LogoURL,
			&team.FoundedYear,
			&team.HomeAddress,
			&team.HomeCity,
			&team.CreatedAt,
			&team.UpdatedAt,
		)
		if err != nil {
			return nil, 0, err
		}
		teams = append(teams, team)
	}

	return teams, total, nil
}

// Update updates a team
func (r *teamRepository) Update(id int, team *models.Team) error {
	query := `
		UPDATE teams
		SET name = $1, logo_url = $2, founded_year = $3, home_address = $4, home_city = $5, updated_at = $6
		WHERE id = $7 AND deleted_at IS NULL
	`

	result, err := r.db.Exec(query,
		team.Name,
		team.LogoURL,
		team.FoundedYear,
		team.HomeAddress,
		team.HomeCity,
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
		return errors.New("tim tidak ditemukan")
	}

	return nil
}

// Delete soft deletes a team
func (r *teamRepository) Delete(id int) error {
	query := `
		UPDATE teams
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
		return errors.New("tim tidak ditemukan")
	}

	return nil
}

// FindByName finds a team by name
func (r *teamRepository) FindByName(name string) (*models.Team, error) {
	query := `
		SELECT id, name, logo_url, founded_year, home_address, home_city, created_at, updated_at
		FROM teams
		WHERE name = $1 AND deleted_at IS NULL
	`

	var team models.Team
	err := r.db.QueryRow(query, name).Scan(
		&team.ID,
		&team.Name,
		&team.LogoURL,
		&team.FoundedYear,
		&team.HomeAddress,
		&team.HomeCity,
		&team.CreatedAt,
		&team.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &team, nil
}
