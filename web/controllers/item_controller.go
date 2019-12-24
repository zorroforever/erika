package controllers

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/sessions"
	"iris/commons"
	"iris/services"
)

type ItemController struct {
	Ctx     iris.Context
	Session *sessions.Session

	UserService services.UserService
	ItemService services.ItemService
}

//func (c *ItemController) GetAllItemList() mvc.Result {
//	userId := c.Session.Get(userIDKey)
//	if _, err := c.ItemService.GetAllItemList(); err != nil {
//		return common.MvcError(err.Error(), c.Ctx)
//	}
//	return mvc.Response{
//		Path: "/user/me",
//	}
//}
/**
获取角色道具栏 道具列表
*/
func (c *ItemController) GetItemlistBy(chId int) {
	itemList := c.ItemService.GetItemListByChId(chId)
	response := commons.NewResponse(itemList)
	c.Ctx.JSON(response)
}

/**
获取角色装备栏 装备列表
*/
func (c *ItemController) GetEquipmentlistBy(chId int) {
	equipmentList := c.ItemService.GetEquipmentListByChId(chId)
	response := commons.NewResponse(equipmentList)
	c.Ctx.JSON(response)
}
