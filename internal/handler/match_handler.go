package handler

import (
	"football-management-api/internal/dto"
	"football-management-api/internal/service"
	"football-management-api/internal/utils"
	"football-management-api/internal/validator"
	"strconv"

	"github.com/gin-gonic/gin"
)

type MatchHandler struct {
	matchService service.MatchService
}

func NewMatchHandler(matchService service.MatchService) *MatchHandler {
	return &MatchHandler{matchService: matchService}
}

// Create handles creating a new match
// @Summary Create a new match
// @Tags matches
// @Accept json
// @Produce json
// @Param match body dto.CreateMatchRequest true "Match data"
// @Success 201 {object} dto.Response
// @Router /matches [post]
func (h *MatchHandler) Create(c *gin.Context) {
	var req dto.CreateMatchRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.SendBadRequest(c, "Data tidak valid", utils.FormatValidationError(err))
		return
	}

	if err := validator.ValidateCreateMatch(req); err != nil {
		utils.SendBadRequest(c, "Validasi gagal", err.Error())
		return
	}

	match, err := h.matchService.Create(req)
	if err != nil {
		utils.SendBadRequest(c, "Gagal membuat pertandingan", err.Error())
		return
	}

	utils.SendCreated(c, "Pertandingan berhasil dibuat", match)
}

// GetByID handles getting a match by ID
// @Summary Get match by ID
// @Tags matches
// @Produce json
// @Param id path int true "Match ID"
// @Success 200 {object} dto.Response
// @Router /matches/{id} [get]
func (h *MatchHandler) GetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.SendBadRequest(c, "ID tidak valid", err.Error())
		return
	}

	match, err := h.matchService.GetByID(id)
	if err != nil {
		utils.SendNotFound(c, "Pertandingan tidak ditemukan", err.Error())
		return
	}

	utils.SendSuccess(c, "Pertandingan ditemukan", match)
}

// GetAll handles getting all matches with pagination
// @Summary Get all matches
// @Tags matches
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Items per page" default(10)
// @Success 200 {object} dto.PaginatedResponse
// @Router /matches [get]
func (h *MatchHandler) GetAll(c *gin.Context) {
	page, limit := utils.GetPaginationParams(c)

	matches, meta, err := h.matchService.GetAll(page, limit)
	if err != nil {
		utils.SendInternalError(c, "Gagal mengambil data pertandingan", err.Error())
		return
	}

	utils.SendPaginated(c, "Data pertandingan berhasil diambil", matches, meta)
}

// Update handles updating a match
// @Summary Update a match
// @Tags matches
// @Accept json
// @Produce json
// @Param id path int true "Match ID"
// @Param match body dto.UpdateMatchRequest true "Match data"
// @Success 200 {object} dto.Response
// @Router /matches/{id} [put]
func (h *MatchHandler) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.SendBadRequest(c, "ID tidak valid", err.Error())
		return
	}

	var req dto.UpdateMatchRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.SendBadRequest(c, "Data tidak valid", utils.FormatValidationError(err))
		return
	}

	if err := validator.ValidateUpdateMatch(req); err != nil {
		utils.SendBadRequest(c, "Validasi gagal", err.Error())
		return
	}

	match, err := h.matchService.Update(id, req)
	if err != nil {
		utils.SendBadRequest(c, "Gagal memperbarui pertandingan", err.Error())
		return
	}

	utils.SendSuccess(c, "Pertandingan berhasil diperbarui", match)
}

// UpdateResult handles updating match result
// @Summary Update match result
// @Tags matches
// @Accept json
// @Produce json
// @Param id path int true "Match ID"
// @Param result body dto.UpdateMatchResultRequest true "Match result data"
// @Success 200 {object} dto.Response
// @Router /matches/{id}/result [put]
func (h *MatchHandler) UpdateResult(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.SendBadRequest(c, "ID tidak valid", err.Error())
		return
	}

	var req dto.UpdateMatchResultRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.SendBadRequest(c, "Data tidak valid", utils.FormatValidationError(err))
		return
	}

	if err := validator.ValidateUpdateMatchResult(req); err != nil {
		utils.SendBadRequest(c, "Validasi gagal", err.Error())
		return
	}

	match, err := h.matchService.UpdateResult(id, req)
	if err != nil {
		utils.SendBadRequest(c, "Gagal memperbarui hasil pertandingan", err.Error())
		return
	}

	utils.SendSuccess(c, "Hasil pertandingan berhasil diperbarui", match)
}

// Delete handles deleting a match
// @Summary Delete a match
// @Tags matches
// @Produce json
// @Param id path int true "Match ID"
// @Success 200 {object} dto.Response
// @Router /matches/{id} [delete]
func (h *MatchHandler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.SendBadRequest(c, "ID tidak valid", err.Error())
		return
	}

	err = h.matchService.Delete(id)
	if err != nil {
		utils.SendBadRequest(c, "Gagal menghapus pertandingan", err.Error())
		return
	}

	utils.SendSuccess(c, "Pertandingan berhasil dihapus", nil)
}
