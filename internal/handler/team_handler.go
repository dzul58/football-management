package handler

import (
	"football-management-api/internal/dto"
	"football-management-api/internal/service"
	"football-management-api/internal/utils"
	"football-management-api/internal/validator"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TeamHandler struct {
	teamService service.TeamService
}

func NewTeamHandler(teamService service.TeamService) *TeamHandler {
	return &TeamHandler{teamService: teamService}
}

// Create handles creating a new team
// @Summary Create a new team
// @Tags teams
// @Accept json
// @Produce json
// @Param team body dto.CreateTeamRequest true "Team data"
// @Success 201 {object} dto.Response
// @Router /teams [post]
func (h *TeamHandler) Create(c *gin.Context) {
	var req dto.CreateTeamRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.SendBadRequest(c, "Data tidak valid", utils.FormatValidationError(err))
		return
	}

	if err := validator.ValidateCreateTeam(req); err != nil {
		utils.SendBadRequest(c, "Validasi gagal", err.Error())
		return
	}

	team, err := h.teamService.Create(req)
	if err != nil {
		utils.SendBadRequest(c, "Gagal membuat tim", err.Error())
		return
	}

	utils.SendCreated(c, "Tim berhasil dibuat", team)
}

// GetByID handles getting a team by ID
// @Summary Get team by ID
// @Tags teams
// @Produce json
// @Param id path int true "Team ID"
// @Success 200 {object} dto.Response
// @Router /teams/{id} [get]
func (h *TeamHandler) GetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.SendBadRequest(c, "ID tidak valid", err.Error())
		return
	}

	team, err := h.teamService.GetByID(id)
	if err != nil {
		utils.SendNotFound(c, "Tim tidak ditemukan", err.Error())
		return
	}

	utils.SendSuccess(c, "Tim ditemukan", team)
}

// GetAll handles getting all teams with pagination
// @Summary Get all teams
// @Tags teams
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Items per page" default(10)
// @Success 200 {object} dto.PaginatedResponse
// @Router /teams [get]
func (h *TeamHandler) GetAll(c *gin.Context) {
	page, limit := utils.GetPaginationParams(c)

	teams, meta, err := h.teamService.GetAll(page, limit)
	if err != nil {
		utils.SendInternalError(c, "Gagal mengambil data tim", err.Error())
		return
	}

	utils.SendPaginated(c, "Data tim berhasil diambil", teams, meta)
}

// Update handles updating a team
// @Summary Update a team
// @Tags teams
// @Accept json
// @Produce json
// @Param id path int true "Team ID"
// @Param team body dto.UpdateTeamRequest true "Team data"
// @Success 200 {object} dto.Response
// @Router /teams/{id} [put]
func (h *TeamHandler) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.SendBadRequest(c, "ID tidak valid", err.Error())
		return
	}

	var req dto.UpdateTeamRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.SendBadRequest(c, "Data tidak valid", utils.FormatValidationError(err))
		return
	}

	if err := validator.ValidateUpdateTeam(req); err != nil {
		utils.SendBadRequest(c, "Validasi gagal", err.Error())
		return
	}

	team, err := h.teamService.Update(id, req)
	if err != nil {
		utils.SendBadRequest(c, "Gagal memperbarui tim", err.Error())
		return
	}

	utils.SendSuccess(c, "Tim berhasil diperbarui", team)
}

// Delete handles deleting a team
// @Summary Delete a team
// @Tags teams
// @Produce json
// @Param id path int true "Team ID"
// @Success 200 {object} dto.Response
// @Router /teams/{id} [delete]
func (h *TeamHandler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.SendBadRequest(c, "ID tidak valid", err.Error())
		return
	}

	err = h.teamService.Delete(id)
	if err != nil {
		utils.SendBadRequest(c, "Gagal menghapus tim", err.Error())
		return
	}

	utils.SendSuccess(c, "Tim berhasil dihapus", nil)
}
