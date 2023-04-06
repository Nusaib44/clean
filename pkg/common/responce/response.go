package responce

import (
	"strings"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Errors  interface{} `json:"errors,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func SuccessResponse(g *gin.Context, message string, data ...interface{}) {
	res := Response{
		Status:  true,
		Message: message,
		Errors:  nil,
		Data:    data,
	}
	g.JSON(200, res)
}

func ErrorResponse(g *gin.Context, message string, err string, data interface{}, code int) {
	splittedError := strings.Split(err, "\n")
	res := Response{
		Status:  false,
		Message: message,
		Errors:  splittedError,
		Data:    data,
	}
	g.JSON(code, res)
}
