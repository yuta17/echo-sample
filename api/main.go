package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/yuta17/hyperlp/database"
	"github.com/yuta17/hyperlp/model"
)

func createSubscribe(c echo.Context) error {
	db := database.DB()
	defer db.Close()

	subscribe := new(model.Subscribe)
	if err := c.Bind(subscribe); err != nil {
		return err
	}

	db.Create(&subscribe)
	return c.String(http.StatusOK, "OK")
}

func main() {
	db := database.DB()
	database.InitMigrate(db)

	e := echo.New()

	// Middleware
	// e.Use(middleware.Logger())
	// e.Use(middleware.Recover())

	// Routes
	e.POST("/subscribes", createSubscribe)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

// curl -X POST http://localhost:1323/subscribes -H 'Content-Type: application/json' -d '{"email":"joe@labstack"}'
