package services

import (
	"github.com/santoshr1016/bookstore_users-api/domain/users"
	"github.com/santoshr1016/bookstore_users-api/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestError){
	/*
	if err := users.Validate(&user); err != nil {
		return nil, err
	}
	*/
	// TODO See how the above function call moved to method call user.Validate()
	if err := user.Validate(); err != nil {
		return nil, err
	}
	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil
}

func GetUser(userId int64) (*users.User, *errors.RestError){
	result := &users.User{Id:userId}
	if err := result.Get(); err != nil {
		return nil, nil
	}
	return result, nil
}
