package model

//UserData
type UserData struct {
	Id         int    `json:"id,omitempty"`
	Uuid       string    `json:"uuid"`
	Name       string    `json:"name"`
	Token      string    `json:"token"`
	InUse      int    `json:"inUse"`
}

