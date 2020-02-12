package controllers

import "C"
import (
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/sessions"
	"iris/commons"
	"iris/datamodels"
	"iris/services"
	"iris/web/viewmodels"
)

type MapController struct {
	Ctx         iris.Context
	Session     *sessions.Session
	MapService  services.MapService
	TaskService services.TaskService
	UserService services.UserService
}

/**
获取地图任务列表
*/
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

/**
获取角色位置信息
*/
func (c *MapController) GetCharacterposBy(chId int) {
	res := c.UserService.GetCharacterposBy(chId)
	response := commons.NewResponse(res)
	c.Ctx.JSON(response)
}

/**
更新角色移动信息
*/
func (c *MapController) PostUpdpms() {
	ss := &viewmodels.ChMoveModel{}
	if err := c.Ctx.ReadJSON(ss); err != nil {
		panic(err.Error())
	} else {
		updss := &datamodels.BizChMoveLib{}
		if ss.ArriveTimeStr != "" {
			updss.ArriveTime = commons.StrCover2Time(ss.ArriveTimeStr)
		}
		updss.SX = ss.SX
		updss.SY = ss.SY
		updss.SMapId = ss.SMapId
		updss.TMapId = ss.TMapId
		updss.TX = ss.TX
		updss.TY = ss.TY
		updss.ChId = ss.ChId
		fmt.Printf("%+v\n", *updss)
		updss.CreateTime = commons.GetNowStr()
		res := c.MapService.DoUpdPersonMoveStatus(*updss)
		response := commons.NewResponse(res)
		c.Ctx.JSON(response)
	}
}
func (c *MapController) PostUpdpmss(chId int, status int) {
	res := c.MapService.DoUpdPersonStatus(chId, status)
	response := commons.NewResponse(res)
	c.Ctx.JSON(response)
}

func (c *MapController) GetCharacterDataBy(chId int) {
	res := c.UserService.GetCharacterDataById(chId)
	response := commons.NewResponse(res)
	fmt.Printf("%+v\n", response)
	c.Ctx.JSON(response)
}
