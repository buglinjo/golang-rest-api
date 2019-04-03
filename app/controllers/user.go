package controllers

import (
	"fmt"

	"github.com/buglinjo/golang-rest-api/app/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func List(c *gin.Context) {
	db, ok := c.MustGet("db").(*gorm.DB)
	if !ok {
		fmt.Println("Can't get DB instance")
	}

	user := &models.User{}

	c.JSON(200, gin.H{
		"user": db.Find(&user).Value,
	})
}
