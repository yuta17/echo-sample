package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/yuta17/hyperlp/model"
)

type SubscribeRepository interface {
	FindByEmail(email string) (*model.Subscribe, error)
}

type SubscribeRepositoryImpl struct {
	DB *gorm.DB
}

func NewSubscribeRepository(db *gorm.DB) SubscribeRepository {
	return &SubscribeRepositoryImpl{db}
}

func (s *SubscribeRepositoryImpl) FindByEmail(email string) (*model.Subscribe, error) {
	subscribe := model.Subscribe{}
	result := s.DB.First(&subscribe, "email = ?", email)
	if result.RecordNotFound() {
		return nil, nil
	}
	if result.Error != nil {
		return nil, result.Error
	}
	return &subscribe, nil
}
