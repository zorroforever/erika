package viewmodels

import "iris/datamodels"

type User struct {
	datamodels.BizUser
	Characters []datamodels.BizUserCharacter `json:"characters"`
}

func (m User) IsValid() bool {
	/* do some checks and return true if it's valid... */
	return m.ID > 0
}
