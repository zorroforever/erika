package repositories

import (
	"github.com/jinzhu/gorm"
	"iris/datamodels"
	"iris/datasource"
)

type ItemRepository interface {
	// 获取所有物品列表
	GetAllItemList() []datamodels.BizItem
	// 按id获取物品
	GetItemById(itemId int) datamodels.BizItem
	// 生成唯一道具ID
	CreateNewItemById(itemId int, uuid string) (bool, datamodels.BizItemLib)
}

func NewItemDBRep(source *gorm.DB) ItemRepository {
	return &itemSQLRepository{source: datasource.DB}
}

type itemSQLRepository struct {
	source *gorm.DB
}

func (r *itemSQLRepository) GetAllItemList() (bizItem []datamodels.BizItem) {
	qc := r.source.Table("BIZ_ITEM").Model(&datamodels.BizItem{})
	qc.Find(&bizItem)
	return bizItem
}

func (r *itemSQLRepository) GetItemById(taskId int) (item datamodels.BizItem) {
	qc := r.source.Table("BIZ_ITEM").Model(&datamodels.BizItem{})
	qc.Where("ID = ?", taskId).Find(&item)
	return item
}

func (r *itemSQLRepository) CreateNewItemById(itemId int, uuid string) (flg bool, item datamodels.BizItemLib) {

	qc := r.source.Table("BIZ_ITEM_LIB").Model(&datamodels.BizItemLib{})

	itemLib := datamodels.BizItemLib{Uuid: uuid, ItemCode: itemId, ItemStatus: 0, CmnDBCol: datamodels.CmnDBCol{CreateTime: "now()"}}
	qc.NewRecord(itemLib) // => 主键为空返回`true`
	qc.Create(&itemLib)

	return true, itemLib
}
