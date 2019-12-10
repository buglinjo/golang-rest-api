package controllers

import (
	"github.com/buglinjo/golang-rest-api/app/models"
	"github.com/buglinjo/golang-rest-api/app/responses"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func List(c *gin.Context) {
	db, _ := c.MustGet("db").(*gorm.DB)

	user := &models.User{}

	responses.Success(c, 200, user.All(db))
}
