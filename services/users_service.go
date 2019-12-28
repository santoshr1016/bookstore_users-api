package services

import (
	"github.com/santoshr1016/bookstore_users-api/domain/users"
	"github.com/santoshr1016/bookstore_users-api/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestError) {
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

func GetUser(userId int64) (*users.User, *errors.RestError) {
	user := &users.User{Id: userId}
	if err := user.Get(); err != nil {
		return nil, err
	}
	return user, nil
}

func UpdateUser(isPatch bool, user users.User) (*users.User, *errors.RestError) {
	//currentUser := &users.User{Id: user.Id}
	//if err := currentUser.Get(); err != nil {
	//	return nil, err
	//}

	currentUser, err := GetUser(user.Id)
	if err != nil {
		return nil, err
	}
	if isPatch {
		if user.FirstName != "" {
			currentUser.FirstName = user.FirstName
		}
		if user.Lastname != "" {
			currentUser.Lastname = user.Lastname
		}
		if user.Email != "" {
			currentUser.Email = user.Email
		}
	} else {
		currentUser.FirstName = user.FirstName
		currentUser.Lastname = user.Lastname
		currentUser.Email = user.Email
	}

	if err := currentUser.Update(); err != nil {
		return nil, err
	}
	return currentUser, nil

}

func DeleteUser(userId int64) *errors.RestError {
	user := &users.User{Id: userId}
	return user.Delete()
}
