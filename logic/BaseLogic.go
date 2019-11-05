package logic

import (
	"fmt"
	"iris/datasource/mybatis"
)

func DoCheckDBURI() bool{
	if mybatis.MysqlUri == "" || mybatis.MysqlUri == "*" {
		fmt.Println("no database url define in MysqlConfig.go , you must set the mysql link!")
		return false
	} else {
		return true
	}
}