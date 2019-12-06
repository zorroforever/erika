package repositories

import (
	"github.com/jinzhu/gorm"
	"iris/datamodels"
	"iris/datasource"
)

/**
数据库 各种对外接口
*/
type CharacterPropertyRepo interface {
	// 按角色id获取角色信息
	GetCharacterPropertyDataByChId(chId int) datamodels.BizChProperty
	// 按角色ID获取用户角色信息
	GetUserCharacterDataByChId(chId int) datamodels.BizUserCharacter
	// 按用户ID获取角色ID
	GetCharacterIdByUserId(userId int) ([]int, bool)
}

type CharacterPropertySQLRepository struct {
	source *gorm.DB
}

func NewCharacterPropertyDBRepo() CharacterPropertyRepo {
	return &CharacterPropertySQLRepository{source: datasource.DB}
}

func (r *CharacterPropertySQLRepository) GetCharacterPropertyDataByChId(chId int) (chp datamodels.BizChProperty) {
	qc := r.source.Table("BIZ_CH_PROPERTY").Model(&datamodels.BizChProperty{})
	qc.Where("CH_ID = ?", chId).Find(&chp).Limit(1)
	return chp
}

func (r *CharacterPropertySQLRepository) GetCharacterIdByUserId(userId int) (chId []int, ex bool) {
	qc := r.source.Table("BIZ_USER_CHARACTER").Model(&datamodels.BizUserCharacter{})
	var res []datamodels.BizUserCharacter
	qc.Where("USER_ID = ?", userId).Find(&res)
	ex = false
	for _, tmp := range res {
		a := datamodels.BizUserCharacter(tmp)
		chId = append(chId, a.ChId)
		ex = true
	}
	//if res != nil {
	//	ex = true
	//} else {
	//	ex = false
	//}
	return chId, ex
}
func (r *CharacterPropertySQLRepository) GetUserCharacterDataByChId(chId int) (chp datamodels.BizUserCharacter) {
	qc := r.source.Table("BIZ_USER_CHARACTER").Model(&datamodels.BizUserCharacter{})
	qc.Where("CH_ID = ?", chId).Find(&chp).Limit(1)
	return chp
}
