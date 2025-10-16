package handler

import (
	"football-management-api/internal/dto"
	"football-management-api/internal/service"
	"football-management-api/internal/utils"
	"football-management-api/internal/validator"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PlayerHandler struct {
	playerService service.PlayerService
}

func NewPlayerHandler(playerService service.PlayerService) *PlayerHandler {
	return &PlayerHandler{playerService: playerService}
}

// Create handles creating a new player
// @Summary Create a new player
// @Tags players
// @Accept json
// @Produce json
// @Param player body dto.CreatePlayerRequest true "Player data"
// @Success 201 {object} dto.Response
// @Router /players [post]
func (h *PlayerHandler) Create(c *gin.Context) {
	var req dto.CreatePlayerRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.SendBadRequest(c, "Data tidak valid", utils.FormatValidationError(err))
		return
	}

	if err := validator.ValidateCreatePlayer(req); err != nil {
		utils.SendBadRequest(c, "Validasi gagal", err.Error())
		return
	}

	player, err := h.playerService.Create(req)
	if err != nil {
		utils.SendBadRequest(c, "Gagal membuat pemain", err.Error())
		return
	}

	utils.SendCreated(c, "Pemain berhasil dibuat", player)
}

// GetByID handles getting a player by ID
// @Summary Get player by ID
// @Tags players
// @Produce json
// @Param id path int true "Player ID"
// @Success 200 {object} dto.Response
// @Router /players/{id} [get]
func (h *PlayerHandler) GetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.SendBadRequest(c, "ID tidak valid", err.Error())
		return
	}

	player, err := h.playerService.GetByID(id)
	if err != nil {
		utils.SendNotFound(c, "Pemain tidak ditemukan", err.Error())
		return
	}

	utils.SendSuccess(c, "Pemain ditemukan", player)
}

// GetAll handles getting all players with pagination
// @Summary Get all players
// @Tags players
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Items per page" default(10)
// @Success 200 {object} dto.PaginatedResponse
// @Router /players [get]
func (h *PlayerHandler) GetAll(c *gin.Context) {
	page, limit := utils.GetPaginationParams(c)

	players, meta, err := h.playerService.GetAll(page, limit)
	if err != nil {
		utils.SendInternalError(c, "Gagal mengambil data pemain", err.Error())
		return
	}

	utils.SendPaginated(c, "Data pemain berhasil diambil", players, meta)
}

// GetByTeamID handles getting players by team ID
// @Summary Get players by team ID
// @Tags players
// @Produce json
// @Param teamId path int true "Team ID"
// @Success 200 {object} dto.Response
// @Router /teams/{teamId}/players [get]
func (h *PlayerHandler) GetByTeamID(c *gin.Context) {
	teamID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.SendBadRequest(c, "Team ID tidak valid", err.Error())
		return
	}

	players, err := h.playerService.GetByTeamID(teamID)
	if err != nil {
		utils.SendBadRequest(c, "Gagal mengambil data pemain", err.Error())
		return
	}

	utils.SendSuccess(c, "Data pemain berhasil diambil", players)
}

// Update handles updating a player
// @Summary Update a player
// @Tags players
// @Accept json
// @Produce json
// @Param id path int true "Player ID"
// @Param player body dto.UpdatePlayerRequest true "Player data"
// @Success 200 {object} dto.Response
// @Router /players/{id} [put]
func (h *PlayerHandler) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.SendBadRequest(c, "ID tidak valid", err.Error())
		return
	}

	var req dto.UpdatePlayerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.SendBadRequest(c, "Data tidak valid", utils.FormatValidationError(err))
		return
	}

	if err := validator.ValidateUpdatePlayer(req); err != nil {
		utils.SendBadRequest(c, "Validasi gagal", err.Error())
		return
	}

	player, err := h.playerService.Update(id, req)
	if err != nil {
		utils.SendBadRequest(c, "Gagal memperbarui pemain", err.Error())
		return
	}

	utils.SendSuccess(c, "Pemain berhasil diperbarui", player)
}

// Delete handles deleting a player
// @Summary Delete a player
// @Tags players
// @Produce json
// @Param id path int true "Player ID"
// @Success 200 {object} dto.Response
// @Router /players/{id} [delete]
func (h *PlayerHandler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.SendBadRequest(c, "ID tidak valid", err.Error())
		return
	}

	err = h.playerService.Delete(id)
	if err != nil {
		utils.SendBadRequest(c, "Gagal menghapus pemain", err.Error())
		return
	}

	utils.SendSuccess(c, "Pemain berhasil dihapus", nil)
}
