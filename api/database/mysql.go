package database

import (
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	"github.com/yuta17/hyperlp/model"
)

func InitMigrate() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	hoge := os.Getenv("DB_USERNAME") + ":" + os.Getenv("DB_PASSWORD") + "@tcp(" + os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + ")/" + os.Getenv("DB_DATABASE") + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", hoge)

	defer db.Close()
	db.AutoMigrate(&model.Subscribe{})

	if err != nil {
		log.Fatal(err)
	}
}
