// file: controllers/user_controller.go

package controllers

import (
	"errors"
	uuid "github.com/iris-contrib/go.uuid"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
	"iris/commons"
	"iris/datamodels"
	"iris/web/viewmodels"
	"strings"

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
type UserController struct {
	// context is auto-binded by Iris on each request,
	// remember that on each incoming request iris creates a new UserController each time,
	// so all fields are request-scoped by-default, only dependency injection is able to set
	// custom fields like the Service which is the same for all requests (static binding)
	// and the Session which depends on the current context (dynamic binding).
	Ctx iris.Context

	// Our UserService, it's an interface which
	// is binded from the main application.
	Service services.UserService

	TaskService services.TaskService

	// Session, binded using dependency injection from the main.go.
	Session *sessions.Session
}

const (
	userIDKey = "UserID"
)

var registerStaticView = mvc.View{
	Name: "user/register.html",
	Data: iris.Map{"Title": "User Registration"},
}

var loginStaticView = mvc.View{
	Name: "user/login.html",
	Data: iris.Map{"Title": "User Login"},
}

func (c *UserController) getCurrentUserID() int64 {
	userID := c.Session.GetInt64Default(userIDKey, 0)
	return userID
}

func (c *UserController) isLoggedIn() bool {
	return c.getCurrentUserID() > 0
}

// GetRegister handles GET: http://localhost:8080/user/register.
func (c *UserController) GetRegister() mvc.Result {
	if c.isLoggedIn() {
		c.logout()
	}

	return registerStaticView
}

// PostRegister handles POST: http://localhost:8080/user/register.
func (c *UserController) PostRegister() (mvc.Result, error) {
	var username = c.Ctx.PostValue("username")
	var password = c.Ctx.PostValue("password")
	var rePassword = c.Ctx.PostValue("rePassword")
	if strings.Compare(password, rePassword) != 0 {
		return nil, errors.New("密码不一致")
	}
	user := datamodels.BizUser{
		Name:     username,
		Password: password,
		InUse:    1,
	}
	c.Service.CreateUser(user)
	return loginStaticView, nil
}

// PostRegister handles POST: http://localhost:8080/user/register.
//func (c *UserController) PostRegister() mvc.Result {
// get firstname, username and password from the form.
//var (
//	firstname = c.Ctx.FormValue("firstname")
//	username  = c.Ctx.FormValue("username")
//	password  = c.Ctx.FormValue("password")
//)

// create the new user, the password will be hashed by the service.
//u, err := c.Service.Create(password, datamodels.User{
//	Username:  username,
//	Firstname: firstname,
//})

// set the user's id to this session even if err != nil,
// the zero id doesn't matters because .getCurrentUserID() checks for that.
// If err != nil then it will be shown, see below on mvc.Response.Err: err.
//c.Session.Set(userIDKey, u.ID)

//return mvc.Response{
// if not nil then this error will be shown instead.
//Err: err,
// redirect to /user/me.
//Path: "/user/me",
// When redirecting from POST to GET request you -should- use this HTTP status code,
// however there're some (complicated) alternatives if you
// search online or even the HTTP RFC.
// Status "See Other" RFC 7231, however iris can automatically fix that
// but it's good to know you can set a custom code;
// Code: 303,
//	}
//}

// GetLogin handles GET: http://localhost:8080/user/login.
func (c *UserController) GetLogin() mvc.Result {
	if c.isLoggedIn() {
		// if it's already logged in then destroy the previous session.
		c.logout()
	}

	return loginStaticView
}

//PostLogin handles POST: http://localhost:8080/user/login.
func (c *UserController) PostLogin() error {
	user := &viewmodels.User{}
	c.Ctx.ReadJSON(user)
	u, characters, found := c.Service.GetByUsernameAndPassword(user.Name, user.Password)

	if !found {
		return errors.New("用户不存在")
	}

	user.Characters = characters
	token, _ := uuid.NewV4()
	user.Token = token.String()
	c.Session.Set(userIDKey, u.ID)
	c.Ctx.Next()
	response := commons.NewResponse(user)
	c.Ctx.JSON(response)
	return nil
}

// GetMe handles GET: http://localhost:8080/user/me.
func (c *UserController) GetMe() mvc.Result {
	if !c.isLoggedIn() {
		// if it's not logged in then redirect user to the login page.
		return mvc.Response{Path: "/user/login"}
	}

	u, found := c.Service.GetByID(c.getCurrentUserID())
	if !found {
		// if the  session exists but for some reason the user doesn't exist in the "database"
		// then logout and re-execute the function, it will redirect the client to the
		// /user/login page.
		c.logout()
		return c.GetMe()
	}
	taskList := c.TaskService.GetAllTaskList()
	c.Ctx.Next()
	return mvc.View{
		Name: "user/me.html",
		Data: iris.Map{
			"Title":    "Profile of " + u.Name,
			"User":     u,
			"TaskList": taskList.Data,
		},
	}
}

func (c *UserController) logout() {
	c.Session.Destroy()
}

// AnyLogout handles All/Any HTTP Methods for: http://localhost:8080/user/logout.
func (c *UserController) AnyLogout() {
	if c.isLoggedIn() {
		c.logout()
	}

	c.Ctx.Redirect("/user/login")
}
