package types

import (
	"fmt"
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

var (
	bcryptSale         = 10
	minFirstNameLength = 2
	minLastNameLength  = 2
	minPasswordLength  = 6
)

type UpdateUserParams struct {
	FirstName string `bson:"firstName,omitempty" json:"firstName"`
	LastName  string `bson:"lastName,omitempty" json:"lastName"`
}

type InsertUserParams struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func (params InsertUserParams) Validate() map[string]string {

	errors := map[string]string{}

	if len(params.FirstName) < minFirstNameLength {
		errors["firstName"] = fmt.Sprintf("Firstname must be at least %d characters", minFirstNameLength)
	}
	if len(params.LastName) < minLastNameLength {
		errors["lastName"] = fmt.Sprintf("Last name must be at least %d characters", minLastNameLength)
	}
	if len(params.Password) < minPasswordLength {
		errors["password"] = fmt.Sprintf("Password must be at least %d characters", minLastNameLength)
	}
	if !isValidEmail(params.Email) {
		errors["email"] = "Email is invalid"
	}

	return errors

}

func (params UpdateUserParams) Validate() map[string]string {
	errors := map[string]string{}

	if params.FirstName == "" && params.LastName == "" {
		errors["empty"] = "At least one of first name or last name must be present "
	} else {
		if params.FirstName != "" && len(params.FirstName) < minFirstNameLength {
			errors["firstName"] = fmt.Sprintf("Firstname must be at least %d characters", minFirstNameLength)
		}
		if params.LastName != "" && len(params.LastName) < minLastNameLength {
			errors["lastName"] = fmt.Sprintf("Last name must be at least %d characters", minLastNameLength)
		}
	}

	return errors
}

func isValidEmail(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(email)
}

type User struct {
	ID                string `bson:"_id,omitempty" json:"id,omitempty"`
	FirstName         string `bson:"firstName" json:"firstName"`
	LastName          string `bson:"lastName" json:"lastName"`
	Email             string `bson:"email" json:"email"`
	EncryptedPassword string `bson:"encryptedPassword" json:"-"`
}

func NewUserFromParams(params InsertUserParams) (*User, error) {

	encrypted, err := EncryptPassword(params.Password)

	if err != nil {
		return nil, err
	}

	return &User{
		FirstName:         params.FirstName,
		LastName:          params.LastName,
		Email:             params.Email,
		EncryptedPassword: string(encrypted),
	}, nil
}

func EncryptPassword(p string) (string, error) {
	encrypted, err := bcrypt.GenerateFromPassword([]byte(p), bcryptSale)

	return string(encrypted), err
}
