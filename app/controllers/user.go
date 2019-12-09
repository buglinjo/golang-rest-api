package controllers

import (
	"fmt"

	"github.com/buglinjo/golang-rest-api/app/models"
	"github.com/buglinjo/golang-rest-api/app/responses"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func List(c *gin.Context) {
	db, ok := c.MustGet("db").(*gorm.DB)
	if !ok {
		fmt.Println("Can't get DB instance")
	}

	user := &models.User{}

	responses.Success(c, 200, user.All(db))
}
