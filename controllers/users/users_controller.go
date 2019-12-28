package users

import (
	"github.com/gin-gonic/gin"
	"github.com/santoshr1016/bookstore_users-api/domain/users"
	"github.com/santoshr1016/bookstore_users-api/services"
	"github.com/santoshr1016/bookstore_users-api/utils/errors"
	"net/http"
	"strconv"
)

func getUserId(userIdParam string) (int64, *errors.RestError) {
	userId, userErr := strconv.ParseInt(userIdParam, 10, 64)
	if userErr != nil {
		return 0, errors.NewBadRequestError("user id should be number")
	}
	return userId, nil
}

func Create(context *gin.Context) {

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

func Get(ctx *gin.Context) {
	userId, userErr := getUserId(ctx.Param("user_id"))
	if userErr != nil {
		ctx.JSON(userErr.Status, userErr.Message)
		return
	}

	user, getErr := services.GetUser(userId)
	if getErr != nil {
		// TODO Handle User Error Creation error
		ctx.JSON(getErr.Status, getErr)
		return
	}

	//context.String(http.StatusNotImplemented, "Implement me")
	ctx.JSON(http.StatusOK, user)
}

func Update(ctx *gin.Context) {
	userId, userErr := getUserId(ctx.Param("user_id"))
	if userErr != nil {
		ctx.JSON(userErr.Status, userErr.Message)
		return
	}

	var user users.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		// TODO Handle json error, Bad Request
		restErr := errors.NewBadRequestError("invalid json body")
		ctx.JSON(restErr.Status, restErr)
		return
	}
	user.Id = userId
	isPartial := ctx.Request.Method == http.MethodPatch

	updatedUser, err := services.UpdateUser(isPartial, user)
	if err != nil {
		ctx.JSON(err.Status, err.Message)
	}
	ctx.JSON(http.StatusOK, updatedUser)
}

func Delete(ctx *gin.Context) {
	userId, userErr := getUserId(ctx.Param("user_id"))
	if userErr != nil {
		ctx.JSON(userErr.Status, userErr.Message)
		return
	}
	if err := services.DeleteUser(userId); err != nil {
		ctx.JSON(err.Status, err.Message)
		return
	}
	ctx.JSON(http.StatusOK, map[string]string{"status": "deleted"})
}

func SearchUser(context *gin.Context) {
	context.String(http.StatusNotImplemented, "Implement me")
}
