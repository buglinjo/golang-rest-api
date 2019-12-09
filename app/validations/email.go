package validations

import (
	"errors"

	"github.com/badoux/checkmail"
	"github.com/buglinjo/golang-rest-api/app/models"
)

func Email(u *models.User) error {
	if u.Email == "" {
		return errors.New("email is required")
	}
	if err := checkmail.ValidateFormat(u.Email); err != nil {
		return errors.New("email is invalid")
	}

	return nil
}
