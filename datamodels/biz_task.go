package datamodels

type BizTask struct {
	ID              int     `json:"id" form:"id" gorm:"Column:ID"`
	Name            string  `json:"name" form:"name" gorm:"Column:NAME"`
	Type            int     `json:"type" form:"type" gorm:"Column:TYPE"`
	Content         string  `json:"content" form:"content" gorm:"Column:CONTENT"`
	Coin            float32 `json:"coin" form:"coin" gorm:"Column:COIN"`
	Experience      int     `json:"experience" form:"experience" gorm:"Column:EXPERIENCE"`
	Honor           int     `json:"honor" form:"honor" gorm:"Column:HONOR"`
	TimeLimit       int     `json:"timeLimit" form:"timeLimit" gorm:"Column:TIME_LIMIT"`
	Remark          string  `json:"remark" form:"remark" gorm:"Column:REMARK"`
	Status          int     `json:"Status" form:"Status" gorm:"Column:STATUS"`
	CooperationFlag int     `json:"CooperationFlag" form:"CooperationFlag" gorm:"Column:COOPERATION_FLAG"`
	MapId           int     `json:"mapId" form:"mapId" gorm:"Column:MAP_ID"`
	PointX          int     `json:"pointX" form:"pointX" gorm:"Column:POINT_X"`
	PointY          int     `json:"pointY" form:"pointY" gorm:"Column:POINT_Y"`
	CreateUser      string  `json:"createUser" form:"createUser" gorm:"Column:CREATE_USER"`
	CreateTime      string  `json:"createTime" form:"createTime" gorm:"Column:CREATE_TIME"`
	UpdateUser      string  `json:"updateUser" form:"updateUser" gorm:"Column:UPDATE_USER"`
	UpdateTime      string  `json:"updateTime" form:"updateTime" gorm:"Column:UPDATE_TIME"`
}
