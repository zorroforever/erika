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

type Response struct {
	Success    bool   `json:"success"`
	ErrCode    string `json:"errCode"`
	ErrMessage string `json:"errMessage"`
	Result     Page   `json:"result"`
}

type Page struct {
	Data       interface{} `json:"data"`
	PageNo     int         `json:"pageNo"`
	PageSize   int         `json:"pageSize"`
	TotalCount int         `json:"totalCount"`
	TotalPage  int         `json:"totalPage"`
}

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
