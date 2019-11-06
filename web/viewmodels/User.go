package viewmodels

import "iris/datamodels"

type User struct {
	datamodels.Biz_user
}

func (m User) IsValid() bool {
	/* do some checks and return true if it's valid... */
	return m.ID > 0
}
