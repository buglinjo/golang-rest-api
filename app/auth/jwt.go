package auth

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func CreateToken(userId string) (string, error) {
	claims := jwt.MapClaims{}
	// TODO: add more claims...
	claims["sub"] = userId
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(os.Getenv("API_KEY")))
}

func ParseToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if token.Method.Alg() != jwt.SigningMethodHS256.Name {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("API_KEY")), nil
	})
}

func ExtractToken(c *gin.Context) string {
	bearerToken := c.GetHeader("Authorization")
	token := strings.Replace(bearerToken, "Bearer ", "", 1)

	return token
}

func ExtractUserId(token *jwt.Token) (int, error) {
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("failed to extract claims")
	}

	userIdStr, ok := claims["sub"].(string)
	if !ok {
		return 0, errors.New("failed to extract user ID")
	}

	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		return 0, errors.New("failed to convert user ID to int")
	}

	return userId, nil
}
