package application

import (
	"github.com/gin-gonic/gin"
	"github.com/huugiii/hello-world-api/internal/app/controller"
	"github.com/huugiii/hello-world-api/internal/app/service"
)

type Application struct {
	engine *gin.Engine
}

func (a *Application) setupRoutes() {
	helloworldController := controller.NewHelloWorldController(service.NewHelloWorldService())
	
	a.engine.GET("/", helloworldController.HelloWorld)
}

func New() *Application {
	gin.SetMode(gin.ReleaseMode)
	return &Application{
		engine: gin.Default(),
	}
}

// Start the application
func (a *Application) Start() {
	a.setupRoutes()
	a.engine.Run(":8080")
}