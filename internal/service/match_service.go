package service

import (
	"errors"
	"football-management-api/internal/dto"
	"football-management-api/internal/models"
	"football-management-api/internal/repository"
	"football-management-api/internal/utils"
)

type MatchService interface {
	Create(req dto.CreateMatchRequest) (*dto.MatchResponse, error)
	GetByID(id int) (*dto.MatchResponse, error)
	GetAll(page, limit int) ([]dto.MatchResponse, dto.PaginationMeta, error)
	Update(id int, req dto.UpdateMatchRequest) (*dto.MatchResponse, error)
	UpdateResult(id int, req dto.UpdateMatchResultRequest) (*dto.MatchResponse, error)
	Delete(id int) error
}

type matchService struct {
	matchRepo  repository.MatchRepository
	teamRepo   repository.TeamRepository
	playerRepo repository.PlayerRepository
	goalRepo   repository.GoalRepository
}

func NewMatchService(
	matchRepo repository.MatchRepository,
	teamRepo repository.TeamRepository,
	playerRepo repository.PlayerRepository,
	goalRepo repository.GoalRepository,
) MatchService {
	return &matchService{
		matchRepo:  matchRepo,
		teamRepo:   teamRepo,
		playerRepo: playerRepo,
		goalRepo:   goalRepo,
	}
}

// Create creates a new match
func (s *matchService) Create(req dto.CreateMatchRequest) (*dto.MatchResponse, error) {
	// Validate teams exist
	_, err := s.teamRepo.FindByID(req.HomeTeamID)
	if err != nil {
		return nil, errors.New("tim home tidak ditemukan")
	}

	_, err = s.teamRepo.FindByID(req.AwayTeamID)
	if err != nil {
		return nil, errors.New("tim away tidak ditemukan")
	}

	match := &models.Match{
		MatchDate:  req.MatchDate,
		MatchTime:  req.MatchTime,
		HomeTeamID: req.HomeTeamID,
		AwayTeamID: req.AwayTeamID,
		Status:     models.StatusScheduled,
	}

	err = s.matchRepo.Create(match)
	if err != nil {
		return nil, err
	}

	// Get match with team info
	createdMatch, err := s.matchRepo.FindByID(match.ID)
	if err != nil {
		return nil, err
	}

	return s.mapToResponse(createdMatch), nil
}

// GetByID gets a match by ID
func (s *matchService) GetByID(id int) (*dto.MatchResponse, error) {
	match, err := s.matchRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	return s.mapToResponse(match), nil
}

// GetAll gets all matches with pagination
func (s *matchService) GetAll(page, limit int) ([]dto.MatchResponse, dto.PaginationMeta, error) {
	offset := utils.CalculateOffset(page, limit)

	matches, total, err := s.matchRepo.FindAll(limit, offset)
	if err != nil {
		return nil, dto.PaginationMeta{}, err
	}

	var responses []dto.MatchResponse
	for _, match := range matches {
		responses = append(responses, *s.mapToResponse(&match))
	}

	meta := dto.PaginationMeta{
		CurrentPage: page,
		PerPage:     limit,
		Total:       total,
		TotalPages:  utils.CalculateTotalPages(total, limit),
	}

	return responses, meta, nil
}

// Update updates a match
func (s *matchService) Update(id int, req dto.UpdateMatchRequest) (*dto.MatchResponse, error) {
	// Get existing match
	existingMatch, err := s.matchRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	// Update fields if provided
	if req.MatchDate != "" {
		existingMatch.MatchDate = req.MatchDate
	}

	if req.MatchTime != "" {
		existingMatch.MatchTime = req.MatchTime
	}

	if req.HomeTeamID != 0 {
		// Validate team exists
		_, err := s.teamRepo.FindByID(req.HomeTeamID)
		if err != nil {
			return nil, errors.New("tim home tidak ditemukan")
		}
		existingMatch.HomeTeamID = req.HomeTeamID
	}

	if req.AwayTeamID != 0 {
		// Validate team exists
		_, err := s.teamRepo.FindByID(req.AwayTeamID)
		if err != nil {
			return nil, errors.New("tim away tidak ditemukan")
		}
		existingMatch.AwayTeamID = req.AwayTeamID
	}

	if req.Status != "" {
		existingMatch.Status = models.MatchStatus(req.Status)
	}

	err = s.matchRepo.Update(id, existingMatch)
	if err != nil {
		return nil, err
	}

	// Get updated match
	updatedMatch, err := s.matchRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	return s.mapToResponse(updatedMatch), nil
}

// UpdateResult updates match result with goals
func (s *matchService) UpdateResult(id int, req dto.UpdateMatchResultRequest) (*dto.MatchResponse, error) {
	// Get existing match
	match, err := s.matchRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	// Validate all players and count goals per team
	homeGoals := 0
	awayGoals := 0

	for _, goalInput := range req.Goals {
		player, err := s.playerRepo.FindByID(goalInput.PlayerID)
		if err != nil {
			return nil, errors.New("pemain dengan ID " + string(rune(goalInput.PlayerID)) + " tidak ditemukan")
		}

		// Check if player belongs to one of the teams in the match
		if player.TeamID != match.HomeTeamID && player.TeamID != match.AwayTeamID {
			return nil, errors.New("pemain " + player.Name + " tidak bermain untuk tim yang bertanding")
		}

		// Count goals per team
		if player.TeamID == match.HomeTeamID {
			homeGoals++
		} else {
			awayGoals++
		}
	}

	// Validate goal count matches the scores
	if homeGoals != req.HomeScore || awayGoals != req.AwayScore {
		return nil, errors.New("jumlah gol per tim tidak sesuai dengan skor yang diberikan")
	}

	// Delete existing goals for this match
	err = s.goalRepo.DeleteByMatchID(id)
	if err != nil {
		return nil, err
	}

	// Create new goals
	for _, goalInput := range req.Goals {
		goal := &models.Goal{
			MatchID:  id,
			PlayerID: goalInput.PlayerID,
			GoalTime: goalInput.GoalTime,
		}
		err = s.goalRepo.Create(goal)
		if err != nil {
			return nil, err
		}
	}

	// Update match result
	err = s.matchRepo.UpdateResult(id, req.HomeScore, req.AwayScore, models.StatusCompleted)
	if err != nil {
		return nil, err
	}

	// Get updated match
	updatedMatch, err := s.matchRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	return s.mapToResponse(updatedMatch), nil
}

// Delete deletes a match
func (s *matchService) Delete(id int) error {
	return s.matchRepo.Delete(id)
}

// mapToResponse maps match model to response DTO
func (s *matchService) mapToResponse(match *models.Match) *dto.MatchResponse {
	response := &dto.MatchResponse{
		ID:         match.ID,
		MatchDate:  match.MatchDate,
		MatchTime:  match.MatchTime,
		HomeTeamID: match.HomeTeamID,
		AwayTeamID: match.AwayTeamID,
		HomeScore:  utils.NullInt32ToIntPtr(match.HomeScore),
		AwayScore:  utils.NullInt32ToIntPtr(match.AwayScore),
		Status:     string(match.Status),
		CreatedAt:  utils.FormatDateTime(match.CreatedAt),
		UpdatedAt:  utils.FormatDateTime(match.UpdatedAt),
	}

	if match.HomeTeam != nil {
		response.HomeTeamName = match.HomeTeam.Name
	}

	if match.AwayTeam != nil {
		response.AwayTeamName = match.AwayTeam.Name
	}

	return response
}
