// file: controllers/user_controller.go

package controllers

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
	"iris/common"
	"iris/services"
)

// UserController is our /user controller.
// UserController is responsible to handle the following requests:
// GET  			/user/register
// POST 			/user/register
// GET 				/user/login
// POST 			/user/login
// GET 				/user/me
// All HTTP Methods /user/logout
type TaskController struct {
	Ctx     iris.Context
	Session *sessions.Session

	UserService services.UserService
	TaskService services.TaskService
}

func (c *TaskController) GetScrambleTaskBy(taskId int) mvc.Result {
	userId := c.Session.Get(userIDKey)
	if _, err := c.TaskService.ScrambleTask(userId.(int64), taskId); err != nil {
		return common.MvcError(err.Error(), c.Ctx)
	}
	return mvc.Response{
		Path: "/user/me",
	}
}
