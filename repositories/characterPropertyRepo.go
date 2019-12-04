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
	GetCharacterPropertyDataByChId(chId int) []datamodels.BizChProperty
}

type CharacterPropertySQLRepository struct {
	source *gorm.DB
}

func NewCharacterPropertyDBRepo() CharacterPropertyRepo {
	return &CharacterPropertySQLRepository{source: datasource.DB}
}

func (r *CharacterPropertySQLRepository) GetCharacterPropertyDataByChId(chId int) (chp []datamodels.BizChProperty) {
	qc := r.source.Table("BIZ_CH_PROPERTY").Model(&datamodels.BizChProperty{})
	qc.Where("CH_ID = ?", chId).Find(&chp)
	return chp
}
