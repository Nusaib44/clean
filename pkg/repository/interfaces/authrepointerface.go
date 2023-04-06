package interfaces

import "clean/pkg/domain"

type AuthRepository interface {
	Register(user domain.Users) (int, error)
	FindUser(email string) (domain.UserResponse, error)
}
