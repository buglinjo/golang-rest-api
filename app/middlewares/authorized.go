package middlewares

import (
	"errors"
	"github.com/buglinjo/golang-rest-api/app/auth"
	"github.com/buglinjo/golang-rest-api/app/models"
	"github.com/buglinjo/golang-rest-api/app/responses"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Authorized() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := auth.ExtractToken(c)
		token, err := auth.ParseToken(tokenString)
		if err != nil || !token.Valid {
			responses.Error(c, http.StatusUnauthorized, errors.New("token is invalid"))
			return
		}

		userId, err := auth.ExtractUserId(token)
		if err != nil {
			responses.Error(c, http.StatusUnauthorized, err)
			return
		}
		userModel := models.User{}
		user, err := userModel.FindById(userId)
		if err != nil {
			responses.Error(c, http.StatusUnauthorized, err)
			return
		}

		c.Set("user", user)
		c.Next()
	}
}
