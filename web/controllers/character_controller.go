package controllers

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/sessions"
	"iris/commons"
	"iris/services"
)

type CharacterController struct {
	Ctx              iris.Context
	Session          *sessions.Session
	CharacterService services.CharacterService
}

func (c *CharacterController) GetChpbBy(chId int) {
	kv := c.CharacterService.GetCharacterPropertyDataByChId(chId)
	response := commons.NewResponse(kv)
	c.Ctx.JSON(response)
}

func (c *CharacterController) GetChlistBy(userId int) {
	chIdList := c.CharacterService.GetCharacterIdByUserId(userId)
	response := commons.NewResponse(chIdList)
	c.Ctx.JSON(response)
}

func (c *CharacterController) GetMe() {
	c.Ctx.Values().Set("message", "character me is here!")
}
