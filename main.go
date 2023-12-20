package main

import (
	"REST-API/helpers"
	userControllers "REST-API/model/user/controllers"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	env := os.Getenv("ENV")
	currentActiveEnv := helpers.GetCurrentEnv(env)
	envFile := currentActiveEnv + ".env"
	if err := godotenv.Load(envFile); err != nil {
		log.Fatal("Error occurred while loading env file:", err)
	}

	// Create a Gin Server
	server := gin.New()

	// Add CORS middleware to allow cross-origin requests
	server.Use(cors.Default())

	// Use the default Gin logger middleware
	server.Use(gin.Logger())

	// Registering Routes
	basepath := server.Group("/")
	userControllers.RegisterUserRoutes(basepath)
	// Start the Gin server
	port := os.Getenv("APP_PORT")
	log.Fatal(server.Run(":" + port))
}
