package repositories

import (
	"github.com/jinzhu/gorm"
	"iris/datamodels"
)

// UserRepository handles the basic operations of a user entity/model.
// It's an interface in order to be testable, i.e a memory user repository or
// a connected to an sql database.
type UserRepository interface {
	//Exec(query Query, action Query, limit int, mode int) (ok bool)

	SelectAll() (users []datamodels.Biz_user, total int)
	//SelectMany(query Query, limit int) (results []datamodels.User)
	//
	//InsertOrUpdate(user datamodels.User) (updatedUser datamodels.User, err error)
	//Delete(query Query, limit int) (deleted bool)
}
func NewUserDBRep(source *gorm.DB) UserRepository{
	return &userSQLRepository{source: source}
}

type userSQLRepository struct {
	source *gorm.DB
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

	qc :=r.source.Model(&datamodels.Biz_user{})
	qc.Select(&users)
	qc.Count(&total)
	return
}

