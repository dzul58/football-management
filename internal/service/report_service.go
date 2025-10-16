package service

import (
	"errors"
	"football-management-api/internal/models"
	"football-management-api/internal/repository"
)

type ReportService interface {
	GetMatchReport(matchID int) (*models.MatchReport, error)
	GetTeamStatistics(teamID int) (*models.TeamStatistics, error)
	GetPlayerStatistics(playerID int) (*models.PlayerStatistics, error)
	GetTopScorers(limit int) ([]models.PlayerStatistics, error)
}

type reportService struct {
	reportRepo repository.ReportRepository
	matchRepo  repository.MatchRepository
	teamRepo   repository.TeamRepository
	playerRepo repository.PlayerRepository
}

func NewReportService(
	reportRepo repository.ReportRepository,
	matchRepo repository.MatchRepository,
	teamRepo repository.TeamRepository,
	playerRepo repository.PlayerRepository,
) ReportService {
	return &reportService{
		reportRepo: reportRepo,
		matchRepo:  matchRepo,
		teamRepo:   teamRepo,
		playerRepo: playerRepo,
	}
}

// GetMatchReport gets detailed match report
func (s *reportService) GetMatchReport(matchID int) (*models.MatchReport, error) {
	// Validate match exists and is completed
	match, err := s.matchRepo.FindByID(matchID)
	if err != nil {
		return nil, errors.New("pertandingan tidak ditemukan")
	}

	if match.Status != models.StatusCompleted {
		return nil, errors.New("laporan hanya tersedia untuk pertandingan yang sudah selesai")
	}

	return s.reportRepo.GetMatchReport(matchID)
}

// GetTeamStatistics gets team statistics
func (s *reportService) GetTeamStatistics(teamID int) (*models.TeamStatistics, error) {
	// Validate team exists
	_, err := s.teamRepo.FindByID(teamID)
	if err != nil {
		return nil, errors.New("tim tidak ditemukan")
	}

	return s.reportRepo.GetTeamStatistics(teamID)
}

// GetPlayerStatistics gets player goal statistics
func (s *reportService) GetPlayerStatistics(playerID int) (*models.PlayerStatistics, error) {
	// Validate player exists
	_, err := s.playerRepo.FindByID(playerID)
	if err != nil {
		return nil, errors.New("pemain tidak ditemukan")
	}

	return s.reportRepo.GetPlayerStatistics(playerID)
}

// GetTopScorers gets top scorers
func (s *reportService) GetTopScorers(limit int) ([]models.PlayerStatistics, error) {
	if limit <= 0 {
		limit = 10
	}
	if limit > 100 {
		limit = 100
	}

	return s.reportRepo.GetTopScorers(limit)
}
