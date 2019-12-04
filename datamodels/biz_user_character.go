package datamodels

/**
  用户角色表
*/
type BizUserCharacter struct {
	ID                int     `json:"id" form:"id" gorm:"Column:ID"`
	UserId            int     `json:"userId" form:"userId" gorm:"Column:USER_ID"`
	ChId              int     `json:"chId" form:"chId" gorm:"Column:CH_ID"`
	ChName            int     `json:"chName" form:"chName" gorm:"Column:CH_NAME"`
	CurrentStatus     int     `json:"currentStatus" form:"currentStatus" gorm:"Column:CURRENT_STATUS"`
	CurrentExperience int     `json:"currentExperience" form:"currentExperience" gorm:"Column:CURRENT_EXPERIENCE"`
	Level             int     `json:"level" form:"level" gorm:"Column:LEVEL"`
	Honor             int     `json:"honor" form:"honor" gorm:"Column:HONOR"`
	MapId             string  `json:"mapId" form:"mapId" gorm:"Column:MAP_ID"`
	ManaDefense       float32 `json:"manaDefense" form:"manaDefense" gorm:"Column:MANA_DEFENSE"`
	PointX            int     `json:"pointX" form:"pointX" gorm:"Column:POINT_X"`
	PointY            int     `json:"pointY" form:"pointY" gorm:"Column:POINT_Y"`
	CmnDBCol
}
