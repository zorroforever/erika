package repositories

import (
	"github.com/jinzhu/gorm"
	"iris/datamodels"
	"iris/datasource"
)

type MapRepository interface {
	InsPersonMoveStatus(in datamodels.BizChMoveLib) bool
}

func NewMapDBRepo() MapRepository {
	return &mapSQLRepository{source: datasource.DB}
}

type mapSQLRepository struct {
	source *gorm.DB
}

func (r *mapSQLRepository) InsPersonMoveStatus(in datamodels.BizChMoveLib) (b bool) {

	qc := r.source.Table("BIZ_CH_MOVE_LIB")
	if err := qc.Create(&in); err == nil {
		return true
	} else {
		return false
	}
}
