package datamodels

type MItem struct {
	ID              int    `json:"id" form:"id" gorm:"Column:ID"`
	ItemType        int    `json:"itemType" form:"itemType" gorm:"Column:ITEM_TYPE"`
	ItemUrl         string `json:"itemUrl" form:"itemUrl" gorm:"Column:ITEM_URL"`
	ItemName        string `json:"itemName" form:"itemName" gorm:"Column:ITEM_NAME"`
	ItemQuality     int    `json:"itemQuality" form:"itemQuality" gorm:"Column:ITEM_QUALITY"`
	ItemDetail      string `json:"itemDetail" form:"itemDetail" gorm:"Column:ITEM_DETAIL"`
	ItemInvalidTime int    `json:"itemInvalidTime" form:"itemInvalidTime" gorm:"Column:ITEM_INVALID_TIME"`
	ItemStatus      int    `json:"itemStatus" form:"itemStatus" gorm:"Column:ITEM_STATUS"`
	TimeLimit       int    `json:"timeLimit" form:"timeLimit" gorm:"Column:TIME_LIMIT"`
	ItemEffect      string `json:"itemEffect" form:"itemEffect" gorm:"Column:ITEM_EFFECT"`
	ItemMaxCount    string `json:"itemMaxCount" form:"itemMaxCount" gorm:"Column:ITEM_MAX_COUNT"`
	Uuid            string `json:"uuid" form:"uuid" gorm:"Column:UUID"`
}
