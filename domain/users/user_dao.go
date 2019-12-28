package users

import (
	"fmt"
	"github.com/santoshr1016/bookstore_users-api/utils/mysql_utils"

	//"github.com/go-sql-driver/mysql"
	"github.com/santoshr1016/bookstore_users-api/datasources/mysql/users_db"
	"github.com/santoshr1016/bookstore_users-api/utils/date_utils"
	"github.com/santoshr1016/bookstore_users-api/utils/errors"
)

//func Save(user User) *errors.RestError{
//	return nil
//}
//
//func Get(userId int64) (*User, *errors.RestError) {
//	return nil, nil
//}
const (
	indexUniqueEmail = "email_UNIQUE"
	noRowsError      = "no rows in result set"
	queryInsertUser  = "INSERT INTO users(first_name, last_name, email, date_created) VALUES(?,?,?,?);"
	queryGetUser     = "SELECT id, first_name, last_name, email, date_created FROM users where id=?;"
)

// Local In memoryDB
//var (
//	userDB = make(map[int64] *User)
//)

func (user *User) Get() *errors.RestError {
	if err := users_db.Client.Ping(); err != nil {
		fmt.Println("Mysql connection error", err)
		panic(err)
	}
	stmt, err := users_db.Client.Prepare(queryGetUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	getResult := stmt.QueryRow(user.Id)

	if err := getResult.Scan(&user.Id, &user.FirstName, &user.Lastname, &user.Email, &user.DateCreated); err != nil {
		fmt.Println("Error while scan")
		return mysql_utils.ParseError(err)
		//if strings.Contains(err.Error(), noRowsError) {
		//	return errors.NewBadRequestError(
		//		fmt.Sprintf("user id %d does not exists", user.Id))
		//}
		//return errors.NewInternalServerError( fmt.Sprintf(" %s error while getting the %d user id", err.Error(), user.Id))
	}

	// TODO Remove this in-memory to Mysql
	/*
		result := userDB[user.Id]
		if result == nil {
			return errors.NewNotFoundError(fmt.Sprint("User %d not found", user.Id))
		}
		user.Id = result.Id
		user.Email = result.Email
		user.FirstName = result.FirstName
		user.Lastname = result.Lastname
		user.DateCreated = result.DateCreated
	*/
	return nil
}

func (user *User) Save() *errors.RestError {
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()
	user.DateCreated = date_utils.GetNowString()
	insertResult, saveErr := stmt.Exec(user.FirstName, user.Lastname, user.Email, user.DateCreated)
	if saveErr != nil {
		return mysql_utils.ParseError(saveErr)
	}
	//sqlErr, ok := saveErr.(*mysql.MySQLError)
	//if !ok {
	//	return errors.NewInternalServerError(fmt.Sprintf("error when trying to save user: %s", saveErr.Error()))
	//}
	//fmt.Println("SQL Error Number: %d", sqlErr.Number)
	//fmt.Println("SQL Error Message: %d", sqlErr.Message)
	//switch sqlErr.Number {
	//case 1062:
	//	return errors.NewBadRequestError( fmt.Sprintf("email %s already exists", user.Email))
	//}

	// http://go-database-sql.org/errors.html
	// TODO Better way to Handle SQL Error, Check Above, The indexUniqueEmail is removed
	//if saveErr != nil {
	//	if strings.Contains(saveErr.Error(), indexUniqueEmail){
	//		return errors.NewBadRequestError(
	//			fmt.Sprintf("email %s already exists", user.Email))
	//	}
	//	return errors.NewInternalServerError(fmt.Sprintf("error when trying to save user: %s", saveErr.Error()))
	// This is same as above, but less efficient
	//result, err := users_db.Client.Exec(queryInsertUser, user.FirstName, user.Lastname, user.Email, user.DateCreated)

	userId, err := insertResult.LastInsertId()
	if err != nil {
		return mysql_utils.ParseError(err)
		//return errors.NewInternalServerError(fmt.Sprintf("error when trying to get lastId: %s", err.Error()))
	}
	user.Id = userId
	//user.DateCreated = date_utils.GetNowString()

	/* In memory DB
	current := userDB[user.Id]
	if current != nil {
		if current.Email == user.Email {
			return errors.NewBadRequestError(fmt.Sprint("Email %s already registered", user.Email))
		}
		return errors.NewBadRequestError(fmt.Sprint("User %d already exists", user.Id))
	}*/
	//now := time.Now().UTC()
	////now := time.Now()
	////user.DateCreated = now.Format("Mon Jan 2 2006 15:04:05 MST")
	//user.DateCreated = now.Format("Mon Jan 2 2006 15:04:05")
	//userDB[user.Id] = user
	//user.DateCreated = date_utils.GetNowString()
	return nil
}
