package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm"
	"github.com/kataras/iris/v12"
	"iris/datasource"
	"iris/repositories"
	"iris/services"
	"iris/web/controllers"

	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
	"time"
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

	config := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=%t&loc=%s",
		datasource.Username,
		datasource.Password,
		datasource.Addr,
		datasource.Name,
		true,
		//"Asia/Shanghai"),
		"Local")
	db, err := gorm.Open("mysql", config)
	db.LogMode(true) // show SQL logger
	if err != nil {
		app.Logger().Fatalf("connect to mysql failed")
		return
	}
	db.DB().SetMaxIdleConns(2) // 用于设置闲置的连接数.设置闲置的连接数则当开启的一个连接使用完成后可以放在池里等候下一次使用。
	db.SingularTable(true)     //设置表名不为负数
	go keepAlive(db)
	iris.RegisterOnInterrupt(func() {
		defer db.Close()
	})

	// ---- Serve our controllers. ----

	// Prepare our repositories and services.
	//db, err := datasource.LoadUsers(datasource.MySQL)
	//if err != nil {
	//	app.Logger().Fatalf("error while loading the users: %v", err)
	//	return
	//}
	userRepo := repositories.NewUserDBRep(db)
	userService := services.NewUserService(userRepo)

	taskRepo := repositories.NewTaskDBRep(db)
	taskService := services.NewTaskService(taskRepo)

	// "/users" based mvc application.
	users := mvc.New(app.Party("/users"))
	// Add the basic authentication(admin:password) middleware
	// for the /users based requests.
	//users.Router.Use(middleware.BasicAuth)
	// Bind the "userService" to the UserController's Service (interface) field.
	users.Register(userService)
	users.Handle(new(controllers.UsersController))

	// "/user" based mvc application.
	sessManager := sessions.New(sessions.Config{
		Cookie:  "kazenotani",
		Expires: 24 * time.Hour,
	})
	user := mvc.New(app.Party("/user"))
	user.Register(
		userService,
		taskService,
		sessManager.Start,
	)
	user.Handle(new(controllers.UserController))

	task := mvc.New(app.Party("/task"))
	task.Register(
		userService,
		taskService,
		sessManager.Start,
	)
	task.Handle(new(controllers.TaskController))

	// http://localhost:8080/noexist
	// and all controller's methods like
	// http://localhost:8080/users/1
	// http://localhost:8080/user/register
	// http://localhost:8080/user/login
	// http://localhost:8080/user/me
	// http://localhost:8080/user/logout
	// basic auth: "admin", "password", see "./middleware/basicauth.go" source file.
	app.Run(
		// Starts the web server at localhost:8080
		iris.Addr("localhost:8080"),
		// Ignores err server closed log when CTRL/CMD+C pressed.
		iris.WithoutServerError(iris.ErrServerClosed),
		// Enables faster json serialization and more.
		iris.WithOptimizations,
	)
}
func keepAlive(dbc *gorm.DB) {
	for {
		dbc.DB().Ping()
		time.Sleep(60 * time.Second)
	}
}
