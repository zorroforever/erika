package services

import (
	"iris/repositories"
)

type MapService interface {
}

func NewMapService() MapService {
	return &mapService{
		repo: repositories.NewMapDBRepo(),
	}
}

type mapService struct {
	repo repositories.MapRepository
}
