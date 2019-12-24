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
	// 按id获取装备
	GetEquipmentById(eqpId int) datamodels.BizEquipment
	// 生成唯一道具ID
	CreateNewItemById(itemId int, uuid string) (bool, datamodels.BizItemLib)
	// 按角色id获取道具列表
	GetItemListByChId(chId int) (count int, lib []datamodels.BizItemLib)
	// 按角色id获取装备列表
	GetEquipmentListByChId(chId int) (count int, lib []datamodels.BizChEquipmentLib)
}

func NewItemDBRep() ItemRepository {
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

func (r *itemSQLRepository) GetItemById(id int) (item datamodels.BizItem) {
	qc := r.source.Table("BIZ_ITEM").Model(&datamodels.BizItem{})
	qc.Where("ID = ?", id).Find(&item)
	return item
}

func (r *itemSQLRepository) GetEquipmentById(id int) (item datamodels.BizEquipment) {
	qc := r.source.Table("BIZ_EQUIPMENT").Model(&datamodels.BizItem{})
	qc.Where("ID = ?", id).Find(&item)
	return item
}
func (r *itemSQLRepository) CreateNewItemById(itemId int, uuid string) (flg bool, item datamodels.BizItemLib) {

	qc := r.source.Table("BIZ_CH_ITEM_LIB").Model(&datamodels.BizItemLib{})

	itemLib := datamodels.BizItemLib{Uuid: uuid, ItemCode: itemId, ItemStatus: 0, CmnDBCol: datamodels.CmnDBCol{CreateTime: "now()"}}
	qc.NewRecord(itemLib)
	qc.Create(&itemLib)

	return true, itemLib
}

func (r *itemSQLRepository) GetItemListByChId(chId int) (cnt int, itemLib []datamodels.BizItemLib) {
	qc := r.source.Table("BIZ_CH_ITEM_LIB").Model(&datamodels.BizItemLib{})
	qc.Where("CH_ID = ?", chId).Count(&cnt)
	if cnt > 0 {
		qc.Where("CH_ID = ?", chId).Find(&itemLib)
	}
	return cnt, itemLib
}

func (r *itemSQLRepository) GetEquipmentListByChId(chId int) (cnt int, itemLib []datamodels.BizChEquipmentLib) {
	qc := r.source.Table("BIZ_CH_EQUIPMENT_LIB").Model(&datamodels.BizChEquipmentLib{})
	qc.Where("CH_ID = ?", chId).Count(&cnt)
	if cnt > 0 {
		qc.Where("CH_ID = ?", chId).Find(&itemLib)
	}
	return cnt, itemLib
}
