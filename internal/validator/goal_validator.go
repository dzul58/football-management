package validator

import (
	"errors"
	"football-management-api/internal/dto"
)

// ValidateCreateGoal validates create goal request
func ValidateCreateGoal(req dto.CreateGoalRequest) error {
	if req.MatchID <= 0 {
		return errors.New("match_id tidak valid")
	}

	if req.PlayerID <= 0 {
		return errors.New("player_id tidak valid")
	}

	if req.GoalTime == "" {
		return errors.New("waktu gol wajib diisi")
	}

	return nil
}
