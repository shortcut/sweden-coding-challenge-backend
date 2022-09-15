package webserver

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func RunServer(routesHandler func(r *gin.Engine) *gin.Engine, portNo string) {
	//setLogFiles
	gin.SetMode(os.Getenv("APP_ENV"))
	apiRouter := gin.Default()
	setMiddlewares(apiRouter)
	routesHandler(apiRouter)
	logrus.Info("Starting HTTP server")

	err := apiRouter.Run(":" + portNo)
	if err != nil {
		panic(err)
	}
}

func setMiddlewares(r *gin.Engine) {
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
}
