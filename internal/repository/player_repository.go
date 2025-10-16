package repository

import (
	"database/sql"
	"errors"
	"football-management-api/internal/models"
	"time"
)

type PlayerRepository interface {
	Create(player *models.Player) error
	FindByID(id int) (*models.Player, error)
	FindAll(limit, offset int) ([]models.Player, int64, error)
	FindByTeamID(teamID int) ([]models.Player, error)
	Update(id int, player *models.Player) error
	Delete(id int) error
	CheckJerseyNumberExists(teamID, jerseyNumber, excludePlayerID int) (bool, error)
}

type playerRepository struct {
	db *sql.DB
}

func NewPlayerRepository(db *sql.DB) PlayerRepository {
	return &playerRepository{db: db}
}

// Create creates a new player
func (r *playerRepository) Create(player *models.Player) error {
	query := `
		INSERT INTO players (team_id, name, height, weight, position, jersey_number, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id
	`

	err := r.db.QueryRow(query,
		player.TeamID,
		player.Name,
		player.Height,
		player.Weight,
		player.Position,
		player.JerseyNumber,
		time.Now(),
		time.Now(),
	).Scan(&player.ID)

	if err != nil {
		return err
	}

	return nil
}

// FindByID finds a player by ID
func (r *playerRepository) FindByID(id int) (*models.Player, error) {
	query := `
		SELECT p.id, p.team_id, p.name, p.height, p.weight, p.position, p.jersey_number, 
		       p.created_at, p.updated_at,
		       t.id, t.name, t.logo_url, t.home_city
		FROM players p
		LEFT JOIN teams t ON p.team_id = t.id AND t.deleted_at IS NULL
		WHERE p.id = $1 AND p.deleted_at IS NULL
	`

	var player models.Player
	var team models.Team
	var logoURL sql.NullString

	err := r.db.QueryRow(query, id).Scan(
		&player.ID,
		&player.TeamID,
		&player.Name,
		&player.Height,
		&player.Weight,
		&player.Position,
		&player.JerseyNumber,
		&player.CreatedAt,
		&player.UpdatedAt,
		&team.ID,
		&team.Name,
		&logoURL,
		&team.HomeCity,
	)

	if err == sql.ErrNoRows {
		return nil, errors.New("pemain tidak ditemukan")
	}

	if err != nil {
		return nil, err
	}

	team.LogoURL = logoURL
	player.Team = &team

	return &player, nil
}

// FindAll finds all players with pagination
func (r *playerRepository) FindAll(limit, offset int) ([]models.Player, int64, error) {
	// Get total count
	var total int64
	countQuery := "SELECT COUNT(*) FROM players WHERE deleted_at IS NULL"
	err := r.db.QueryRow(countQuery).Scan(&total)
	if err != nil {
		return nil, 0, err
	}

	// Get players
	query := `
		SELECT p.id, p.team_id, p.name, p.height, p.weight, p.position, p.jersey_number,
		       p.created_at, p.updated_at,
		       t.id, t.name, t.logo_url, t.home_city
		FROM players p
		LEFT JOIN teams t ON p.team_id = t.id AND t.deleted_at IS NULL
		WHERE p.deleted_at IS NULL
		ORDER BY p.created_at DESC
		LIMIT $1 OFFSET $2
	`

	rows, err := r.db.Query(query, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var players []models.Player
	for rows.Next() {
		var player models.Player
		var team models.Team
		var logoURL sql.NullString

		err := rows.Scan(
			&player.ID,
			&player.TeamID,
			&player.Name,
			&player.Height,
			&player.Weight,
			&player.Position,
			&player.JerseyNumber,
			&player.CreatedAt,
			&player.UpdatedAt,
			&team.ID,
			&team.Name,
			&logoURL,
			&team.HomeCity,
		)
		if err != nil {
			return nil, 0, err
		}

		team.LogoURL = logoURL
		player.Team = &team
		players = append(players, player)
	}

	return players, total, nil
}

// FindByTeamID finds all players by team ID
func (r *playerRepository) FindByTeamID(teamID int) ([]models.Player, error) {
	query := `
		SELECT id, team_id, name, height, weight, position, jersey_number, created_at, updated_at
		FROM players
		WHERE team_id = $1 AND deleted_at IS NULL
		ORDER BY jersey_number ASC
	`

	rows, err := r.db.Query(query, teamID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var players []models.Player
	for rows.Next() {
		var player models.Player
		err := rows.Scan(
			&player.ID,
			&player.TeamID,
			&player.Name,
			&player.Height,
			&player.Weight,
			&player.Position,
			&player.JerseyNumber,
			&player.CreatedAt,
			&player.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		players = append(players, player)
	}

	return players, nil
}

// Update updates a player
func (r *playerRepository) Update(id int, player *models.Player) error {
	query := `
		UPDATE players
		SET team_id = $1, name = $2, height = $3, weight = $4, position = $5, jersey_number = $6, updated_at = $7
		WHERE id = $8 AND deleted_at IS NULL
	`

	result, err := r.db.Exec(query,
		player.TeamID,
		player.Name,
		player.Height,
		player.Weight,
		player.Position,
		player.JerseyNumber,
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
		return errors.New("pemain tidak ditemukan")
	}

	return nil
}

// Delete soft deletes a player
func (r *playerRepository) Delete(id int) error {
	query := `
		UPDATE players
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
		return errors.New("pemain tidak ditemukan")
	}

	return nil
}

// CheckJerseyNumberExists checks if a jersey number exists for a team
func (r *playerRepository) CheckJerseyNumberExists(teamID, jerseyNumber, excludePlayerID int) (bool, error) {
	query := `
		SELECT COUNT(*) 
		FROM players 
		WHERE team_id = $1 AND jersey_number = $2 AND id != $3 AND deleted_at IS NULL
	`

	var count int
	err := r.db.QueryRow(query, teamID, jerseyNumber, excludePlayerID).Scan(&count)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}
