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

	var user users.User
	if err := context.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		context.JSON(restErr.Status, restErr)
		return
	}

	result, saveErr := services.UsersService.CreateUser(user)
	if saveErr != nil {
		context.JSON(saveErr.Status, saveErr)
		return
	}

	context.JSON(http.StatusCreated, result.Marshall(context.GetHeader("X-Public") == "true"))
}

func Get(ctx *gin.Context) {
	userId, userErr := getUserId(ctx.Param("user_id"))
	if userErr != nil {
		ctx.JSON(userErr.Status, userErr.Message)
		return
	}

	user, getErr := services.UsersService.GetUser(userId)
	if getErr != nil {
		// TODO Handle User Error Creation error
		//ctx.JSON(getErr.Status, getErr)
		ctx.JSON(getErr.Status, getErr)
		return
	}

	//context.String(http.StatusNotImplemented, "Implement me")
	//ctx.JSON(http.StatusOK, user)
	ctx.JSON(http.StatusOK, user.Marshall(ctx.GetHeader("X-Public") == "true"))
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

	updatedUser, err := services.UsersService.UpdateUser(isPartial, user)
	if err != nil {
		ctx.JSON(err.Status, err.Message)
	}
	ctx.JSON(http.StatusOK, updatedUser.Marshall(ctx.GetHeader("X-Public") == "true"))
}

func Delete(ctx *gin.Context) {
	userId, userErr := getUserId(ctx.Param("user_id"))
	if userErr != nil {
		ctx.JSON(userErr.Status, userErr.Message)
		return
	}
	if err := services.UsersService.DeleteUser(userId); err != nil {
		ctx.JSON(err.Status, err.Message)
		return
	}
	ctx.JSON(http.StatusOK, map[string]string{"status": "deleted"})
}

func Search(ctx *gin.Context) {
	status := ctx.Query("status")
	usersGot, err := services.UsersService.SearchUser(status)
	if err != nil {
		ctx.JSON(err.Status, err.Message)
		return
	}
	ctx.JSON(http.StatusOK, usersGot.Marshall(ctx.GetHeader("X-Public") == "true"))
}

func Login(ctx *gin.Context) {
	var request users.LoginRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		ctx.JSON(restErr.Status, restErr)
		return
	}
	user, err := services.UsersService.LoginUser(request)
	if err != nil {
		ctx.JSON(err.Status, err.Message)
	}
	//ctx.JSON(http.StatusOK, user)
	ctx.JSON(http.StatusOK, user.Marshall(ctx.GetHeader("X-Public") == "true"))
}
