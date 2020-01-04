package services

import (
	"github.com/santoshr1016/bookstore_users-api/domain/users"
	"github.com/santoshr1016/bookstore_users-api/utils/crypto_utils"
	"github.com/santoshr1016/bookstore_users-api/utils/date_utils"
	"github.com/santoshr1016/bookstore_users-api/utils/errors"
)

// This kind of structuring has flaws
// You cannot mock the inside services functions, so cannot test
// Its a package function(which are highly not recommended)
var (
	UsersService usersServiceInterface = &usersService{}
)

type usersService struct {
}

type usersServiceInterface interface {
	CreateUser(users.User) (*users.User, *errors.RestError)
	GetUser(int64) (*users.User, *errors.RestError)
	UpdateUser(bool, users.User) (*users.User, *errors.RestError)
	DeleteUser(int64) *errors.RestError
	SearchUser(string) (users.ArrayOfUsers, *errors.RestError)
}

func (s *usersService) CreateUser(user users.User) (*users.User, *errors.RestError) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	user.DateCreated = date_utils.GetNowDbFormat()
	user.Status = users.StatusActive
	user.Password = crypto_utils.GetMd5(user.Password)
	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *usersService) GetUser(userId int64) (*users.User, *errors.RestError) {
	user := &users.User{Id: userId}
	if err := user.Get(); err != nil {
		return nil, err
	}
	return user, nil
}

func (s *usersService) UpdateUser(isPatch bool, user users.User) (*users.User, *errors.RestError) {
	//currentUser := &users.User{Id: user.Id}
	//if err := currentUser.Get(); err != nil {
	//	return nil, err
	//}
	currentUser, err := UsersService.GetUser(user.Id)
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

func (s *usersService) DeleteUser(userId int64) *errors.RestError {
	user := &users.User{Id: userId}
	return user.Delete()
}
func (s *usersService) SearchUser(status string) (users.ArrayOfUsers, *errors.RestError) {
	dao := users.User{}
	return dao.FindByStatus(status)
}
