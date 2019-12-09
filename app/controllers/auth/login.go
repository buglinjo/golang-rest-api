package auth

import (
	"fmt"

	"github.com/buglinjo/golang-rest-api/app/models"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var u models.User
	_ = c.ShouldBind(&u)
	fmt.Println(u)
}
