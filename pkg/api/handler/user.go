package handler

import "clean/pkg/usecase/interfaces"

type Userhandler struct {
	Userservice interfaces.Userusecase
}

func NewUserHandler(user_service interfaces.Userusecase) *Userhandler {
	return &Userhandler{Userservice: user_service}
}
