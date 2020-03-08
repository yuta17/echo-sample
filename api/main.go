package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/yuta17/hyperlp/controller"
	"github.com/yuta17/hyperlp/database"
	"github.com/yuta17/hyperlp/repository"
	"github.com/yuta17/hyperlp/service"
)

func main() {
	db := database.DB()
	defer db.Close()
	database.InitMigrate(db)

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	subscribeRepository := repository.NewSubscribeRepository(db)
	subscribeService := service.NewSubscribeService(subscribeRepository)
	subscribeController := controller.NewSubscribeController(subscribeService, db)
	e.POST("/subscribes", subscribeController.CreateSubscribe)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

// curl -X POST http://localhost:1323/subscribes -H 'Content-Type: application/json' -d '{"email":"joe@labstack"}'
