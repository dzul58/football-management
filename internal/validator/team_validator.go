package validator

import (
	"errors"
	"football-management-api/internal/dto"
)

// ValidateCreateTeam validates create team request
func ValidateCreateTeam(req dto.CreateTeamRequest) error {
	if req.Name == "" {
		return errors.New("nama tim wajib diisi")
	}

	if req.FoundedYear < 1800 || req.FoundedYear > 2100 {
		return errors.New("tahun berdiri tidak valid")
	}

	if req.HomeAddress == "" {
		return errors.New("alamat markas wajib diisi")
	}

	if req.HomeCity == "" {
		return errors.New("kota markas wajib diisi")
	}

	return nil
}

// ValidateUpdateTeam validates update team request
func ValidateUpdateTeam(req dto.UpdateTeamRequest) error {
	if req.FoundedYear != 0 {
		if req.FoundedYear < 1800 || req.FoundedYear > 2100 {
			return errors.New("tahun berdiri tidak valid")
		}
	}

	return nil
}
