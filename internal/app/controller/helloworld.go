package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/huugiii/hello-world-api/internal/app/service"
)

type HelloWorldController struct {
	service *service.HelloWorldService
}

const (
	defaultMessage = "Hello, World !"
)

func (c *HelloWorldController) HelloWorld(ctx *gin.Context) {
	message := ctx.Request.URL.Query().Get("message")
	if message == "" {
		message = defaultMessage
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": message,
	})
}

func NewHelloWorldController(service *service.HelloWorldService) *HelloWorldController {
	return &HelloWorldController{
		service: service,
	}
}
