package controller

import (
	"main/entities"
	"main/service"
	"main/validators"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type VideoController interface {
	FindAll() []entities.Video
	Save(ctx *gin.Context) error
}

type controller struct {
	service service.VideoService
}

var validate *validator.Validate

func New(service service.VideoService) VideoController {
	validate = validator.New()
	validate.RegisterValidation("is-cool", validators.ValidateCoolTitle)

	return &controller{
		service: service,
	}
}

func (c *controller) FindAll() []entities.Video {
	return c.service.FindAll()
}

func (c *controller) Save(ctx *gin.Context) error {
	var video entities.Video
	err := ctx.BindJSON(&video)
	if err != nil {
		return err
	}

	err = validate.Struct(video)

	if err != nil {
		return err
	}

	c.service.Save(video)
	return nil
}
