package routes

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/paul-ilves/wanaku-api-go/controllers"
	"os"
)

func SetupRoutes(a *fiber.App) {
	a.Get("/", handleIndex)

	//create routes group
	router := a.Group("/api")
	router.Post("/login", controllers.HandleLogin)
	router.Post("/logout", controllers.HandleLogout)
	router.Post("/register", controllers.HandleRegister)

	router.Use(jwtware.New(jwtware.Config{
		SigningMethod: "HS256",
		SigningKey:    []byte(os.Getenv("JWT_SECRET")),
	}))

	router.Get("/users", controllers.HandleGetAllUsers)
	//router.Get("/users/me", controllers.HandleLoggedUser) // todo implement
	router.Get("/users/:userID", controllers.HandleGetUser)

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
