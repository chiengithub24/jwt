package main

import (
	"instagram-clone/internal/config"
	"instagram-clone/internal/handlers"
	"instagram-clone/internal/repository"
	"instagram-clone/internal/services"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load config
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Cannot load config:", err)
	}

	// Initialize database
	db, err := repository.NewDatabase(cfg)
	if err != nil {
		log.Fatal("Cannot connect to database:", err)
	}

	// Check database connection
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Cannot get database instance:", err)
	}
	err = sqlDB.Ping()
	if err != nil {
		log.Fatal("Cannot ping database:", err)
	}

	// Initialize repositories, services and handlers
	userRepo := repository.NewUserRepository(db)
	authService := services.NewAuthService(userRepo, cfg.JWT.SecretKey)
	authHandler := handlers.NewAuthHandler(authService)

	// Initialize Gin
	r := gin.Default()

	// Routes
	api := r.Group("/api")
	{
		// Auth routes
		auth := api.Group("/auth")
		{
			auth.POST("/register", authHandler.Register)
			auth.POST("/login", authHandler.Login)
		}

		// Health check
		api.GET("/health", func(c *gin.Context) {
			c.JSON(200, gin.H{"status": "ok"})
		})
	}

	// Start server
	serverAddr := ":" + cfg.Server.Port
	log.Printf("Server starting on %s", serverAddr)
	if err := r.Run(serverAddr); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
