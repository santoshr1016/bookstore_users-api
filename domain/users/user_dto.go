package users

import (
	"github.com/santoshr1016/bookstore_users-api/utils/errors"
	"strings"
)

type User struct {
	Id int64 `json:"id"`
	FirstName string `json:"first_name"`
	Lastname string `json:"last_name"`
	Email string `json:"email"`
	DateCreated string `json:"date_created"`
}
/* TODO Why this function is moved to method over the user
"Users package" is responsible for the validity of the user and not any other module
Users will be validating itself by the struct that is passed.
*/
/*
func Validate(user *User) *errors.RestError{
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return errors.NewBadRequestError("invalid email address")
	}
	return nil
}
*/

func (user *User) Validate() *errors.RestError{
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return errors.NewBadRequestError("invalid email address")
	}
	return nil
}