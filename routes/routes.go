package routes

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/paul-ilves/wanaku-api-go/config"
	"github.com/paul-ilves/wanaku-api-go/handlers"
)

func SetupRoutes(a *fiber.App) {
	a.Get("/", handleIndex)

	//create routes group
	router := a.Group("/api")
	router.Post("/login", handlers.HandleLogin)
	router.Post("/logout", handlers.HandleLogout)
	router.Post("/register", handlers.HandleRegister)
	router.Post("/refresh", handlers.RefreshToken)

	router.Use(jwtware.New(jwtware.Config{
		SigningMethod: "HS256",
		SigningKey:    []byte(config.C.JWTSecret),
	}))

	router.Get("/users", handlers.HandleGetAllUsers)
	router.Get("/users/me", handlers.HandleGetCurrentUser)
	router.Get("/users/:userID", handlers.HandleGetUser)

	router.Use(handlePageNotFound)
	a.Use(handlePageNotFound)
}

func handleIndex(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"allSystems": "Go!",
	})
}

func handlePageNotFound(c *fiber.Ctx) error {
	return c.Status(404).JSON(fiber.Map{
		"error":   true,
		"message": "path not found",
	})
}
