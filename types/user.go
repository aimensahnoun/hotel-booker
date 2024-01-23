package types

import (
	"fmt"
	"regexp"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

var (
	bcryptSale         = 10
	minFirstNameLength = 2
	minLastNameLength  = 2
	minPasswordLength  = 6
)

type InsertUserParams struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func (params InsertUserParams) Validate() []string {

	errors := []string{}

	if len(params.FirstName) < minFirstNameLength {
		errors = append(errors, fmt.Sprintf("Firstname must be at least %d characters", minFirstNameLength))
	}
	if len(params.LastName) < minLastNameLength {
		errors = append(errors, fmt.Sprintf("Last name must be at least %d characters", minLastNameLength))
	}
	if len(params.Password) < minPasswordLength {
		errors = append(errors, fmt.Sprintf("Password must be at least %d characters", minLastNameLength))
	}
	if !isValidEmail(params.Email) {
		errors = append(errors, fmt.Sprintf("Email is invalid"))
	}

	return errors

}

func isValidEmail(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(email)
}

type User struct {
	ID                primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	FirstName         string             `bson:"firstName" json:"firstName"`
	LastName          string             `bson:"lastName" json:"lastName"`
	Email             string             `bson:"email" json:"email"`
	EncryptedPassword string             `bson:"encryptedPassword" json:"-"`
}

func NewUserFromParams(params InsertUserParams) (*User, error) {
	encrypted, err := bcrypt.GenerateFromPassword([]byte(params.Password), 10)

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
