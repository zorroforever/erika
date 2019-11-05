package controller

import (
	"github.com/kataras/iris/v12"
	"iris/logic"
)
/**
获取全部用户
 */
func DoFetchUser(app *iris.Application){
	app.Get("/doFetchUser", func(ctx iris.Context) {
		//id := ctx.URLParam("id")
		page := ctx.URLParamDefault("page", "1")
		//name := ctx.FormValue("name")
		message := ctx.FormValue("message")
		var result,err = logic.FetchAllUser()
			if err != nil {
				panic(err)
			}
		ctx.JSON(iris.Map{
			"state":     0,
			"page":      page,
			"message":   message,
			"data_list": result,
		})
	})
}