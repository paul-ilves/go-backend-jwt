package handlers

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/paul-ilves/wanaku-api-go/config"
	"github.com/paul-ilves/wanaku-api-go/services"
	"github.com/paul-ilves/wanaku-api-go/utils"
	"time"
)

// HandleLogin retrieves Email and Password and passes them for check inside services.CheckEmailAndPassword. If the check is OK - it returns the accessToken+refreshToken pair.
func HandleLogin(c *fiber.Ctx) error {
	var r struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&r); err != nil {
		c.ClearCookie("access_token", "refresh_token")
		return sendError(err.Error(), fiber.StatusBadRequest, c)
	}

	userAgent := string(c.Request().Header.UserAgent())
	response, appError := services.CheckEmailAndPassword(r.Email, r.Password, userAgent)
	if appError != nil {
		c.ClearCookie("access_token", "refresh_token")
		return sendError(appError.Error(), fiber.StatusForbidden, c)
	}

	accessCookie, refreshCookie := generateCookies(response["accessToken"], response["refreshToken"])
	c.Cookie(accessCookie)
	c.Cookie(refreshCookie)

	return c.Status(200).JSON(response)
}

// HandleRegister retrieves data to register a new user, handles the flow to services.RegisterNewUser and in case of happy path - returns the newly created user DTO and accessToken+refreshToken pair.
func HandleRegister(c *fiber.Ctx) error {
	var r struct {
		Email       string `json:"email"`
		Password    string `json:"password"`
		FirstName   string `json:"firstName"`
		LastName    string `json:"lastName"`
		PhoneNumber string `json:"phoneNumber"`
	}
	if err := c.BodyParser(&r); err != nil {
		return sendError(err.Error(), fiber.StatusBadRequest, c)
	}

	userAgent := string(c.Request().Header.UserAgent())
	response, appError := services.RegisterNewUser(services.UserDto{
		FirstName:   r.FirstName,
		LastName:    r.LastName,
		Email:       r.Email,
		PhoneNumber: r.PhoneNumber,
	}, r.Password, userAgent)
	if appError != nil {
		return sendError(appError.Error(), fiber.StatusBadRequest, c)
	}

	accessCookie, refreshCookie := generateCookies(response["accessToken"].(string), response["refreshToken"].(string))
	c.Cookie(accessCookie)
	c.Cookie(refreshCookie)
	return c.Status(200).JSON(response)
}

// HandleLogout receives a refreshToken, invalidates it in the DB and clears th Client's token cookies.
func HandleLogout(c *fiber.Ctx) error {
	refreshToken := c.Get("refreshToken")
	if refreshToken == "" {
		return sendError("Header refreshToken was not sent", 400, c)
	}
	appError := services.InvalidateToken(refreshToken)
	if appError != nil {
		return sendError(appError.Error(), 500, c)
	}

	return c.SendStatus(200)
}

// RefreshToken checks AccessToken (if it's valid and expired), checks RefreshToken (if present in DB, related to user and valid) and if all conditions met - sends a new token pair
func RefreshToken(c *fiber.Ctx) error {
	// TODO implement RefreshToken
	//1. Check Access Token. It should be valid AND expired
	//2. Check Refresh Token. It should be present in DB, related to the user_id inside access token claims AND non-expired
	//3. If conditions met => invalidate the old accessToken and issue a new token pair

	var r struct {
		AccessToken  string `json:"accessToken"`
		RefreshToken string `json:"refreshToken"`
	}
	if err := c.BodyParser(&r); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error":   true,
			"message": "cannot parse request body",
		})
	}

	if r.AccessToken == "" || r.RefreshToken == "" {
		return c.Status(400).JSON(fiber.Map{
			"error":   true,
			"message": "refreshToken or accessToken is missing",
		})
	}

	type CustomClaims struct {
		jwt.RegisteredClaims
		Role string `json:"role"`
	}

	// check accessToken validity
	token, err := jwt.ParseWithClaims(r.AccessToken, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(config.C.JWTSecret), nil
	})
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "true", "message": "JWT Error:" + err.Error()})
	}

	// validate the essential claims
	if !token.Valid {
		return c.Status(401).JSON(fiber.Map{"error": "true", "message": "Token is invalid"})
	}

	claims, ok := token.Claims.(*CustomClaims)
	if !ok {
		return sendError("Bad type assertion :(", 400, c)
	}

	//check if the access token is expired
	if claims.ExpiresAt.After(time.Now()) {
		return sendError("your access token hasn't expired yet!", 400, c)
	}

	return c.Status(200).JSON(fiber.Map{"message": "ok!"})
}

func sendError(msg string, code int, c *fiber.Ctx) error {
	utils.Error("error in handler: " + msg)
	return c.Status(code).JSON(fiber.Map{"error": "true", "message": msg})
}

func generateCookies(accessToken, refreshToken string) (*fiber.Cookie, *fiber.Cookie) {
	accessCookie := &fiber.Cookie{
		Name:     "access_token",
		Value:    accessToken,
		Expires:  time.Now().Add(config.C.AccessTokenLifetime()),
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
