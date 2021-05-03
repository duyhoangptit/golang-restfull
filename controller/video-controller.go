package controller

import (
	"main/entities"
	"main/service"

	"github.com/gin-gonic/gin"
)

type VideoController interface {
	FindAll() []entities.Video
	Save(ctx *gin.Context) entities.Video
}

type controller struct {
	service service.VideoService
}

func New(service service.VideoService) VideoController {
	return &controller {
		service: service,
	}
}

func (c *controller) FindAll() []entities.Video {
	return c.service.FindAll()
}

func (c *controller) Save(ctx *gin.Context) entities.Video {
	var video entities.Video
	ctx.BindJSON(&video)
	c.service.Save(video)
	return video
}
