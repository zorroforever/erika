package services

import (
	"iris/datamodels"
	"iris/repositories"
)

type CharacterService interface {
	// 按角色id获取角色信息
	GetCharacterPropertyDataByChId(chId int) []datamodels.BizChProperty
}

func NewCharacterService() CharacterService {
	return &characterService{
		repo: repositories.NewCharacterPropertyDBRepo(),
	}
}

type characterService struct {
	repo repositories.CharacterPropertyRepo
}

func (c *characterService) GetCharacterPropertyDataByChId(chId int) []datamodels.BizChProperty {
	return c.repo.GetCharacterPropertyDataByChId(chId)
}
