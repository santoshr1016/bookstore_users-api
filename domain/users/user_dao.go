package users

import (
	"fmt"
	"github.com/santoshr1016/bookstore_users-api/utils/mysql_utils"

	"github.com/santoshr1016/bookstore_users-api/datasources/mysql/users_db"
	"github.com/santoshr1016/bookstore_users-api/utils/date_utils"
	"github.com/santoshr1016/bookstore_users-api/utils/errors"
)

const (
	queryInsertUser       = "INSERT INTO users(first_name, last_name, email, date_created, status, password) VALUES(?,?,?,?,?,?);"
	queryGetUser          = "SELECT id, first_name, last_name, email, date_created, status, password FROM users where id=?;"
	queryUpdateUser       = "Update users SET first_name=?, last_name=?, email=?, status=?, password=? where id=?;"
	queryDeleteUser       = "Delete from users where id=?;"
	queryFindUserByStatus = "Select id, first_name, last_name, email, date_created, status from users where status=?;"
)

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

	if err := getResult.Scan(
		&user.Id, &user.FirstName, &user.Lastname, &user.Email,
		&user.DateCreated, &user.Status, &user.Password); err != nil {
		fmt.Println("Error while scan")
		return mysql_utils.ParseError(err)
	}

	return nil
}

func (user *User) Save() *errors.RestError {
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()
	insertResult, saveErr := stmt.Exec(user.FirstName, user.Lastname, user.Email,
		user.DateCreated, user.Status, user.Password)
	if saveErr != nil {
		return mysql_utils.ParseError(saveErr)
	}
	userId, err := insertResult.LastInsertId()
	if err != nil {
		return mysql_utils.ParseError(err)
	}
	user.Id = userId

	return nil
}

func (user *User) Update() *errors.RestError {
	stmt, err := users_db.Client.Prepare(queryUpdateUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()
	user.DateCreated = date_utils.GetNowString()
	_, saveErr := stmt.Exec(user.FirstName, user.Lastname, user.Email,
		user.Id, user.Status, user.Password)
	if saveErr != nil {
		return mysql_utils.ParseError(saveErr)
	}

	return nil

}

func (user *User) Delete() *errors.RestError {
	stmt, err := users_db.Client.Prepare(queryDeleteUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()
	if _, delErr := stmt.Exec(user.Id); delErr != nil {
		return mysql_utils.ParseError(delErr)
	}
	return nil
}

func (user *User) FindByStatus(status string) ([]User, *errors.RestError) {
	stmt, err := users_db.Client.Prepare(queryFindUserByStatus)
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	rows, err := stmt.Query(status)
	if err != nil {
		return nil, mysql_utils.ParseError(err)
	}
	defer rows.Close()

	results := make([]User, 0)

	for rows.Next() {
		var user User
		if err := rows.Scan(&user.Id, &user.FirstName, &user.Lastname,
			&user.Email, &user.DateCreated, &user.Status); err != nil {
			return nil, mysql_utils.ParseError(err)
		}
		results = append(results, user)
	}

	if len(results) == 0 {
		return nil, errors.NewNotFoundError(fmt.Sprint("No users with matching status found %s ", status))
	}
	return results, nil
}
