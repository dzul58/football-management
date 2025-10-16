package service

import (
	"errors"
	"football-management-api/internal/dto"
	"football-management-api/internal/models"
	"football-management-api/internal/repository"
	"football-management-api/internal/utils"
)

type PlayerService interface {
	Create(req dto.CreatePlayerRequest) (*dto.PlayerResponse, error)
	GetByID(id int) (*dto.PlayerResponse, error)
	GetAll(page, limit int) ([]dto.PlayerResponse, dto.PaginationMeta, error)
	GetByTeamID(teamID int) ([]dto.PlayerResponse, error)
	Update(id int, req dto.UpdatePlayerRequest) (*dto.PlayerResponse, error)
	Delete(id int) error
}

type playerService struct {
	playerRepo repository.PlayerRepository
	teamRepo   repository.TeamRepository
}

func NewPlayerService(playerRepo repository.PlayerRepository, teamRepo repository.TeamRepository) PlayerService {
	return &playerService{
		playerRepo: playerRepo,
		teamRepo:   teamRepo,
	}
}

// Create creates a new player
func (s *playerService) Create(req dto.CreatePlayerRequest) (*dto.PlayerResponse, error) {
	// Validate team exists
	_, err := s.teamRepo.FindByID(req.TeamID)
	if err != nil {
		return nil, errors.New("tim tidak ditemukan")
	}

	// Check if jersey number already exists in the team
	exists, err := s.playerRepo.CheckJerseyNumberExists(req.TeamID, req.JerseyNumber, 0)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, errors.New("nomor punggung sudah digunakan oleh pemain lain di tim ini")
	}

	player := &models.Player{
		TeamID:       req.TeamID,
		Name:         req.Name,
		Height:       req.Height,
		Weight:       req.Weight,
		Position:     models.PlayerPosition(req.Position),
		JerseyNumber: req.JerseyNumber,
	}

	err = s.playerRepo.Create(player)
	if err != nil {
		return nil, err
	}

	// Get player with team info
	createdPlayer, err := s.playerRepo.FindByID(player.ID)
	if err != nil {
		return nil, err
	}

	return s.mapToResponse(createdPlayer), nil
}

// GetByID gets a player by ID
func (s *playerService) GetByID(id int) (*dto.PlayerResponse, error) {
	player, err := s.playerRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	return s.mapToResponse(player), nil
}

// GetAll gets all players with pagination
func (s *playerService) GetAll(page, limit int) ([]dto.PlayerResponse, dto.PaginationMeta, error) {
	offset := utils.CalculateOffset(page, limit)

	players, total, err := s.playerRepo.FindAll(limit, offset)
	if err != nil {
		return nil, dto.PaginationMeta{}, err
	}

	var responses []dto.PlayerResponse
	for _, player := range players {
		responses = append(responses, *s.mapToResponse(&player))
	}

	meta := dto.PaginationMeta{
		CurrentPage: page,
		PerPage:     limit,
		Total:       total,
		TotalPages:  utils.CalculateTotalPages(total, limit),
	}

	return responses, meta, nil
}

// GetByTeamID gets all players in a team
func (s *playerService) GetByTeamID(teamID int) ([]dto.PlayerResponse, error) {
	// Validate team exists
	_, err := s.teamRepo.FindByID(teamID)
	if err != nil {
		return nil, errors.New("tim tidak ditemukan")
	}

	players, err := s.playerRepo.FindByTeamID(teamID)
	if err != nil {
		return nil, err
	}

	var responses []dto.PlayerResponse
	for _, player := range players {
		responses = append(responses, *s.mapToResponseSimple(&player))
	}

	return responses, nil
}

// Update updates a player
func (s *playerService) Update(id int, req dto.UpdatePlayerRequest) (*dto.PlayerResponse, error) {
	// Get existing player
	existingPlayer, err := s.playerRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	// Update fields if provided
	if req.TeamID != 0 {
		// Validate new team exists
		_, err := s.teamRepo.FindByID(req.TeamID)
		if err != nil {
			return nil, errors.New("tim tidak ditemukan")
		}
		existingPlayer.TeamID = req.TeamID
	}

	if req.Name != "" {
		existingPlayer.Name = req.Name
	}

	if req.Height != 0 {
		existingPlayer.Height = req.Height
	}

	if req.Weight != 0 {
		existingPlayer.Weight = req.Weight
	}

	if req.Position != "" {
		existingPlayer.Position = models.PlayerPosition(req.Position)
	}

	if req.JerseyNumber != 0 {
		// Check if jersey number is being changed
		if req.JerseyNumber != existingPlayer.JerseyNumber {
			// Check if new jersey number already exists in the team
			exists, err := s.playerRepo.CheckJerseyNumberExists(existingPlayer.TeamID, req.JerseyNumber, id)
			if err != nil {
				return nil, err
			}
			if exists {
				return nil, errors.New("nomor punggung sudah digunakan oleh pemain lain di tim ini")
			}
		}
		existingPlayer.JerseyNumber = req.JerseyNumber
	}

	err = s.playerRepo.Update(id, existingPlayer)
	if err != nil {
		return nil, err
	}

	// Get updated player
	updatedPlayer, err := s.playerRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	return s.mapToResponse(updatedPlayer), nil
}

// Delete deletes a player
func (s *playerService) Delete(id int) error {
	return s.playerRepo.Delete(id)
}

// mapToResponse maps player model to response DTO with team info
func (s *playerService) mapToResponse(player *models.Player) *dto.PlayerResponse {
	response := &dto.PlayerResponse{
		ID:           player.ID,
		TeamID:       player.TeamID,
		Name:         player.Name,
		Height:       player.Height,
		Weight:       player.Weight,
		Position:     string(player.Position),
		JerseyNumber: player.JerseyNumber,
		CreatedAt:    utils.FormatDateTime(player.CreatedAt),
		UpdatedAt:    utils.FormatDateTime(player.UpdatedAt),
	}

	if player.Team != nil {
		response.TeamName = player.Team.Name
	}

	return response
}

// mapToResponseSimple maps player model to response DTO without team info
func (s *playerService) mapToResponseSimple(player *models.Player) *dto.PlayerResponse {
	return &dto.PlayerResponse{
		ID:           player.ID,
		TeamID:       player.TeamID,
		Name:         player.Name,
		Height:       player.Height,
		Weight:       player.Weight,
		Position:     string(player.Position),
		JerseyNumber: player.JerseyNumber,
		CreatedAt:    utils.FormatDateTime(player.CreatedAt),
		UpdatedAt:    utils.FormatDateTime(player.UpdatedAt),
	}
}
