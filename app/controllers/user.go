package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func List(c *gin.Context) {
	db, ok := c.MustGet("db").(*gorm.DB)
	if !ok {
		fmt.Println("Can't get DB instance")
	}

	type Result struct {
		Id   int
		Name string
	}

	result := Result{}

	db.Raw("SELECT * FROM users").Scan(&result)

	c.JSON(200, gin.H{
		"result": result,
	})
}
