package services

import (
	"iris/datamodels"
	"iris/repositories"
)

type MapService interface {
	DoUpdPersonMoveStatus(in datamodels.BizChMoveLib) (res bool)
}

func NewMapService() MapService {
	return &mapService{
		repo: repositories.NewMapDBRepo(),
	}
}

type mapService struct {
	repo repositories.MapRepository
}

func (s *mapService) DoUpdPersonMoveStatus(in datamodels.BizChMoveLib) (b bool) {
	b = s.repo.InsPersonMoveStatus(in)
	s.repo.UpdPersonPosition(in)
	return b
}
