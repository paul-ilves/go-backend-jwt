package config

import (
	"fmt"
	"time"
)

// C contains all the config parameters for the App, such as Server parameters, DB, secrets and etc.
var C *Config = nil

// Config defines the config parameters of the App which are extracted from the environment variables.
type Config struct {
	ServerHost        string `envconfig:"SERVER_HOST"`
	ServerPort        int    `envconfig:"SERVER_PORT"`
	DBHost            string `envconfig:"DB_HOST"`
	DBPort            int    `envconfig:"DB_PORT"`
	DBName            string `envconfig:"DB_NAME"`
	DBUser            string `envconfig:"DB_USER"`
	DBPassword        string `envconfig:"DB_PASSWORD"`
	JWTSecret         string `envconfig:"JWT_SECRET"`
	RTLifetimeHours   int    `envconfig:"REFRESH_TOKEN_LIFETIME_HOURS"`
	ATLifetimeMinutes int    `envconfig:"ACCESS_TOKEN_LIFETIME_MINUTES"`
}

// DBUrl spits out a ready-to use DB connection string so that you don't have to format it manually.
func (c *Config) DBUrl() string {
	return fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s sslmode=disable",
		c.DBHost, c.DBPort, c.DBName, c.DBUser, c.DBPassword)
}

// ServerUrl spits out a ready-to use server URL in format host:port.
func (c *Config) ServerUrl() string {
	return fmt.Sprintf("%s:%d", c.ServerHost, c.ServerPort)
}

func (c *Config) AccessTokenLifetime() time.Duration {
	return time.Minute * time.Duration(c.ATLifetimeMinutes)
}
