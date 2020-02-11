package repositories

import (
	"github.com/jinzhu/gorm"
	"iris/datamodels"
	"iris/datasource"
)

// UserRepository handles the basic operations of a user entity/model.
// It's an interface in order to be testable, i.e a memory user repository or
// a connected to an sql database.
type UserRepository interface {
	//Exec(query Query, action Query, limit int, mode int) (ok bool)

	SelectAll() (users []datamodels.Biz_user, total int)
	GetID(id int64) (user datamodels.Biz_user, found bool)
	GetByUsernameAndPassword(username string, password string) (user datamodels.Biz_user, found bool)
	//SelectMany(query Query, limit int) (results []datamodels.User)
	//
	//InsertOrUpdate(user datamodels.User) (updatedUser datamodels.User, err error)
	//Delete(query Query, limit int) (deleted bool)
	CreateUser(user datamodels.Biz_user) (bizUser datamodels.Biz_user, found bool)
	GetChDataByChId(id int) (r datamodels.BizUserCharacter)
	GetChPropertyByChId(id int) (r datamodels.BizChProperty)
	DoCheckMoveStatus(id int) (c string)
	UpdMoveStatus(id int, status int)
}

type userSQLRepository struct {
	source *gorm.DB
}

func NewUserDBRep() UserRepository {
	source := datasource.DB.Table("BIZ_USER")
	return &userSQLRepository{source: source}
}

// Select receives a query function
// which is fired for every single user model inside
// our imaginary data source.
// When that function returns true then it stops the iteration.
//
// It returns the query's return last known boolean value
// and the last known user model
// to help callers to reduce the LOC.
//
// It's actually a simple but very clever prototype function
// I'm using everywhere since I firstly think of it,
// hope you'll find it very useful as well.
func (r *userSQLRepository) SelectAll() (users []datamodels.Biz_user, total int) {

	qc := r.source.Model(&datamodels.Biz_user{})
	qc.Find(&users)
	qc.Count(&total)
	return
}
func (r *userSQLRepository) GetID(id int64) (user datamodels.Biz_user, found bool) {
	qc := r.source.Model(&datamodels.Biz_user{})
	qc.Where("ID = ?", id).Find(&user)
	if user.IsValid() == false {
		found = false
	} else {
		found = true
	}
	return
}

func (r *userSQLRepository) GetByUsernameAndPassword(username string, password string) (user datamodels.Biz_user, found bool) {
	qc := r.source.Model(&datamodels.Biz_user{})
	qc.Where("name = ? AND password = ?", username, password).Find(&user)
	if user.IsValid() {
		found = true
	} else {
		found = false
	}
	return user, found
}

func (r *userSQLRepository) CreateUser(user datamodels.Biz_user) (datamodels.Biz_user, bool) {
	r.source.Create(&user)
	return user, true
}

func (r *userSQLRepository) GetChDataByChId(chId int) (res datamodels.BizUserCharacter) {
	qc := r.source.Table("BIZ_USER_CHARACTER").Model(&datamodels.BizUserCharacter{})
	qc.Where("CH_ID = ?", chId).Find(&res).Limit(1)
	return res
}

func (r *userSQLRepository) DoCheckMoveStatus(chId int) (res string) {
	var bcml datamodels.BizChMoveLib
	qc := r.source.Table("BIZ_CH_MOVE_LIB").Model(&datamodels.BizChMoveLib{})
	qc.Where("ARRIVE_TIME < now() ").Order("ID desc").Find(&bcml).Limit(1)
	return bcml.ArriveTime
}

func (r *userSQLRepository) UpdMoveStatus(chId int, status int) {
	qc := r.source.Table("BIZ_USER_CHARACTER").Model(&datamodels.BizUserCharacter{})
	qc.Where("CH_ID = ?", chId).Updates(datamodels.BizUserCharacter{CurrentStatus: status})
}

func (r *userSQLRepository) GetChPropertyByChId(chId int) (chp datamodels.BizChProperty) {
	qc := r.source.Table("BIZ_CH_PROPERTY").Model(&datamodels.BizChProperty{})
	qc.Where("CH_ID = ?", chId).Find(&chp).Limit(1)
	return chp
}
