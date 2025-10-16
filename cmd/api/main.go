package main

import (
	"fmt"
	"log"

	"football-management-api/internal/config"
	"football-management-api/internal/routes"
	"football-management-api/pkg/logger"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Failed to load configuration:", err)
	}

	// Initialize logger
	if err := logger.InitLogger(cfg.Log.File); err != nil {
		log.Fatal("Failed to initialize logger:", err)
	}

	logger.Info(fmt.Sprintf("Starting %s v%s in %s mode",
		cfg.App.Name,
		cfg.App.Version,
		cfg.App.Env,
	))

	// Initialize database
	_, err = config.InitDatabase(cfg)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer config.CloseDatabase()

	logger.Info("Database connected successfully")

	// Setup router
	router := routes.SetupRouter()

	// Start server
	address := cfg.GetServerAddress()
	logger.Info(fmt.Sprintf("Server starting on %s", address))

	fmt.Printf("\n")
	fmt.Printf("========================================\n")
	fmt.Printf("  %s\n", cfg.App.Name)
	fmt.Printf("  Version: %s\n", cfg.App.Version)
	fmt.Printf("  Environment: %s\n", cfg.App.Env)
	fmt.Printf("========================================\n")
	fmt.Printf("  Server running on: http://%s\n", address)
	fmt.Printf("  API Base URL: http://%s/api/v1\n", address)
	fmt.Printf("  Health Check: http://%s/api/v1/health\n", address)
	fmt.Printf("========================================\n")
	fmt.Printf("\n")

	if err := router.Run(address); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
