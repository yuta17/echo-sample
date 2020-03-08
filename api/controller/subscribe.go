package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
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

	db.Create(&subscribe)
	return c.String(http.StatusOK, "OK")
}
