package controllers

import (
	"github.com/buglinjo/golang-rest-api/app/models"
	"github.com/buglinjo/golang-rest-api/app/responses"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct{}

func (uc *UserController) List(c *gin.Context) {
	user := &models.User{}

	responses.Success(c, http.StatusOK, user.All())
}

func (uc *UserController) Profile(c *gin.Context) {
	user, _ := c.Get("user")

	responses.Success(c, http.StatusOK, user)
}
