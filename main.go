package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/paul-ilves/wanaku-api-go/config"
	"github.com/paul-ilves/wanaku-api-go/routes"
	"log"
	"os"
)

func main() {
	app := fiber.New()
	routes.SetupRoutes(app)
	log.Fatal(app.Listen(config.C.ServerHostPort))
}

func init() {
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

	serverHostPort := fmt.Sprintf("%s:%s", os.Getenv("SERVER_HOST"), os.Getenv("SERVER_PORT"))
	dbConnectionURL := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"))

	config.C = &config.Config{
		ServerHostPort:  serverHostPort,
		DBConnectionURL: dbConnectionURL,
	}
}
