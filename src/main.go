package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/mojtabaRKS/link_shortener/internal/webserver"
	"github.com/mojtabaRKS/link_shortener/pkg/db"
	"github.com/mojtabaRKS/link_shortener/pkg/env"
)

func init() {
	env.Load()
	db.Load()
}

func main() {
	routesHandler := func(r *gin.Engine) *gin.Engine {
		return webserver.Routes(r)
	}

	webserver.RunServer(routesHandler, os.Getenv("PORT"))
}
