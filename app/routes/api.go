package routes

import (
	"github.com/buglinjo/golang-rest-api/app/controllers"
	"github.com/buglinjo/golang-rest-api/app/controllers/auth"
	"github.com/buglinjo/golang-rest-api/app/middlewares"
	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	r := gin.Default()

	api := r.Group("api")
	api.Use(middlewares.Api())
	{
		r.Use(middlewares.Api())

		api.POST("auth/login", auth.Login)

		authGroup := api.Group("/")
		authGroup.Use(middlewares.Authorized())
		{
			uc := controllers.UserController{}
			authGroup.GET("users/list", uc.List)
			authGroup.GET("/", uc.Profile)
		}
	}

	return r
}
