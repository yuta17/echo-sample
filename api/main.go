package main

import (
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"github.com/yuta17/hyperlp/database"
	"github.com/yuta17/hyperlp/model"
)

func createSubscribe(c echo.Context) error {
	db, err := gorm.Open("mysql",
		os.Getenv("DB_USERNAME")+":"+os.Getenv("DB_PASSWORD")+"@tcp("+os.Getenv("DB_HOST")+":"+os.Getenv("DB_PORT")+")/"+os.Getenv("DB_DATABASE")+"?charset=utf8mb4&parseTime=True&loc=Local",
	)
	defer db.Close()
	if err != nil {
		log.Panic(err)
	}

	subscribe := new(model.Subscribe)
	if err := c.Bind(subscribe); err != nil {
		return err
	}

	db.Create(&subscribe)
	return c.String(http.StatusOK, "OK")
}

func main() {
	database.InitMigrate()
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
