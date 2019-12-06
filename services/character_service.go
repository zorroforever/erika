package services

import (
	"iris/commons"
	"iris/repositories"
	"iris/web/viewmodels"
	"strconv"
)

type CharacterService interface {
	// 按角色id获取角色信息
	GetCharacterPropertyDataByChId(chId int) []viewmodels.CharacterDispModel
	// 按用户ID获取角色ID 角色名字等信息
	GetCharacterIdByUserId(userId int) []viewmodels.CharacterDispModel
}

func NewCharacterService() CharacterService {
	return &characterService{
		repo: repositories.NewCharacterPropertyDBRepo(),
	}
}

type characterService struct {
	repo repositories.CharacterPropertyRepo
}

func (c *characterService) GetCharacterIdByUserId(userId int) []viewmodels.CharacterDispModel {
	var kv []viewmodels.CharacterDispModel
	chId, b := c.repo.GetCharacterIdByUserId(userId)
	if b {
		for _, item := range chId {
			ucd := c.repo.GetUserCharacterDataByChId(item)
			kv = append(kv, viewmodels.CharacterDispModel{
				Key:   strconv.Itoa(item),
				Name:  ucd.ChName,
				Value: strconv.Itoa(item),
			})
		}
	}
	return kv
}
func (c *characterService) GetCharacterPropertyDataByChId(chId int) []viewmodels.CharacterDispModel {
	uc := c.repo.GetUserCharacterDataByChId(chId)
	chp := c.repo.GetCharacterPropertyDataByChId(chId)
	var kv []viewmodels.CharacterDispModel
	kv = append(kv, viewmodels.CharacterDispModel{
		Key:   "currentExperience",
		Name:  "当前经验",
		Value: strconv.Itoa(uc.CurrentExperience),
	})
	kv = append(kv, viewmodels.CharacterDispModel{
		Key:   "currentStatus",
		Name:  "当前状态",
		Value: strconv.Itoa(uc.CurrentStatus),
	})
	kv = append(kv, viewmodels.CharacterDispModel{
		Key:   "honor",
		Name:  "荣耀",
		Value: strconv.Itoa(uc.Honor),
	})
	kv = append(kv, viewmodels.CharacterDispModel{
		Key:   "hp",
		Name:  "HP",
		Value: strconv.Itoa(chp.HelPoint),
	})
	kv = append(kv, viewmodels.CharacterDispModel{
		Key:   "manaPoint",
		Name:  "MP",
		Value: strconv.Itoa(chp.ManaPoint),
	})
	kv = append(kv, viewmodels.CharacterDispModel{
		Key:   "Physical",
		Name:  "体力",
		Value: strconv.Itoa(chp.Physical),
	})
	kv = append(kv, viewmodels.CharacterDispModel{
		Key:   "phyAttack",
		Name:  "物理攻击力",
		Value: commons.FloatToString(chp.PhyAttack),
	})

	kv = append(kv, viewmodels.CharacterDispModel{
		Key:   "manaAttack",
		Name:  "魔法攻击力",
		Value: commons.FloatToString(chp.ManaAttack),
	})

	kv = append(kv, viewmodels.CharacterDispModel{
		Key:   "phyDefense",
		Name:  "物理防御力",
		Value: commons.FloatToString(chp.PhyDefense),
	})

	kv = append(kv, viewmodels.CharacterDispModel{
		Key:   "phyDefense",
		Name:  "物理防御力",
		Value: commons.FloatToString(chp.PhyDefense),
	})

	kv = append(kv, viewmodels.CharacterDispModel{
		Key:   "str",
		Name:  "力量",
		Value: strconv.Itoa(chp.Str),
	})
	kv = append(kv, viewmodels.CharacterDispModel{
		Key:   "dex",
		Name:  "灵巧",
		Value: strconv.Itoa(chp.Dex),
	})
	kv = append(kv, viewmodels.CharacterDispModel{
		Key:   "agl",
		Name:  "敏捷",
		Value: strconv.Itoa(chp.Agl),
	})
	kv = append(kv, viewmodels.CharacterDispModel{
		Key:   "luk",
		Name:  "幸运",
		Value: strconv.Itoa(chp.Luk),
	})
	kv = append(kv, viewmodels.CharacterDispModel{
		Key:   "vit",
		Name:  "耐力",
		Value: strconv.Itoa(chp.Vit),
	})
	kv = append(kv, viewmodels.CharacterDispModel{
		Key:   "int",
		Name:  "智力",
		Value: strconv.Itoa(chp.Int),
	})

	kv = append(kv, viewmodels.CharacterDispModel{
		Key:   "evade",
		Name:  "回避力",
		Value: strconv.Itoa(chp.Evade),
	})
	return kv
}
