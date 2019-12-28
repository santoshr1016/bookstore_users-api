package mysql_utils

import (
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/santoshr1016/bookstore_users-api/utils/errors"
	"strings"
)

const (
	noRowsError = "no rows in result set"
)

func ParseError(err error) *errors.RestError {
	fmt.Println("Inside ParseError")
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), noRowsError) {
			return errors.NewInternalServerError("No rows found")
		}
		return errors.NewInternalServerError("Error parsing the rows in DB")
	}
	fmt.Println("SQL Error Number: %d", sqlErr.Number)
	fmt.Println("SQL Error Message: %d", sqlErr.Message)
	switch sqlErr.Number {
	case 1062:
		return errors.NewBadRequestError("Duplicate data")
	}
	return errors.NewInternalServerError("Error processing the request")
}
