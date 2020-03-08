package service

import (
	"github.com/yuta17/hyperlp/repository"
)

type subscribeService struct {
	repo repository.SubscribeRepository
}

func (s *subscribeService) IsDuplicated(email string) (bool, error) {
	subscribe, err := s.repo.FindByEmail(email)
	if err != nil {
		return false, err
	}
	if subscribe != nil {
		return true, nil
	}
	return false, nil
}
