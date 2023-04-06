package server

import (
	"clean/pkg/api/handler"

	"github.com/gin-gonic/gin"
)

type ServerHttp struct {
	engine *gin.Engine
}

func Newserver(auth handler.AuthHandler) *ServerHttp {
	engine := gin.New()

	engine.Use(gin.Logger())
	engine.LoadHTMLGlob("templates/*.html")

	userapi := engine.Group("user")
	{
		userapi.POST("/register", auth.Signin)
	}
	return &ServerHttp{engine: engine}
}
func (sh *ServerHttp) Start() {

	sh.engine.Run(":8080")

}
