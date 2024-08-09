package middleware

import (
	"github.com/gin-gonic/gin"
	"golang-ddd-demo/application"
	"log"
	"os"
)

type ServerImpl struct {
	router *gin.Engine
}

func InitServer() Server {
	serverImpl := &ServerImpl{}
	router := gin.Default()
	router.Use(CorsMiddleware())
	application.InitCatController(router)
	serverImpl.router = router
	return serverImpl
}

func (api *ServerImpl) RunServer() {
	appPort := os.Getenv("PORT")
	if appPort == "" {
		appPort = os.Getenv("LOCAL_PORT") //localhost
	}
	log.Fatal(api.router.Run(":" + appPort))
}
