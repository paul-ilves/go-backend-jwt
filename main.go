package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/paul-ilves/wanaku-api-go/config"
	"github.com/paul-ilves/wanaku-api-go/repository"
	"github.com/paul-ilves/wanaku-api-go/routes"
	"github.com/paul-ilves/wanaku-api-go/utils"
	"log"
	"os"
	"sync"
)

func main() {
	app := fiber.New()
	routes.SetupRoutes(app)
	log.Fatal(app.Listen(config.C.ServerUrl()))
}

// init loads the env variables from the .env file, checks if any of them are missing and passes them to config.C singleton for global use.
// it also prevents the application from starting if DB connection cannot be made
func init() {
	once := sync.Once{}
	once.Do(func() {
		var err error
		if err := godotenv.Load(); err != nil {
			log.Fatal("Could not load .env file")
		}
		envVars := []string{"SERVER_HOST", "SERVER_PORT", "DB_HOST", "DB_PORT", "DB_NAME", "DB_USER", "DB_PASSWORD", "JWT_SECRET", "REFRESH_TOKEN_LIFETIME_HOURS", "ACCESS_TOKEN_LIFETIME_MINUTES"}
		for _, envVar := range envVars {
			if os.Getenv(envVar) == "" {
				log.Fatalf("FATAL: environment variable %v is missing", envVar)
			}
		}
		config.C = new(config.Config)
		if err = envconfig.Process("", config.C); err != nil {
			utils.Error("Problem with populating config.C with env variables: " + err.Error())
		}

		//check if db connection can be established
		repository.OpenDBConnection()
		defer repository.CloseDBConnection()
	})
}
