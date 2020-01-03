package controllers

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/sessions"
	"iris/commons"
	"iris/services"
)

type MapController struct {
	Ctx         iris.Context
	Session     *sessions.Session
	MapService  services.MapService
	TaskService services.TaskService
	UserService services.UserService
}

func (c *MapController) GetTasklistBy(mapId int) {
	kv := c.TaskService.GetTaskListByMapId(mapId)
	var response commons.Response
	if kv == nil {
		response = commons.NewErrorResponse(-1, "没有任务在该地图发布。")
	} else {
		response = commons.NewResponse(kv)
	}
	c.Ctx.JSON(response)

}

func (c *MapController) GetCharacterposBy(chId int) {
	res := c.UserService.GetCharacterposBy(chId)
	c.Ctx.JSON(res)
}
