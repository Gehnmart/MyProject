package user

import "errors"

func ValidateUser(user User) error {
	if len(user.Password) < 8 || len(user.Password) > 72 {
		return errors.New("password must be between 8 and 72 characters")
	}

	if len(user.Name) < 3 || len(user.Name) > 32 {
		return errors.New("name must be between 3 and 72 characters")
	}

	if len(user.Email) < 3 || len(user.Email) > 64 {
		return errors.New("email must be between 3 and 72 characters")
	}

	return nil
}
