package app

import "github.com/gin-gonic/gin"

var (
	router = gin.Default()
)

func StartApplication() {
	mapUrls()
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
