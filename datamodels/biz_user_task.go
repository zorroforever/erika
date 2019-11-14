package datamodels

type BizUserTask struct {
	UserId     int     `json:"userId" form:"userId" gorm:"Column:USER_ID"`
	TaskId     int     `json:"taskId" form:"taskId" gorm:"Column:TASK_ID"`
	Coin       float32 `json:"coin" form:"coin" gorm:"Column:COIN"`
	Experience int     `json:"experience" form:"experience" gorm:"Column:EXPERIENCE"`
	Honor      int     `json:"honor" form:"honor" gorm:"Column:HONOR"`
	Log        string  `json:"log" form:"log" gorm:"Column:LOG"`
	Remark     string  `json:"remark" form:"remark" gorm:"Column:REMARK"`
	CreateUser string  `json:"createUser" form:"createUser" gorm:"Column:CREATE_USER"`
	CreateTime string  `json:"createTime" form:"createTime" gorm:"Column:CREATE_TIME"`
	UpdateUser string  `json:"updateUser" form:"updateUser" gorm:"Column:UPDATE_USER"`
	UpdateTime string  `json:"updateTime" form:"updateTime" gorm:"Column:UPDATE_TIME"`
}
