package mybatis

import (
	"github.com/jinzhu/gorm"
	"time"
)

//mysql链接格式为         用户名:密码@(数据库链接地址:端口)/数据库名称   例如root:root@(127.0.0.1:3306)/test?charset=utf8&parseTime=True&loc=Local
//此处请按格式填写你的mysql链接，这里用*号代替
const MysqlUri = "robin:robin@(192.168.112.17)/gravity_product?charset=utf8&parseTime=True&loc=Local"

//定义数据库模型
//例子：Activity 活动数据
type Activity struct {
	Id         string    `json:"id,omitempty"`
	Uuid       string    `json:"uuid"`
	Name       string    `json:"name"`
	PcLink     string    `json:"pcLink"`
	H5Link     string    `json:"h5Link"`
	Remark     string    `json:"remark"`
	Version    int       `json:"version"`
	CreateTime time.Time `json:"createTime"`
	DeleteFlag int       `json:"deleteFlag"`
}

type Activity2 struct {
	Id         string    `json:"id"`
	Uuid       string    `json:"uuid,omitempty"`
	Name       string    `json:"name"`
	PcLink     string    `json:"pc_link"`
	H5Link     string    `json:"h5_link"`
	Remark     string    `json:"remark"`
	Version    int       `json:"version"`
	CreateTime time.Time `json:"create_time"`
	DeleteFlag int       `json:"delete_flag"`
}
type User struct {
	gorm.Model
	Salt      string `gorm:"type:varchar(255)" json:"salt"`
	Username  string `gorm:"type:varchar(32)" json:"username"`
	Password  string `gorm:"type:varchar(200);column:password" json:"-"`
	Languages string `gorm:"type:varchar(200);column:languages" json:"languages"`
}

func (u User) TableName() string {
	return "gorm_user"
}

type UserSerializer struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Salt      string    `json:"salt"`
	UserName  string    `json:"user_name"`
	Password  string    `json:"-"`
	Languages string    `json:"languages"`
}

func (self User) Serializer() UserSerializer {
	return UserSerializer{
		ID:        self.ID,
		CreatedAt: self.CreatedAt.Truncate(time.Second),
		UpdatedAt: self.UpdatedAt.Truncate(time.Second),
		Salt:      self.Salt,
		Password:  self.Password,
		Languages: self.Languages,
		UserName:  self.Username,
	}
}