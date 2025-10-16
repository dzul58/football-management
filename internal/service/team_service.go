package service

import (
	"errors"
	"football-management-api/internal/dto"
	"football-management-api/internal/models"
	"football-management-api/internal/repository"
	"football-management-api/internal/utils"
)

type TeamService interface {
	Create(req dto.CreateTeamRequest) (*dto.TeamResponse, error)
	GetByID(id int) (*dto.TeamResponse, error)
	GetAll(page, limit int) ([]dto.TeamResponse, dto.PaginationMeta, error)
	Update(id int, req dto.UpdateTeamRequest) (*dto.TeamResponse, error)
	Delete(id int) error
}

type teamService struct {
	teamRepo repository.TeamRepository
}

func NewTeamService(teamRepo repository.TeamRepository) TeamService {
	return &teamService{teamRepo: teamRepo}
}

// Create creates a new team
func (s *teamService) Create(req dto.CreateTeamRequest) (*dto.TeamResponse, error) {
	// Check if team name already exists
	existingTeam, err := s.teamRepo.FindByName(req.Name)
	if err != nil {
		return nil, err
	}
	if existingTeam != nil {
		return nil, errors.New("nama tim sudah digunakan")
	}

	team := &models.Team{
		Name:        req.Name,
		LogoURL:     utils.StringToNullString(req.LogoURL),
		FoundedYear: req.FoundedYear,
		HomeAddress: req.HomeAddress,
		HomeCity:    req.HomeCity,
	}

	err = s.teamRepo.Create(team)
	if err != nil {
		return nil, err
	}

	return s.mapToResponse(team), nil
}

// GetByID gets a team by ID
func (s *teamService) GetByID(id int) (*dto.TeamResponse, error) {
	team, err := s.teamRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	return s.mapToResponse(team), nil
}

// GetAll gets all teams with pagination
func (s *teamService) GetAll(page, limit int) ([]dto.TeamResponse, dto.PaginationMeta, error) {
	offset := utils.CalculateOffset(page, limit)

	teams, total, err := s.teamRepo.FindAll(limit, offset)
	if err != nil {
		return nil, dto.PaginationMeta{}, err
	}

	var responses []dto.TeamResponse
	for _, team := range teams {
		responses = append(responses, *s.mapToResponse(&team))
	}

	meta := dto.PaginationMeta{
		CurrentPage: page,
		PerPage:     limit,
		Total:       total,
		TotalPages:  utils.CalculateTotalPages(total, limit),
	}

	return responses, meta, nil
}

// Update updates a team
func (s *teamService) Update(id int, req dto.UpdateTeamRequest) (*dto.TeamResponse, error) {
	// Get existing team
	existingTeam, err := s.teamRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	// Check if name is being changed and already exists
	if req.Name != "" && req.Name != existingTeam.Name {
		teamWithName, err := s.teamRepo.FindByName(req.Name)
		if err != nil {
			return nil, err
		}
		if teamWithName != nil {
			return nil, errors.New("nama tim sudah digunakan")
		}
		existingTeam.Name = req.Name
	}

	// Update fields if provided
	if req.LogoURL != "" {
		existingTeam.LogoURL = utils.StringToNullString(req.LogoURL)
	}
	if req.FoundedYear != 0 {
		existingTeam.FoundedYear = req.FoundedYear
	}
	if req.HomeAddress != "" {
		existingTeam.HomeAddress = req.HomeAddress
	}
	if req.HomeCity != "" {
		existingTeam.HomeCity = req.HomeCity
	}

	err = s.teamRepo.Update(id, existingTeam)
	if err != nil {
		return nil, err
	}

	// Get updated team
	updatedTeam, err := s.teamRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	return s.mapToResponse(updatedTeam), nil
}

// Delete deletes a team
func (s *teamService) Delete(id int) error {
	return s.teamRepo.Delete(id)
}

// mapToResponse maps team model to response DTO
func (s *teamService) mapToResponse(team *models.Team) *dto.TeamResponse {
	return &dto.TeamResponse{
		ID:          team.ID,
		Name:        team.Name,
		LogoURL:     utils.NullStringToString(team.LogoURL),
		FoundedYear: team.FoundedYear,
		HomeAddress: team.HomeAddress,
		HomeCity:    team.HomeCity,
		CreatedAt:   utils.FormatDateTime(team.CreatedAt),
		UpdatedAt:   utils.FormatDateTime(team.UpdatedAt),
	}
}
