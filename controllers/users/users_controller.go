package users

import (
	"github.com/gin-gonic/gin"
	"github.com/santoshr1016/bookstore_users-api/domain/users"
	"github.com/santoshr1016/bookstore_users-api/services"
	"github.com/santoshr1016/bookstore_users-api/utils/errors"
	"net/http"
	"strconv"
)

func GetUser(context *gin.Context){
	userId, userErr := strconv.ParseInt(context.Param("user_id"), 10, 64)
	if userErr != nil {
		err := errors.NewBadRequestError("User id is not a number")
		context.JSON(err.Status, err)
		return
	}

	user, getErr := services.GetUser(userId)
	if getErr != nil {
		// TODO Handle User Error Creation error
		context.JSON(getErr.Status, getErr)
		return
	}

	//context.String(http.StatusNotImplemented, "Implement me")
	context.JSON(http.StatusOK, user)
}

func CreateUser(context *gin.Context){

	// TODO This is one way of reading the http body
	/*
	var user users.User
	bytes, err := ioutil.ReadAll(context.Request.Body)
	if err != nil {
		// Handle the error
		return
	}
	if err := json.Unmarshal(bytes, &user); err != nil {
		// Handle json error
		fmt.Println("error", err)
		return
	}
	fmt.Println(user)
	*/
	// TODO This is Shorter version of the above
	var user users.User
	if err := context.ShouldBindJSON(&user); err != nil {
		// TODO Handle json error, Bad Request
		restErr := errors.NewBadRequestError("invalid json body")
		context.JSON(restErr.Status, restErr)
		return
	}

	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		// TODO Handle User Error Creation error
		context.JSON(saveErr.Status, saveErr)
		return
	}

	//context.String(http.StatusNotImplemented, "Implement me")
	context.JSON(http.StatusCreated, result)
}

func SearchUser(context *gin.Context){
	context.String(http.StatusNotImplemented, "Implement me")
}
