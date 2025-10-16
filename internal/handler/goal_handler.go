package handler

import (
	"football-management-api/internal/dto"
	"football-management-api/internal/service"
	"football-management-api/internal/utils"
	"football-management-api/internal/validator"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GoalHandler struct {
	goalService service.GoalService
}

func NewGoalHandler(goalService service.GoalService) *GoalHandler {
	return &GoalHandler{goalService: goalService}
}

// Create handles creating a new goal
// @Summary Create a new goal
// @Tags goals
// @Accept json
// @Produce json
// @Param goal body dto.CreateGoalRequest true "Goal data"
// @Success 201 {object} dto.Response
// @Router /goals [post]
func (h *GoalHandler) Create(c *gin.Context) {
	var req dto.CreateGoalRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.SendBadRequest(c, "Data tidak valid", utils.FormatValidationError(err))
		return
	}

	if err := validator.ValidateCreateGoal(req); err != nil {
		utils.SendBadRequest(c, "Validasi gagal", err.Error())
		return
	}

	goal, err := h.goalService.Create(req)
	if err != nil {
		utils.SendBadRequest(c, "Gagal membuat data gol", err.Error())
		return
	}

	utils.SendCreated(c, "Data gol berhasil dibuat", goal)
}

// GetByMatchID handles getting goals by match ID
// @Summary Get goals by match ID
// @Tags goals
// @Produce json
// @Param matchId path int true "Match ID"
// @Success 200 {object} dto.Response
// @Router /matches/{matchId}/goals [get]
func (h *GoalHandler) GetByMatchID(c *gin.Context) {
	matchID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.SendBadRequest(c, "Match ID tidak valid", err.Error())
		return
	}

	goals, err := h.goalService.GetByMatchID(matchID)
	if err != nil {
		utils.SendBadRequest(c, "Gagal mengambil data gol", err.Error())
		return
	}

	utils.SendSuccess(c, "Data gol berhasil diambil", goals)
}

// Delete handles deleting a goal
// @Summary Delete a goal
// @Tags goals
// @Produce json
// @Param id path int true "Goal ID"
// @Success 200 {object} dto.Response
// @Router /goals/{id} [delete]
func (h *GoalHandler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.SendBadRequest(c, "ID tidak valid", err.Error())
		return
	}

	err = h.goalService.Delete(id)
	if err != nil {
		utils.SendBadRequest(c, "Gagal menghapus data gol", err.Error())
		return
	}

	utils.SendSuccess(c, "Data gol berhasil dihapus", nil)
}
