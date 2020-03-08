package service

import (
	"github.com/yuta17/hyperlp/repository"
)

type SubscribeService struct {
	Repo repository.SubscribeRepository
}

func NewSubscribeService(r repository.SubscribeRepository) *SubscribeService {
	return &SubscribeService{
		Repo: r,
	}
}

func (s *SubscribeService) IsDuplicated(email string) (bool, error) {
	subscribe, err := s.Repo.FindByEmail(email)
	if err != nil {
		return false, err
	}
	if subscribe != nil {
		return true, nil
	}
	return false, nil
}
