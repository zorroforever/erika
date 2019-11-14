package datamodels

type BizTask struct {
	ID         int     `json:"id" form:"id" gorm:"Column:ID"`
	Name       string  `json:"name" form:"name" gorm:"Column:NAME"`
	Type       int     `json:"type" form:"type" gorm:"Column:TYPE"`
	Content    string  `json:"content" form:"content" gorm:"Column:CONTENT"`
	Coin       float32 `json:"coin" form:"coin" gorm:"Column:COIN"`
	Experience int     `json:"experience" form:"experience" gorm:"Column:EXPERIENCE"`
	Honor      int     `json:"honor" form:"honor" gorm:"Column:HONOR"`
	TimeLimit  int     `json:"timeLimit" form:"timeLimit" gorm:"Column:TIME_LIMIT"`
	Remark     string  `json:"remark" form:"remark" gorm:"Column:REMARK"`
}
