package repositories

import (
	"github.com/jinzhu/gorm"
	"iris/datamodels"
	"iris/datasource"
	"strconv"
)

type MapRepository interface {
	InsPersonMoveStatus(in datamodels.BizChMoveLib) bool
	UpdPersonPosition(in datamodels.BizChMoveLib)
	UpdPersonStatus(int, int) bool
}

func NewMapDBRepo() MapRepository {
	return &mapSQLRepository{source: datasource.DB}
}

type mapSQLRepository struct {
	source *gorm.DB
}

func (r *mapSQLRepository) InsPersonMoveStatus(in datamodels.BizChMoveLib) (b bool) {
	//in.ArriveTimeStr = time.Unix(in.ArriveTime, 0).Format("2006-01-02 15:04:05")
	qc := r.source.Table("BIZ_CH_MOVE_LIB")
	if err := qc.Create(&in); err == nil {
		return true
	} else {
		return false
	}
}
func (r *mapSQLRepository) UpdPersonPosition(in datamodels.BizChMoveLib) {

	qc := r.source.Table("BIZ_USER_CHARACTER")
	qc.Model(&datamodels.BizUserCharacter{}).Where("CH_ID = ?", in.ChId).Updates(datamodels.BizUserCharacter{MapId: strconv.Itoa(in.TMapId), PointX: in.TX, PointY: in.TY})
}

func (r *mapSQLRepository) UpdPersonStatus(chId int, status int) (b bool) {
	qc := r.source.Table("BIZ_USER_CHARACTER")
	qc.Model(&datamodels.BizUserCharacter{}).Where("CH_ID = ?", chId).Updates(datamodels.BizUserCharacter{CurrentStatus: status})
	return true
}
