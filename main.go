package main

import (
	"account-bank-backend/config"
	"account-bank-backend/routes"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// CORS Middleware
func CORSAllow() gin.HandlerFunc {
	cfg := cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}

	cfg.AllowAllOrigins = true
	return cors.New(cfg)
}

func main() {
	gin.SetMode(gin.ReleaseMode)

	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️  Warning: No .env file found")
	}

	// Connect Database
	config.ConnectDatabase()

	// Create Router
	router := gin.New()

	// Middleware
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(CORSAllow())

	// Setup API Routes
	routes.SetupRouter(router, config.DB)

	// Run Server
	log.Printf("🚀 Server running")
	log.Printf("Server running on port %s", os.Getenv("HTTP_PORT"))
	// router.Run(fmt.Sprintf(":%s", os.Getenv("HTTP_PORT")))

	err := router.Run(fmt.Sprintf(":%s", os.Getenv("HTTP_PORT")))
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
