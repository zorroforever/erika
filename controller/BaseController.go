package controller

import (
	"github.com/kataras/iris/v12"
)
/**
路由统一注册
 */
func DoRouter(app *iris.Application){
	app.HandleDir("/web", "web")


	tmpl := iris.HTML("web/views", ".html")
	tmpl.Reload(true)
	tmpl.AddFunc("greet", func(s string) string {
		return "Greetings " + s + "!"
	})

	// Register the views engine to the views,
	// this will load the templates.
	app.RegisterView(tmpl)
	app.Get("/",func(ctx iris.Context) {
		ctx.View("index.html")
	} )

	DoFetchUser(app)
}
