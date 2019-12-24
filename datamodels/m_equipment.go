package datamodels

type MEqpItem struct {
	ID          int     `json:"id" form:"id" gorm:"Column:ID"`
	Type        int     `json:"type" form:"type" gorm:"Column:TYPE"`
	Name        string  `json:"name" form:"name" gorm:"Column:NAME"`
	HelPoint    int     `json:"helPoint" form:"helPoint" gorm:"Column:HEL_POINT"`
	ManaPoint   int     `json:"manaPoint" form:"manaPoint" gorm:"Column:MANA_POINT"`
	PhyAttack   float32 `json:"phyAttack" form:"phyAttack" gorm:"Column:PHY_ATTACK"`
	ManaAttack  float32 `json:"manaAttack" form:"manaAttack" gorm:"Column:MANA_ATTACK"`
	PhyDefense  float32 `json:"phyDefense" form:"phyDefense" gorm:"Column:PHY_DEFENSE"`
	ManaDefense float32 `json:"manaDefense" form:"manaDefense" gorm:"Column:MANA_DEFENSE"`
	Evade       float32 `json:"evade" form:"evade" gorm:"Column:EVADE"`
	Str         int     `json:"str" form:"str" gorm:"Column:STR"`
	Agl         int     `json:"agl" form:"agl" gorm:"Column:AGL"`
	Int         int     `json:"int" form:"int" gorm:"Column:INT"`
	Vit         int     `json:"vit" form:"vit" gorm:"Column:VIT"`
	Dex         int     `json:"dex" form:"dex" gorm:"Column:DEX"`
	Luk         int     `json:"luk" form:"luk" gorm:"Column:LUK"`
}
