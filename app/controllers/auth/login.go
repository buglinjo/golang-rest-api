package auth

import (
	"errors"
	"github.com/badoux/checkmail"
	"github.com/buglinjo/golang-rest-api/app/auth"
	"github.com/buglinjo/golang-rest-api/app/models"
	"github.com/buglinjo/golang-rest-api/app/responses"
	"github.com/gin-gonic/gin"
	"strconv"
)

type login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(c *gin.Context) {
	var l login
	_ = c.ShouldBind(&l)

	err := l.validate()
	if err != nil {
		responses.Error(c, 426, err)
		return
	}

	u := &models.User{}

	u, err = u.FindByEmail(l.Email)
	if err != nil {
		responses.Error(c, 401, errors.New("email or password is incorrect"))
		return
	}

	err = u.VerifyPassword(l.Password)
	if err != nil {
		responses.Error(c, 401, errors.New("email or password is incorrect"))
		return
	}

	token, err := auth.CreateToken(strconv.FormatUint(uint64(u.Id), 10))
	if err != nil {
		responses.Error(c, 401, errors.New("could not create token"))
		return
	}

	responses.Success(c, 200, token)
}

func (l *login) validate() error {
	if l.Password == "" {
		return errors.New("password is required")
	}
	if l.Email == "" {
		return errors.New("email is required")
	}
	if err := checkmail.ValidateFormat(l.Email); err != nil {
		return errors.New("email is invalid")
	}

	return nil
}
