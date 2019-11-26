package services

import (
	"fmt"
	uuid "github.com/iris-contrib/go.uuid"
	"iris/datamodels"
	"iris/repositories"
)

type ItemService interface {
	GetAllItemList() []datamodels.BizItem
	GetItemById(itemId int) datamodels.BizItem
	CreateNewItemById(itemId int) (bool, datamodels.BizItemLib)
	GetItemListByRoleId(roleId int) []datamodels.MItem
}

func NewItemService() ItemService {
	return &itemService{
		repo: repositories.NewItemDBRep(),
	}
}

type itemService struct {
	repo repositories.ItemRepository
}

func (s *itemService) GetAllItemList() []datamodels.BizItem {
	return s.repo.GetAllItemList()
}

func (s *itemService) GetItemById(itemId int) datamodels.BizItem {
	return s.repo.GetItemById(itemId)
}

func (s *itemService) CreateNewItemById(itemId int) (bool, datamodels.BizItemLib) {
	u2, err := uuid.NewV4()
	if err != nil {
		fmt.Printf("Something went wrong: %s", err)
		return false, datamodels.BizItemLib{
			ID:         0,
			Uuid:       "",
			ItemCode:   0,
			RoleId:     0,
			ItemStatus: 0,
			CmnDBCol:   datamodels.CmnDBCol{},
		}
	}
	return s.repo.CreateNewItemById(itemId, u2.String())
}

func (s *itemService) GetItemListByRoleId(roleId int) (rMItem []datamodels.MItem) {
	// 1. 获取item列表
	count, itemList := s.repo.GetItemListByRoleId(roleId)
	if count > 0 {
		// 2。补充item信息
		rMItem = make([]datamodels.MItem, count)
		for i, item := range itemList {
			code := item.ItemCode
			itemDetail := s.repo.GetItemById(code)
			rMItem[i] = datamodels.MItem{
				ID:              code,
				ItemType:        itemDetail.ItemType,
				ItemName:        itemDetail.ItemName,
				ItemQuality:     itemDetail.ItemQuality,
				ItemDetail:      itemDetail.ItemDetail,
				ItemInvalidTime: itemDetail.ItemInvalidTime,
				ItemStatus:      item.ItemStatus,
				TimeLimit:       itemDetail.TimeLimit,
				ItemEffect:      itemDetail.ItemEffect,
				ItemMaxCount:    itemDetail.ItemMaxCount,
				Uuid:            item.Uuid,
			}
		}
	}
	return rMItem
}
