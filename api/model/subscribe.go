package model

import (
	"github.com/jinzhu/gorm"
)

type Subscribe struct {
	gorm.Model
	Email string `gorm:"unique;not null"`
}
