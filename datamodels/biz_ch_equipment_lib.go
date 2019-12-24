package datamodels

/**
  角色装备表
*/
type BizChEquipmentLib struct {
	ID          int `json:"id" form:"id" gorm:"Column:ID"`
	UserId      int `json:"userId" form:"userId" gorm:"Column:USER_ID"`
	ChId        int `json:"chId" form:"chId" gorm:"Column:CH_ID"`
	EquipmentId int `json:"equipmentId" form:"equipmentId" gorm:"Column:EQUIPMENT_ID"`
	Durability  int `json:"durability" form:"durability" gorm:"Column:DURABILITY"`
	CmnDBCol
}
