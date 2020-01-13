package datamodels

type CmnDBCol struct {
	Remark     string `json:"remark" form:"remark" gorm:"Column:REMARK"`
	CreateUser string `json:"createUser" form:"createUser" gorm:"Column:CREATE_USER"`
	CreateTime string `json:"createTime" form:"createTime" gorm:"Column:CREATE_TIME"`
	UpdateUser string `json:"updateUser" form:"updateUser" gorm:"Column:UPDATE_USER"`
	UpdateTime string `json:"updateTime" form:"updateTime" gorm:"Column:UPDATE_TIME"`
}

type CmnDBCol2 struct {
	CreateUser string `json:"createUser" form:"createUser" gorm:"Column:CREATE_USER"`
	CreateTime string `json:"createTime" form:"createTime" gorm:"Column:CREATE_TIME"`
	UpdateUser string `json:"updateUser" form:"updateUser" gorm:"Column:UPDATE_USER"`
	UpdateTime string `json:"updateTime" form:"updateTime" gorm:"Column:UPDATE_TIME"`
}
