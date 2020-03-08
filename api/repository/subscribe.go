package repository

import (
	"github.com/yuta17/hyperlp/model"
)

type SubscribeRepository interface {
	FindByEmail(email string) (*model.Subscribe, error)
}
