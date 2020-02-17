package viewmodels

type ChDataModel struct {
	MapId    int    `json:"mapId"`
	UserId   int    `json:"userId"`
	ChName   string `json:"chName"`
	Dex      int    `json:"dex"`
	Sx       int    `json:"sx"`
	Sy       int    `json:"sy"`
	CdStart  string `json:"cdStart"`
	CdEnd    string `json:"cdEnd"`
	CdNow    string `json:"cdNow"`
	IsNeedCD int    `json:"isNeedCD"`
}
