package initialize

import (
	"log"
	"os"

	"server/docs"

	"github.com/joho/godotenv"
)

func init() {
	// To load our environmental variables.
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}

func App() {
	serviceRegister()

	// Programatically set swagger info
	docs.SwaggerInfo.Title = "Swagger Go-DDD-Layout API"
	docs.SwaggerInfo.Description = "This is a simple Go-DDD-Layout project where we have implemented basic CRUD operations for users. \nIt differentiates between administrator roles and regular user permissions. \nThrough this project, you can learn how to build an engineered Go project and explore the new perspectives that domain-driven design architecture brings to us."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "go-ddd-layout.com"
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	appPort := os.Getenv("APP_PORT")
	if appPort == "" {
		appPort = "8080" // default port
	}

	log.Fatal(routeEngine().Run(":" + appPort))
}
