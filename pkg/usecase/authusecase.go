package usecase

import (
	"clean/pkg/domain"
	interfaces "clean/pkg/repository/interfaces"
	services "clean/pkg/usecase/interfaces"
	"errors"
	"fmt"
)

type authUseCase struct {
	authRepo interfaces.AuthRepository
}

func NewAuthUseCase(repo interfaces.AuthRepository) services.AuthenticationUsecase {
	return &authUseCase{authRepo: repo}
}

func (c *authUseCase) Register(user domain.Users) (domain.Users, error) {
	retun, _ := c.authRepo.FindUser(user.Email)
	//fmt.Println(user.Email, retun.Email)
	fmt.Println(retun.Email, user.Email)
	if retun.Email == user.Email {
		return user, errors.New("user already exists ")

	}

	_, err := c.authRepo.Register(user)
	fmt.Println(err)
	if err != nil {
		return user, errors.New("email id already exists")
	}
	fmt.Println(user, "register")

	//fmt.Printf("\n\n %v ")
	return user, err
}

// ---------------------------------find user-----------------
func (c *authUseCase) FindUser(email string) (domain.UserResponse, error) {
	user, err := c.authRepo.FindUser(email)
	return user, err
}
