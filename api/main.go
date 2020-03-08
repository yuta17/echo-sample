package main

import (
	"github.com/labstack/echo/v4"
	"github.com/yuta17/hyperlp/controller"
	"github.com/yuta17/hyperlp/database"
)

func main() {
	db := database.DB()
	database.InitMigrate(db)

	e := echo.New()

	// Middleware
	// e.Use(middleware.Logger())
	// e.Use(middleware.Recover())

	// Routes
	e.POST("/subscribes", controller.CreateSubscribe)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

// curl -X POST http://localhost:1323/subscribes -H 'Content-Type: application/json' -d '{"email":"joe@labstack"}'
