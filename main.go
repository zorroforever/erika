package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"iris/commons"
	"iris/services"
	"iris/web/controllers"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

//var xmlBytes = []byte(`
//<?xml version="1.0" encoding="UTF-8"?>
//<!DOCTYPE mapper PUBLIC "-//mybatis.org//DTD Mapper 3.0//EN"
//"https://raw.githubusercontent.com/zhuxiujia/GoMybatis/master/mybatis-3-mapper.dtd">
//<mapper>
//    <select id="SelectAll">
//        SELECT C_APP_ID,N_APP_ID FROM BIZ_APP_COMMON
//    </select>
//</mapper>
//`)

//func get_ly_datas()  (list [] map[string]string){
//	//dbURI := fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=Local",userName, password, ip,dbName)
//	//db, err := gorm.Open("mysql", dbURI)
//	//checkErr(err)
//	//db.First(&user)
//	//defer db.Close()
//	mybatis.FinsUserData()
//	return nil
//}

//主函数
func main() {
	//app := iris.Default()
	//
	//// 路由器
	//controller.DoRouter(app)
	//
	////启动服务监听本地8080端口
	//app.Run(iris.Addr("127.0.0.1:8080"))
	app := iris.New()
	// You got full debug messages, useful when using MVC and you want to make
	// sure that your code is aligned with the Iris' MVC Architecture.
	app.Logger().SetLevel("debug")

	// Load the template files.
	tmpl := iris.HTML("./web/views", ".html").
		Layout("shared/layout.html").
		Reload(true)
	app.RegisterView(tmpl)

	app.HandleDir("/public", "./web/public")

	app.OnAnyErrorCode(func(ctx iris.Context) {
		ctx.ViewData("Message", ctx.Values().
			GetStringDefault("message", "The page you're looking for doesn't exist"))
		ctx.View("shared/error.html")
	})

	// 初始化controller
	initRoute(app)

	app.UseGlobal(before)
	app.DoneGlobal(after)

	app.Run(
		// Starts the web server at localhost:8080
		iris.Addr("localhost:8080"),
		// Ignores err server closed log when CTRL/CMD+C pressed.
		iris.WithoutServerError(iris.ErrServerClosed),
		// Enables faster json serialization and more.
		iris.WithOptimizations,
	)
}

func before(ctx iris.Context) {
	println(fmt.Sprintf("Before the handler: %s", ctx.Path()))
	ctx.Next()
}

func after(ctx iris.Context) {
	println(fmt.Sprintf("After the handler: %s", ctx.Path()))
}

func initRoute(app *iris.Application) {
	// ---- Serve our controllers. ----

	// Prepare our repositories and services.
	//db, err := datasource.LoadUsers(datasource.MySQL)
	//if err != nil {
	//	app.Logger().Fatalf("error while loading the users: %v", err)
	//	return
	//}
	userService := services.NewUserService()
	taskService := services.NewTaskService()
	itemService := services.NewItemService()

	// "/users" based mvc application.
	users := mvc.New(app.Party("/users"))
	// Add the basic authentication(admin:password) middleware
	// for the /users based requests.
	//users.Router.Use(middleware.BasicAuth)
	// Bind the "userService" to the UserController's Service (interface) field.
	users.Register(userService)
	users.Handle(new(controllers.UsersController))

	// "/user" based mvc application.
	user := mvc.New(app.Party("/user"))
	user.Register(
		userService,
		taskService,
		commons.SessManager.Start,
	)
	user.Handle(new(controllers.UserController))

	task := mvc.New(app.Party("/task"))
	task.Register(
		userService,
		taskService,
		commons.SessManager.Start,
	)
	task.Handle(new(controllers.TaskController))

	item := mvc.New(app.Party("/item"))
	item.Register(
		userService,
		itemService,
		commons.SessManager.Start,
	)
	item.Handle(new(controllers.ItemController))

	// http://localhost:8080/noexist
	// and all controller's methods like
	// http://localhost:8080/users/1
	// http://localhost:8080/user/register
	// http://localhost:8080/user/login
	// http://localhost:8080/user/me
	// http://localhost:8080/user/logout
	// basic auth: "admin", "password", see "./middleware/basicauth.go" source file.
}
