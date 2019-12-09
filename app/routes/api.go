package routes

import (
	"github.com/buglinjo/golang-rest-api/app/controllers"
	"github.com/buglinjo/golang-rest-api/app/controllers/auth"
	"github.com/buglinjo/golang-rest-api/app/middlewares"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func Setup(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	r.Use(middlewares.Api(db))

	r.GET("/", controllers.List)

	r.POST("auth/login", auth.Login)

	return r
}
