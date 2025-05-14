package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/huugiii/hello-world-api/internal/app/service"
)

type HelloWorldController struct {
	service *service.HelloWorldService
	prefix  string
}

const (
	defaultMessage = "Hello, World !"
)

func (c *HelloWorldController) HelloWorld(ctx *gin.Context) {
	message := ctx.Request.URL.Query().Get("message")
	if message == "" {
		message = defaultMessage
	}

	message = c.prefix + " - " + message

	ctx.JSON(http.StatusOK, gin.H{
		"message": message,
	})
}

func NewHelloWorldController(service *service.HelloWorldService, prefix string) *HelloWorldController {
	return &HelloWorldController{
		service: service,
		prefix:  prefix,
	}
}
