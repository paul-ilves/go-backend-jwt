package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/paul-ilves/wanaku-api-go/services"
	"time"
)

func HandleLogin(c *fiber.Ctx) error {
	var r struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&r); err != nil {
		c.ClearCookie("access_token", "refresh_token")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	userAgent := string(c.Request().Header.UserAgent())
	response, appError := services.CheckEmailAndPassword(r.Email, r.Password, userAgent)
	if appError != nil {
		c.ClearCookie("access_token", "refresh_token")
		return c.Status(403).JSON(fiber.Map{
			"error":   true,
			"message": appError.Error(),
		})
	}

	accessCookie, refreshCookie := generateCookies(response["accessToken"], response["refreshToken"])
	c.Cookie(accessCookie)
	c.Cookie(refreshCookie)

	return c.Status(200).JSON(response)
}

func HandleRegister(c *fiber.Ctx) error {
	var r struct {
		Email       string `json:"email"`
		Password    string `json:"password"`
		FirstName   string `json:"firstName"`
		LastName    string `json:"lastName"`
		PhoneNumber string `json:"phoneNumber"`
	}
	if err := c.BodyParser(&r); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	userAgent := string(c.Request().Header.UserAgent())
	response, appError := services.RegisterNewUser(services.UserDto{
		FirstName:   r.FirstName,
		LastName:    r.LastName,
		Email:       r.Email,
		PhoneNumber: r.PhoneNumber,
	}, r.Password, userAgent)
	if appError != nil {
		return c.Status(400).JSON(fiber.Map{
			"error":   true,
			"message": appError.Error(),
		})
	}

	accessCookie, refreshCookie := generateCookies(response["accessToken"].(string), response["refreshToken"].(string))
	c.Cookie(accessCookie)
	c.Cookie(refreshCookie)
	return c.Status(200).JSON(response)
}

func HandleLogout(c *fiber.Ctx) error {
	refreshToken := c.Get("refreshToken")
	if refreshToken == "" {
		return c.Status(400).JSON(fiber.Map{
			"error":   true,
			"message": "Header refreshToken was not sent",
		})
	}
	appError := services.InvalidateToken(refreshToken)
	if appError != nil {
		return c.Status(500).JSON(fiber.Map{
			"error":   true,
			"message": appError.Error(),
		})
	}

	return c.SendStatus(200)
}

func RefreshToken(c *fiber.Ctx) error {
	refreshToken := c.Get("refreshToken")
	if refreshToken == "" {
		return c.Status(400).JSON(fiber.Map{
			"error":   true,
			"message": "Header refreshToken was not sent",
		})
	}
	return nil
}

func generateCookies(accessToken, refreshToken string) (*fiber.Cookie, *fiber.Cookie) {
	accessCookie := &fiber.Cookie{
		Name:     "access_token",
		Value:    accessToken,
		Expires:  time.Now().Add(24 * time.Hour),
		HTTPOnly: true,
		Secure:   true,
	}

	refreshCookie := &fiber.Cookie{
		Name:     "refresh_token",
		Value:    refreshToken,
		Expires:  time.Now().Add(10 * 24 * time.Hour),
		HTTPOnly: true,
		Secure:   true,
	}

	return accessCookie, refreshCookie
}

//todo create endpoint to gain new access token
