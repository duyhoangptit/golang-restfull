package main

import (
	"main/controller"
	"main/service"

	"github.com/gin-gonic/gin"
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func main() {
	server := gin.Default()

	server.GET("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"code":    "00",
			"message": "Success",
			"data":    videoController.FindAll(),
		})
	})

	server.POST("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"code":    "00",
			"message": "Success",
			"data":    videoController.Save(ctx),
		})
	})

	server.Run(":8080")
}
