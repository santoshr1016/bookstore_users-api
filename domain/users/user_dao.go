package users

import (
	"fmt"
	"github.com/santoshr1016/bookstore_users-api/utils/errors"
)

//func Save(user User) *errors.RestError{
//	return nil
//}
//
//func Get(userId int64) (*User, *errors.RestError) {
//	return nil, nil
//}

var (
	userDB = make(map[int64] *User)

)

func (user *User) Get() *errors.RestError {
	result := userDB[user.Id]
	if result == nil {
		return errors.NewNotFoundError(fmt.Sprint("User %d not found", user.Id))
	}
	user.Id = result.Id
	user.Email = result.Email
	user.FirstName = result.FirstName
	user.Lastname = result.Lastname
	user.DateCreated = result.DateCreated
	return nil
}

func (user *User) Save() *errors.RestError{
	current := userDB[user.Id]
	if current != nil {
		if current.Email == user.Email {
			return errors.NewBadRequestError(fmt.Sprint("Email %s already registered", user.Email))
		}
		return errors.NewBadRequestError(fmt.Sprint("User %d already exists", user.Id))
	}
	userDB[user.Id] = user
	return  nil
}
