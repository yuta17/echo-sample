package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"github.com/yuta17/hyperlp/database"
	"github.com/yuta17/hyperlp/model"
)

func CreateSubscribe(c echo.Context) error {
	db := database.DB()
	defer db.Close()

	subscribe := new(model.Subscribe)
	if err := c.Bind(subscribe); err != nil {
		return err
	}

	if err := db.Create(&subscribe).Error; err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, subscribe)
}
