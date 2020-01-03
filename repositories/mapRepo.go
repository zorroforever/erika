package repositories

import (
	"github.com/jinzhu/gorm"
	"iris/datasource"
)

type MapRepository interface {
}

func NewMapDBRepo() MapRepository {
	return &mapSQLRepository{source: datasource.DB}
}

type mapSQLRepository struct {
	source *gorm.DB
}
