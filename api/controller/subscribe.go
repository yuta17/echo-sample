package controller

import (
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"github.com/yuta17/hyperlp/model"
	"github.com/yuta17/hyperlp/service"
)

type SubscribeController struct {
	service *service.SubscribeService
	db      *gorm.DB
}

func NewSubscribeController(service *service.SubscribeService, db *gorm.DB) *SubscribeController {
	return &SubscribeController{
		service: service,
		db:      db,
	}
}

func (s *SubscribeController) CreateSubscribe(c echo.Context) error {
	db := s.db
	subscribe := new(model.Subscribe)
	if err := c.Bind(subscribe); err != nil {
		return err
	}

	IsDuplicatedEmail, err := s.service.IsDuplicated(subscribe.Email)
	if IsDuplicatedEmail == true {
		return c.JSON(422, "this email has been registered.")
	}

	if err = db.Create(&subscribe).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusCreated, subscribe)
}
