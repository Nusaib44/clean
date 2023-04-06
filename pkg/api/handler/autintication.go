package handler

import (
	"clean/pkg/common/responce"
	"clean/pkg/domain"
	services "clean/pkg/usecase/interfaces"
	"net/http"

	"github.com/gin-gonic/gin"
)

// type Authinticationhandler struct {
// 	authUsecase interfaces.AuthenticationUsecase
// }

// func Newauthintication(aut interfaces.AuthenticationUsecase) *Authinticationhandler {
// 	return &Authinticationhandler{authUsecase: aut}
// }

type AuthHandler struct {
	authUseCase services.AuthenticationUsecase
}

func NewAuthHandler(usecase services.AuthenticationUsecase) *AuthHandler {
	return &AuthHandler{
		authUseCase: usecase,
	}

}

func (ah *AuthHandler) Signin(g *gin.Context) {
	var body domain.Users
	if err := g.BindJSON(&body); err != nil {
		responce.ErrorResponse(g, "failed to bind", err.Error(), body, http.StatusBadRequest)
		return
	}

	user, err := ah.authUseCase.Register(body)
	reply := "welcome" + user.First_Name
	if err != nil {
		responce.ErrorResponse(g, "error to registor", err.Error(), user, 500)
	}
	responce.SuccessResponse(g, reply, user)

}
