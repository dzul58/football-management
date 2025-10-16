package routes

import (
	"football-management-api/internal/config"
	"football-management-api/internal/handler"
	"football-management-api/internal/middleware"
	"football-management-api/internal/repository"
	"football-management-api/internal/service"

	"github.com/gin-gonic/gin"
)

// SetupRouter sets up all routes and dependencies
func SetupRouter() *gin.Engine {
	// Set Gin mode
	if config.GlobalConfig.IsDevelopment() {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()

	// Apply middleware
	router.Use(middleware.Recovery())
	router.Use(middleware.Logger())
	router.Use(middleware.CORS())
	router.Use(middleware.ErrorHandler())

	// Initialize repositories
	db := config.GetDB()
	teamRepo := repository.NewTeamRepository(db)
	playerRepo := repository.NewPlayerRepository(db)
	matchRepo := repository.NewMatchRepository(db)
	goalRepo := repository.NewGoalRepository(db)
	reportRepo := repository.NewReportRepository(db)

	// Initialize services
	teamService := service.NewTeamService(teamRepo)
	playerService := service.NewPlayerService(playerRepo, teamRepo)
	matchService := service.NewMatchService(matchRepo, teamRepo, playerRepo, goalRepo)
	goalService := service.NewGoalService(goalRepo, matchRepo, playerRepo)
	reportService := service.NewReportService(reportRepo, matchRepo, teamRepo, playerRepo)

	// Initialize handlers
	teamHandler := handler.NewTeamHandler(teamService)
	playerHandler := handler.NewPlayerHandler(playerService)
	matchHandler := handler.NewMatchHandler(matchService)
	goalHandler := handler.NewGoalHandler(goalService)
	reportHandler := handler.NewReportHandler(reportService)

	// API v1 routes
	v1 := router.Group("/api/v1")
	{
		// Health check
		v1.GET("/health", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"status":  "ok",
				"message": "Football Management API is running",
				"version": config.GlobalConfig.App.Version,
			})
		})

		// Teams routes
		teams := v1.Group("/teams")
		{
			teams.POST("", teamHandler.Create)
			teams.GET("", teamHandler.GetAll)
			teams.GET("/:id", teamHandler.GetByID)
			teams.PUT("/:id", teamHandler.Update)
			teams.DELETE("/:id", teamHandler.Delete)
			teams.GET("/:id/players", playerHandler.GetByTeamID)
		}

		// Players routes
		players := v1.Group("/players")
		{
			players.POST("", playerHandler.Create)
			players.GET("", playerHandler.GetAll)
			players.GET("/:id", playerHandler.GetByID)
			players.PUT("/:id", playerHandler.Update)
			players.DELETE("/:id", playerHandler.Delete)
		}

		// Matches routes
		matches := v1.Group("/matches")
		{
			matches.POST("", matchHandler.Create)
			matches.GET("", matchHandler.GetAll)
			matches.GET("/:id", matchHandler.GetByID)
			matches.PUT("/:id", matchHandler.Update)
			matches.DELETE("/:id", matchHandler.Delete)
			matches.PUT("/:id/result", matchHandler.UpdateResult)
			matches.GET("/:id/goals", goalHandler.GetByMatchID)
		}

		// Goals routes
		goals := v1.Group("/goals")
		{
			goals.POST("", goalHandler.Create)
			goals.DELETE("/:id", goalHandler.Delete)
		}

		// Reports routes
		reports := v1.Group("/reports")
		{
			reports.GET("/matches/:id", reportHandler.GetMatchReport)
			reports.GET("/teams/:id/statistics", reportHandler.GetTeamStatistics)
			reports.GET("/players/:id/statistics", reportHandler.GetPlayerStatistics)
			reports.GET("/top-scorers", reportHandler.GetTopScorers)
		}
	}

	return router
}
