package validator

import (
	"errors"
	"football-management-api/internal/config"
	"football-management-api/internal/dto"
	"football-management-api/internal/utils"
)

// ValidateCreatePlayer validates create player request
func ValidateCreatePlayer(req dto.CreatePlayerRequest) error {
	if req.TeamID <= 0 {
		return errors.New("team_id tidak valid")
	}

	if req.Name == "" {
		return errors.New("nama pemain wajib diisi")
	}

	if req.Height < 100 || req.Height > 250 {
		return errors.New("tinggi badan harus antara 100-250 cm")
	}

	if req.Weight < 30 || req.Weight > 200 {
		return errors.New("berat badan harus antara 30-200 kg")
	}

	if !utils.Contains(config.ValidPlayerPositions(), req.Position) {
		return errors.New("posisi pemain tidak valid. Pilihan: Penyerang, Gelandang, Bertahan, Penjaga Gawang")
	}

	if req.JerseyNumber < 1 || req.JerseyNumber > 99 {
		return errors.New("nomor punggung harus antara 1-99")
	}

	return nil
}

// ValidateUpdatePlayer validates update player request
func ValidateUpdatePlayer(req dto.UpdatePlayerRequest) error {
	if req.TeamID < 0 {
		return errors.New("team_id tidak valid")
	}

	if req.Height != 0 {
		if req.Height < 100 || req.Height > 250 {
			return errors.New("tinggi badan harus antara 100-250 cm")
		}
	}

	if req.Weight != 0 {
		if req.Weight < 30 || req.Weight > 200 {
			return errors.New("berat badan harus antara 30-200 kg")
		}
	}

	if req.Position != "" {
		if !utils.Contains(config.ValidPlayerPositions(), req.Position) {
			return errors.New("posisi pemain tidak valid. Pilihan: Penyerang, Gelandang, Bertahan, Penjaga Gawang")
		}
	}

	if req.JerseyNumber != 0 {
		if req.JerseyNumber < 1 || req.JerseyNumber > 99 {
			return errors.New("nomor punggung harus antara 1-99")
		}
	}

	return nil
}
