package service

import (
	"errors"
	"football-management-api/internal/dto"
	"football-management-api/internal/models"
	"football-management-api/internal/repository"
	"football-management-api/internal/utils"
)

type GoalService interface {
	Create(req dto.CreateGoalRequest) (*dto.GoalResponse, error)
	GetByMatchID(matchID int) ([]dto.GoalResponse, error)
	Delete(id int) error
}

type goalService struct {
	goalRepo   repository.GoalRepository
	matchRepo  repository.MatchRepository
	playerRepo repository.PlayerRepository
}

func NewGoalService(
	goalRepo repository.GoalRepository,
	matchRepo repository.MatchRepository,
	playerRepo repository.PlayerRepository,
) GoalService {
	return &goalService{
		goalRepo:   goalRepo,
		matchRepo:  matchRepo,
		playerRepo: playerRepo,
	}
}

// Create creates a new goal
func (s *goalService) Create(req dto.CreateGoalRequest) (*dto.GoalResponse, error) {
	// Validate match exists
	match, err := s.matchRepo.FindByID(req.MatchID)
	if err != nil {
		return nil, errors.New("pertandingan tidak ditemukan")
	}

	// Validate player exists
	player, err := s.playerRepo.FindByID(req.PlayerID)
	if err != nil {
		return nil, errors.New("pemain tidak ditemukan")
	}

	// Validate player is in one of the teams playing
	if player.TeamID != match.HomeTeamID && player.TeamID != match.AwayTeamID {
		return nil, errors.New("pemain tidak bermain dalam pertandingan ini")
	}

	goal := &models.Goal{
		MatchID:  req.MatchID,
		PlayerID: req.PlayerID,
		GoalTime: req.GoalTime,
	}

	err = s.goalRepo.Create(goal)
	if err != nil {
		return nil, err
	}

	// Get created goal with details
	createdGoal, err := s.goalRepo.FindByID(goal.ID)
	if err != nil {
		return nil, err
	}

	return s.mapToResponse(createdGoal), nil
}

// GetByMatchID gets all goals in a match
func (s *goalService) GetByMatchID(matchID int) ([]dto.GoalResponse, error) {
	// Validate match exists
	_, err := s.matchRepo.FindByID(matchID)
	if err != nil {
		return nil, errors.New("pertandingan tidak ditemukan")
	}

	goals, err := s.goalRepo.FindByMatchID(matchID)
	if err != nil {
		return nil, err
	}

	var responses []dto.GoalResponse
	for _, goal := range goals {
		responses = append(responses, *s.mapToResponse(&goal))
	}

	return responses, nil
}

// Delete deletes a goal
func (s *goalService) Delete(id int) error {
	return s.goalRepo.Delete(id)
}

// mapToResponse maps goal model to response DTO
func (s *goalService) mapToResponse(goal *models.Goal) *dto.GoalResponse {
	response := &dto.GoalResponse{
		ID:        goal.ID,
		MatchID:   goal.MatchID,
		PlayerID:  goal.PlayerID,
		GoalTime:  goal.GoalTime,
		CreatedAt: utils.FormatDateTime(goal.CreatedAt),
	}

	if goal.Player != nil {
		response.PlayerName = goal.Player.Name
		if goal.Player.Team != nil {
			response.TeamName = goal.Player.Team.Name
		}
	}

	return response
}
