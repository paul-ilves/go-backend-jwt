package services

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/paul-ilves/wanaku-api-go/config"
	"github.com/paul-ilves/wanaku-api-go/repository"
	"github.com/paul-ilves/wanaku-api-go/utils"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"strconv"
	"time"
)

func DecodeToken(tokenString string) (*UserDto, *utils.AppError) {
	claims := jwt.MapClaims{}
	prefix := tokenString[:7]
	if prefix != "Bearer " {
		return nil, &utils.AppError{
			Message: "Invalid Auth Token header",
			Code:    400,
		}
	}

	tokenString = tokenString[7:]
	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.C.JWTSecret), nil
	})
	if err != nil {
		errMessage := fmt.Errorf("we got an error: %s", err.Error())
		fmt.Println(errMessage)
		return nil, &utils.AppError{
			Message: errMessage.Error(),
			Code:    400,
		}
	}

	subValue := claims["sub"]
	fmt.Println("sub is ", subValue)

	//TODO get the user data from the DB

	userId, err := strconv.ParseUint(subValue.(string), 10, 32)

	fmt.Println("userId is: ", userId)

	user, appErr := repository.SelectUserByID(userId)
	if appErr != nil {
		return nil, appErr
	}
	userDto := toDTO(*user)
	return &userDto, nil
}

func CheckEmailAndPassword(email, password, userAgent string) (map[string]string, *utils.AppError) {
	u, err := repository.SelectUserByEmail(email)
	if err != nil {
		return nil, err
	}

	//verify password
	passErr := bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password))
	if passErr != nil {
		return nil, &utils.AppError{
			Message: passErr.Error(),
			Code:    403,
		}
	}

	//create JWT
	accessToken, refreshToken, appError := generateTokenPair(u.ID, userAgent)
	if appError != nil {
		return nil, appError
	}

	return map[string]string{"accessToken": accessToken, "refreshToken": refreshToken}, nil
}

func RegisterNewUser(u UserDto, password, userAgent string) (map[string]interface{}, *utils.AppError) {
	//validate user fields

	//if valid => create new user in DB
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return nil, &utils.AppError{
			Message: err.Error(),
			Code:    500,
		}
	}
	user := u.toEntity()
	user.PasswordHash = string(hashed)
	savedUser, appError := repository.InsertUser(user)
	if appError != nil {
		return nil, appError
	}

	savedUserDto := toDTO(*savedUser)

	//generate JWT
	accessToken, refreshToken, appError := generateTokenPair(savedUserDto.ID, userAgent)
	if appError != nil {
		return nil, appError
	}

	//create and return response map
	return map[string]interface{}{"user": savedUserDto, "accessToken": accessToken, "refreshToken": refreshToken}, nil
}

func generateTokenPair(userId uint64, userAgent string) (string, string, *utils.AppError) {
	accessToken, appError := generateAccessToken(userId)
	if appError != nil {
		return "", "", appError
	}
	refreshToken, appError := generateRefreshToken(userId, userAgent)
	if appError != nil {
		return "", "", appError
	}

	return accessToken, refreshToken, nil
}

func generateRefreshToken(userID uint64, userAgent string) (string, *utils.AppError) {
	expiresAt := time.Now().Add(time.Hour * time.Duration(config.C.RTLifetimeHours))
	sign := randomString(64)
	rt, appError := repository.InsertRefreshToken(userID, sign, expiresAt, userAgent)
	if appError != nil {
		return "", &utils.AppError{
			Message: appError.Error(),
			Code:    500,
		}
	}

	return rt.Sign, nil
}

func InvalidateToken(refreshToken string) *utils.AppError {
	appError := repository.DeleteToken(refreshToken)
	if appError != nil {
		return appError
	}

	return nil
}

func generateAccessToken(subId uint64) (string, *utils.AppError) {
	type CustomClaims struct {
		jwt.RegisteredClaims
		Role string `json:"role"`
	}

	user, err := repository.SelectUserByID(subId)
	if err != nil {
		return "", err
	}

	role := user.Role.String

	myClaims := CustomClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   strconv.FormatUint(subId, 10),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(config.C.AccessTokenLifetime())),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		Role: role,
	}

	unsignedJwt := jwt.NewWithClaims(jwt.SigningMethodHS256, myClaims)
	signedString, passErr := unsignedJwt.SignedString([]byte(config.C.JWTSecret))

	if passErr != nil {
		return "", &utils.AppError{
			Message: passErr.Error(),
			Code:    500,
		}
	}
	return signedString, nil
}

func randomString(length int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, length)
	rand.Read(b)
	return fmt.Sprintf("%x", b)[:length]
}
