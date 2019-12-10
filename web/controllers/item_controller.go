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
func (c *ItemController) GetItemlistBy(chId int) {
	itemList := c.ItemService.GetItemListByRoleId(chId)
	response := commons.NewResponse(itemList)
	c.Ctx.JSON(response)
}
