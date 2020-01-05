package app

import (
	"github.com/gin-gonic/gin"
	"github.com/santoshr1016/bookstore_users-api/logger"
)

var (
	router = gin.Default()
)

func StartApplication() {
	mapUrls()
	logger.Info("about to start the application.....")
	/*
		router.GET("/ping", func(context *gin.Context) {
			context.JSON(200, gin.H{
				"message": "pong",
			})
		})
		TODO : Move this to controller
	*/

	router.Run(":8989")
}
