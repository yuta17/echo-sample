package main

import (
	"net/http"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

type Subscribe struct {
	gorm.Model
	ID    int    `gorm:"primary_key`
	Email string `json:email`
}

func createSubscribe(c echo.Context) error {
	db, err := gorm.Open("mysql",
		os.Getenv("DB_USERNAME")+":"+os.Getenv("DB_PASSWORD")+"@tcp("+os.Getenv("DB_HOST")+":"+os.Getenv("DB_PORT")+")/"+os.Getenv("DB_DATABASE")+"?charset=utf8mb4&parseTime=True&loc=Local",
	)
	defer db.Close()
	if err != nil {
		log.Panic(err)
	}

	subscribe := new(Subscribe)
	if err := c.Bind(subscribe); err != nil {
		return err
	}

	db.Create(&subscribe)
	return c.String(http.StatusOK, "OK")
}

func InitMigrate() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := gorm.Open("mysql",
		os.Getenv("DB_USERNAME")+":"+os.Getenv("DB_PASSWORD")+"@tcp("+os.Getenv("DB_HOST")+":"+os.Getenv("DB_PORT")+")/"+os.Getenv("DB_DATABASE")+"?charset=utf8mb4&parseTime=True&loc=Local",
	)

	defer db.Close()
	db.AutoMigrate(&Subscribe{})

	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	InitMigrate()
	e := echo.New()

	// Middleware
	// e.Use(middleware.Logger())
	// e.Use(middleware.Recover())

	// Routes
	e.POST("/subscribes", createSubscribe)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
