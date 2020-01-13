package commons

import (
	"database/sql/driver"
	"fmt"
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

type KvModel struct {
	Key   string `json:"key"`
	Value string `json:"value"`
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

func FloatToString(input_num float32) string {
	// to convert a float number to a string
	return strconv.FormatFloat(float64(input_num), 'f', 6, 64)
}

type LocalTime struct {
	time.Time
}

func (t LocalTime) MarshalJSON() ([]byte, error) {
	//格式化秒
	seconds := t.Unix()
	return []byte(strconv.FormatInt(seconds, 10)), nil
}
func (t LocalTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}
func (t *LocalTime) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = LocalTime{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}

func GetNowStr() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

const CH_FREE int = 1
const CH_MOVING int = 2
const CH_TASKING int = 3
