package interfaces

import "clean/pkg/domain"

type AuthenticationUsecase interface {
	Register(use domain.Users) (domain.Users, error)
}
