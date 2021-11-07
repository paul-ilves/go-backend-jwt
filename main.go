package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/paul-ilves/wanaku-api-go/routes"
	"log"
	"os"
)

func main() {
	checkEnvVars()
	app := fiber.New()
	routes.SetupRoutes(app)
	//repository.OpenDBConnection()

	serverAddress := fmt.Sprintf("%s:%s", os.Getenv("SERVER_HOST"), os.Getenv("SERVER_PORT"))
	log.Fatal(app.Listen(serverAddress))
}

func checkEnvVars() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Could not load .env file")
	}
	envVars := []string{"SERVER_HOST", "SERVER_PORT", "DB_HOST", "DB_PORT", "DB_NAME", "DB_USER", "DB_PASSWORD"}
	for _, envVar := range envVars {
		if os.Getenv(envVar) == "" {
			log.Fatalf("FATAL: environment variable %s is missing", envVar)
		}
	}
}
