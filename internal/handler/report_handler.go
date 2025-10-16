package handler

import (
	"football-management-api/internal/service"
	"football-management-api/internal/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ReportHandler struct {
	reportService service.ReportService
}

func NewReportHandler(reportService service.ReportService) *ReportHandler {
	return &ReportHandler{reportService: reportService}
}

// GetMatchReport handles getting detailed match report
// @Summary Get match report
// @Tags reports
// @Produce json
// @Param matchId path int true "Match ID"
// @Success 200 {object} dto.Response
// @Router /reports/matches/{matchId} [get]
func (h *ReportHandler) GetMatchReport(c *gin.Context) {
	matchID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.SendBadRequest(c, "Match ID tidak valid", err.Error())
		return
	}

	report, err := h.reportService.GetMatchReport(matchID)
	if err != nil {
		utils.SendBadRequest(c, "Gagal mengambil laporan pertandingan", err.Error())
		return
	}

	utils.SendSuccess(c, "Laporan pertandingan berhasil diambil", report)
}

// GetTeamStatistics handles getting team statistics
// @Summary Get team statistics
// @Tags reports
// @Produce json
// @Param teamId path int true "Team ID"
// @Success 200 {object} dto.Response
// @Router /reports/teams/{teamId}/statistics [get]
func (h *ReportHandler) GetTeamStatistics(c *gin.Context) {
	teamID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.SendBadRequest(c, "Team ID tidak valid", err.Error())
		return
	}

	stats, err := h.reportService.GetTeamStatistics(teamID)
	if err != nil {
		utils.SendBadRequest(c, "Gagal mengambil statistik tim", err.Error())
		return
	}

	utils.SendSuccess(c, "Statistik tim berhasil diambil", stats)
}

// GetPlayerStatistics handles getting player statistics
// @Summary Get player statistics
// @Tags reports
// @Produce json
// @Param playerId path int true "Player ID"
// @Success 200 {object} dto.Response
// @Router /reports/players/{playerId}/statistics [get]
func (h *ReportHandler) GetPlayerStatistics(c *gin.Context) {
	playerID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.SendBadRequest(c, "Player ID tidak valid", err.Error())
		return
	}

	stats, err := h.reportService.GetPlayerStatistics(playerID)
	if err != nil {
		utils.SendBadRequest(c, "Gagal mengambil statistik pemain", err.Error())
		return
	}

	utils.SendSuccess(c, "Statistik pemain berhasil diambil", stats)
}

// GetTopScorers handles getting top scorers
// @Summary Get top scorers
// @Tags reports
// @Produce json
// @Param limit query int false "Limit" default(10)
// @Success 200 {object} dto.Response
// @Router /reports/top-scorers [get]
func (h *ReportHandler) GetTopScorers(c *gin.Context) {
	limit := 10
	if limitStr := c.Query("limit"); limitStr != "" {
		if l, err := strconv.Atoi(limitStr); err == nil && l > 0 {
			limit = l
		}
	}

	scorers, err := h.reportService.GetTopScorers(limit)
	if err != nil {
		utils.SendInternalError(c, "Gagal mengambil data top scorer", err.Error())
		return
	}

	utils.SendSuccess(c, "Data top scorer berhasil diambil", scorers)
}
