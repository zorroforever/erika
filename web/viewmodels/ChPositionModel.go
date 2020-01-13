package viewmodels

type ChPositionModel struct {
	MapId      int    `json:"mapId"`
	PosX       int    `json:"posX"`
	PosY       int    `json:"posY"`
	ArriveTime string `json:"arriveTime"`
}
