package controllers

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/sessions"
	"iris/services"
)

type CharacterController struct {
	Ctx              iris.Context
	Session          *sessions.Session
	CharacterService services.CharacterService
}

func (c *CharacterController) GetChpbBy(chId int) {
	allTaskList := c.CharacterService.GetCharacterPropertyDataByChId(chId)
	c.Ctx.JSON(allTaskList)
}

func (c *CharacterController) GetMe() {
	c.Ctx.Values().Set("message", "character me is here!")
}
