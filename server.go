package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"main/controller"
	"main/database"
	"main/middlewares"
	"main/service"

	"github.com/gin-gonic/gin"
	gindump "github.com/tpkeeper/gin-dump"
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func setupLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {
	setupLogOutput()

	server := gin.Default()

	db := database.DBConn()

	fmt.Println("Db connect", db.Ping())

	server.Use(gin.Recovery(), middlewares.Logger(),
		middlewares.BasicAuth(), gindump.Dump())

	server.GET("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"code":    "00",
			"message": "Success",
			"data":    videoController.FindAll(),
		})
	})

	server.POST("/videos", func(ctx *gin.Context) {
		err := videoController.Save(ctx)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"code":    "00",
				"message": "Success",
			})
		}
	})

	server.Run(":8080")
}
