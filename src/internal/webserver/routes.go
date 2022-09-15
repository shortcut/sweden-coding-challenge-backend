package webserver

import "github.com/gin-gonic/gin"

func Routes(r *gin.Engine) *gin.Engine {
	rg := r.RouterGroup.Group("/api/v1")
	{
		rg.POST("register", auth_v1.Register)
	}

	return r
}
