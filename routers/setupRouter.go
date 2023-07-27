package routers

import (
	"github.com/gin-gonic/gin"
	"online.eric/ginWebServer/middlewares"
	"online.eric/ginWebServer/routers/user"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", GetPing)

	authorized := r.Group("/", middlewares.BasicAuth())
	{
		authorized.GET("/user/:name", user.GetUser)
		authorized.POST("setUser", user.PostSetUser)
	}

	return r
}
