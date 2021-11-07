package repository

import (
	"github.com/paul-ilves/wanaku-api-go/utils"
	"time"
)

type RefreshToken struct {
	ID        uint64    `db:"id"`
	Sign      string    `db:"sign"`
	IssuedAt  time.Time `db:"issued_at"`
	ExpiresAt time.Time `db:"expires_at"`
	UserID    uint64    `db:"user_id"`
	DeviceID  string    `db:"device_id"`
}

// 1. a user requests a new access token by providing his non-expired refresh token
// happy path: 2. the server creates and returns a new access token + new refresh token
// sad path:  2. if the refresh token is obsolete -> c.ClearCookie("access_token", "refresh_token") AND return 401

func SelectTokenByUserID(userID uint64) (*RefreshToken, *utils.AppError) {
	OpenDBConnection()
	query := `select * from "refresh_token" where "user_id" = $1`
	var t RefreshToken
	err := client.Get(&t, query, userID)
	if err != nil {
		appErr := logError(err)
		return nil, &appErr
	}
	defer CloseDBConnection()
	return &t, nil
}

func InsertRefreshToken(userID uint64, sign string, expiresAt time.Time, userAgent string) (*RefreshToken, *utils.AppError) {
	OpenDBConnection()

	query := `INSERT INTO "refresh_token" (sign, expires_at, user_id, device_id, issued_at) values ($1, $2, $3, $4, current_timestamp) returning id`
	var tokenID uint64
	err := client.QueryRow(query, sign, expiresAt, userID, userAgent).Scan(&tokenID)
	if err != nil {
		appError := logError(err)
		return nil, &appError
	}
	appError := deleteRefreshTokensWithDiffID(userID, tokenID, userAgent)
	if appError != nil {
		return nil, appError
	}
	query = `select * from "refresh_token" where "id" = $1`
	var t RefreshToken
	err = client.Get(&t, query, tokenID)
	if err != nil {
		appError := logError(err)
		return nil, &appError
	}

	defer CloseDBConnection()
	return &t, nil
}

func DeleteToken(refreshTokenString string) *utils.AppError {
	OpenDBConnection()
	query := `delete from "refresh_token" where "sign" = $1`
	_, err := client.Exec(query, refreshTokenString)
	if err != nil {
		appError := logError(err)
		return &appError
	}
	defer CloseDBConnection()
	return nil
}

func deleteRefreshTokensWithDiffID(userID, tokenID uint64, userAgent string) *utils.AppError {
	query := `delete from "refresh_token" where "user_id" = $1 and "device_id" = $2 and "id" != $3`
	_, err := client.Exec(query, userID, userAgent, tokenID)
	if err != nil {
		appError := logError(err)
		return &appError
	}
	return nil
}
