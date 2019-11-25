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
}

func NewItemService(repo repositories.ItemRepository) ItemService {
	return &itemService{
		repo: repo,
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
