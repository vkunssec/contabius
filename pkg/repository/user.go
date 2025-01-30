package repository

import (
	"github.com/vkunssec/contabius/pkg/domain"
)

func GetUserById(id string) (domain.User, error) {
	return domain.User{}, nil
}

func GetUserByEmail(email string) (domain.User, error) {
	return domain.User{}, nil
}

func LoginUser(user domain.UserLogin) (domain.User, error) {
	return domain.User{}, nil
}

func CreateUser(user domain.UserRequest) (domain.User, error) {

	return domain.User{}, nil
}

func UpdateUser(id string, user domain.UserUpdateRequest) (domain.User, error) {
	return domain.User{}, nil
}

func UpdateUserPassword(id string, user domain.UserUpdatePassword) (domain.User, error) {
	return domain.User{}, nil
}
