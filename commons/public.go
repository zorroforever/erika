package commons

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
	"strconv"
	"time"
)

var SessManager = sessions.New(sessions.Config{
	Cookie:  "kazenotani",
	Expires: 24 * time.Hour,
})

/*
MVC错误页面显示
*/
func MvcError(msg string, ctx iris.Context) mvc.View {
	return mvc.View{
		Name: "shared/error.html",
		Data: iris.Map{
			"Title":   "请求异常",
			"Code":    strconv.Itoa(ctx.GetStatusCode()),
			"Message": msg,
		},
	}
}
