package services

import (
	"github.com/paul-ilves/wanaku-api-go/repository"
	"github.com/paul-ilves/wanaku-api-go/utils"
)

func GetAllUsers() (*[]UserDto, *utils.AppError) {
	users, err := repository.SelectAllUsers()
	if err != nil {
		return nil, err
	}

	var userDTOs = make([]UserDto, 0)
	for _, user := range *users {
		userDTOs = append(userDTOs, toDTO(user))
	}

	return &userDTOs, nil
}

func GetUser(userID uint64) (*UserDto, *utils.AppError) {
	u, err := repository.SelectUserByID(userID)
	if err != nil {
		return nil, err
	}

	userDTO := toDTO(*u)
	return &userDTO, nil
}

func (u UserDto) toEntity() repository.User {
	return repository.User{
		FirstName:   u.FirstName,
		LastName:    u.LastName,
		Email:       u.Email,
		PhoneNumber: u.PhoneNumber,
	}
}

func toDTO(u repository.User) UserDto {
	return UserDto{
		ID:          u.ID,
		FirstName:   u.FirstName,
		LastName:    u.LastName,
		Email:       u.Email,
		PhoneNumber: u.PhoneNumber,
	}
}
