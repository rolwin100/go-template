package auth

import (
	"errors"

	"github.com/rolwin100/adopler/app/utils"
)

func (a *Auth) SignUp(email string, password string, name string) error {
	// Create a new user.
	repository := a.RepositoryConfig.Repository()
	isUserPresent, err := repository.GetUserByEmail(email)
	if err != nil {
		// Return status 400 and error message.
		return err
	}

	if isUserPresent.Next() {
		// User already exists.
		return errors.New("User already exists.")
	}
	hashedPassword := utils.EncryptPassword(password)
	repository.CreateUser(name, email, hashedPassword)
	return nil
}
