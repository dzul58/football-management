package validator

import (
	"errors"
	"football-management-api/internal/config"
	"football-management-api/internal/dto"
	"football-management-api/internal/utils"
)

// ValidateCreateMatch validates create match request
func ValidateCreateMatch(req dto.CreateMatchRequest) error {
	if req.MatchDate == "" {
		return errors.New("tanggal pertandingan wajib diisi")
	}

	// Validate date format
	if _, err := utils.ParseDate(req.MatchDate); err != nil {
		return errors.New("format tanggal tidak valid. Gunakan format YYYY-MM-DD")
	}

	if req.MatchTime == "" {
		return errors.New("waktu pertandingan wajib diisi")
	}

	// Validate time format
	if _, err := utils.ParseTime(req.MatchTime); err != nil {
		return errors.New("format waktu tidak valid. Gunakan format HH:MM:SS")
	}

	if req.HomeTeamID <= 0 {
		return errors.New("home_team_id tidak valid")
	}

	if req.AwayTeamID <= 0 {
		return errors.New("away_team_id tidak valid")
	}

	if req.HomeTeamID == req.AwayTeamID {
		return errors.New("tim home dan away tidak boleh sama")
	}

	return nil
}

// ValidateUpdateMatch validates update match request
func ValidateUpdateMatch(req dto.UpdateMatchRequest) error {
	if req.MatchDate != "" {
		if _, err := utils.ParseDate(req.MatchDate); err != nil {
			return errors.New("format tanggal tidak valid. Gunakan format YYYY-MM-DD")
		}
	}

	if req.MatchTime != "" {
		if _, err := utils.ParseTime(req.MatchTime); err != nil {
			return errors.New("format waktu tidak valid. Gunakan format HH:MM:SS")
		}
	}

	if req.HomeTeamID != 0 && req.AwayTeamID != 0 {
		if req.HomeTeamID == req.AwayTeamID {
			return errors.New("tim home dan away tidak boleh sama")
		}
	}

	if req.Status != "" {
		if !utils.Contains(config.ValidMatchStatuses(), req.Status) {
			return errors.New("status tidak valid. Pilihan: Scheduled, Completed, Cancelled")
		}
	}

	return nil
}

// ValidateUpdateMatchResult validates update match result request
func ValidateUpdateMatchResult(req dto.UpdateMatchResultRequest) error {
	if req.HomeScore < 0 {
		return errors.New("skor home tidak boleh negatif")
	}

	if req.AwayScore < 0 {
		return errors.New("skor away tidak boleh negatif")
	}

	if len(req.Goals) == 0 {
		if req.HomeScore > 0 || req.AwayScore > 0 {
			return errors.New("detail gol wajib diisi jika ada skor")
		}
	}

	totalGoals := req.HomeScore + req.AwayScore
	if len(req.Goals) != totalGoals {
		return errors.New("jumlah detail gol harus sama dengan total skor")
	}

	for i, goal := range req.Goals {
		if goal.PlayerID <= 0 {
			return errors.New("player_id pada detail gol ke-" + string(rune(i+1)) + " tidak valid")
		}
		if goal.GoalTime == "" {
			return errors.New("waktu gol pada detail gol ke-" + string(rune(i+1)) + " wajib diisi")
		}
	}

	return nil
}
