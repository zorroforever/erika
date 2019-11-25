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

// 前端每页每页页数
var PageSize = 10

type Response struct {
	Success    bool        `json:"success"`
	ErrCode    int         `json:"errCode"`
	ErrMessage string      `json:"errMessage"`
	Result     interface{} `json:"result"`
}

type Page struct {
	Data       interface{} `json:"data"`
	PageNo     int         `json:"pageNo"`
	PageSize   int         `json:"pageSize"`
	TotalCount int         `json:"totalCount"`
	TotalPage  int         `json:"totalPage"`
}

func NewResponse(data interface{}) Response {
	json := Response{
		Success:    true,
		ErrCode:    iris.StatusOK,
		ErrMessage: "",
		Result:     data,
	}
	return json
}

func NewPageResponse(page Page) Response {
	json := Response{
		Success:    true,
		ErrCode:    iris.StatusInternalServerError,
		ErrMessage: "",
		Result:     page,
	}
	return json
}

func NewErrorResponse(errCode int, errMessage string) Response {
	json := Response{
		Success:    false,
		ErrCode:    errCode,
		ErrMessage: errMessage,
		Result:     nil,
	}
	return json
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
