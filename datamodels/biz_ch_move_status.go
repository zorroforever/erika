package datamodels

type BizChMoveLib struct {
	ID         int    `json:"id" form:"id" gorm:"Column:ID"`
	ChId       int    `json:"chId" form:"chId" gorm:"Column:CH_ID"`
	SMapId     int    `json:"sMapId" form:"sMapId" gorm:"Column:S_MAP_ID"`
	TMapId     int    `json:"tMapId" form:"tMapId" gorm:"Column:T_MAP_ID"`
	SX         int    `json:"sX" form:"sX" gorm:"Column:S_X"`
	SY         int    `json:"sY" form:"sY" gorm:"Column:S_Y"`
	TX         int    `json:"tX" form:"tX" gorm:"Column:T_X"`
	TY         int    `json:"tY" form:"tY" gorm:"Column:T_Y"`
	ArriveTime string `json:"arriveTime" form:"arriveTime" gorm:"Column:ARRIVE_TIME"`
	CmnDBCol
}
