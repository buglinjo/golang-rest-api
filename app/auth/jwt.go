package auth

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func CreateToken(userId uint32) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = userId
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(os.Getenv("API_KEY")))
}

func TokenValid(c *gin.Context) {
	tokenString := ExtractToken(c)
	fmt.Println(tokenString)
}

func ExtractToken(c *gin.Context) string {
	keys := c.GetHeader("Authorization")
	return keys
}