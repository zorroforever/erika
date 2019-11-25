package datamodels

/**
  道具属性表
*/
type BizItemLib struct {
	ID         int    `json:"id" form:"id" gorm:"Column:ID"`
	Uuid       string `json:"uuid" form:"uuid" gorm:"Column:UUID"`
	ItemCode   int    `json:"itemCode" form:"itemCode" gorm:"Column:ITEM_CODE"`
	RoleId     int    `json:"roleId" form:"roleId" gorm:"Column:ROLE_ID"`
	ItemStatus int    `json:"itemStatus" form:"itemStatus" gorm:"Column:ITEM_STATUS"`
	CmnDBCol
}
