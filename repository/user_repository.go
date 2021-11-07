package repository

import (
	"database/sql"
	"github.com/lib/pq"
	"github.com/paul-ilves/wanaku-api-go/utils"
)

type User struct {
	ID             uint64         `db:"id"`
	FirstName      string         `db:"first_name"`
	LastName       string         `db:"last_name"`
	LastAddressID  sql.NullInt32  `db:"last_address"`
	Email          string         `db:"email"`
	PhoneNumber    string         `db:"phone_number"`
	PasswordHash   string         `db:"password_hash"`
	Role           sql.NullString `db:"role"`
	Status         sql.NullString `db:"status"`
	CreatedAt      pq.NullTime    `db:"created_at"`
	BirthDate      pq.NullTime    `db:"birthdate"`
	OrganisationID sql.NullInt32  `db:"organisation_id"`
}

func SelectAllUsers() (*[]User, *utils.AppError) {
	OpenDBConnection()
	query := `select * from "user"`
	var users []User
	err := client.Select(&users, query)
	if err != nil {
		appErr := logError(err)
		return nil, &appErr
	}
	defer CloseDBConnection()
	return &users, nil
}

func SelectUserByID(userID uint64) (*User, *utils.AppError) {
	OpenDBConnection()
	query := `select * from "user" where "id" = $1`
	var u User
	err := client.Get(&u, query, userID)
	if err != nil {
		appErr := logError(err)
		return nil, &appErr
	}
	defer CloseDBConnection()
	return &u, nil
}

func SelectUserByEmail(userEmail string) (*User, *utils.AppError) {
	OpenDBConnection()
	query := `select * from "user" where "email" = $1`
	var u User
	err := client.Get(&u, query, userEmail)
	if err != nil {
		appErr := logError(err)
		return nil, &appErr
	}
	defer CloseDBConnection()
	return &u, nil
}

func InsertUser(u User) (*User, *utils.AppError) {
	OpenDBConnection()

	query := `INSERT INTO "user" (first_name, last_name, email, phone_number, password_hash, created_at) values ($1, $2, $3, $4, $5, $6) returning id`
	var lastInsertId uint64
	err := client.QueryRow(query, u.FirstName, u.LastName, u.Email, u.PhoneNumber, u.PasswordHash, u.CreatedAt).Scan(&lastInsertId)
	if err != nil {
		appErr := logError(err)
		return nil, &appErr
	}
	savedUser, appError := SelectUserByID(lastInsertId)
	if appError != nil {
		return nil, appError
	}

	defer CloseDBConnection()
	return savedUser, nil
}

func logError(err error) utils.AppError {
	utils.Error("Got some problem in db: " + err.Error())
	var errCode uint = 500
	if err == sql.ErrNoRows {
		errCode = 404
	}
	return utils.AppError{
		Message: err.Error(),
		Code:    errCode,
	}
}
