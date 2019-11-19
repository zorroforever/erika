package datasource

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/kataras/iris/v12"
	"time"
)

var DB *gorm.DB

const (
	Username = "robin"
	Password = "robin"
	Addr     = "192.168.7.197"
	Name     = "kazenotani"
)

func init() {
	println("初始化数据库开始")
	DB = InitDB()
	println("初始化数据库结束")
}

func InitDB() (dbc *gorm.DB) {
	config := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=%t&loc=%s",
		Username,
		Password,
		Addr,
		Name,
		true,
		//"Asia/Shanghai"),
		"Local")
	db, err := gorm.Open("mysql", config)
	db.LogMode(true) // show SQL logger
	if err != nil {
		panic("failed to connect database")
	}
	db.DB().SetMaxIdleConns(2) // 用于设置闲置的连接数.设置闲置的连接数则当开启的一个连接使用完成后可以放在池里等候下一次使用。
	db.SingularTable(true)     //设置表名不为负数
	go keepAlive(db)
	iris.RegisterOnInterrupt(func() {
		defer db.Close()
	})
	return db
}

func keepAlive(dbc *gorm.DB) {
	for {
		dbc.DB().Ping()
		time.Sleep(60 * time.Second)
	}
}
