package datamodels

// User is our User example model.
// Keep note that the tags for public-use (for our web app)
// should be kept in other file like "web/viewmodels/biz_user.go"
// which could wrap by embedding the datamodels.User or
// define completely new fields instead but for the shake
// of the example, we will use this datamodel
// as the only one User model in our application.
type Biz_user struct {
	ID       int64  `json:"id" form:"id" gorm:"Column:ID"`
	Name     string `json:"name" form:"name" gorm:"Column:NAME"`
	Password string `json:"password" form:"password" gorm:"Column:PASSWORD"`
	Token    string `json:"token" form:"token" gorm:"Column:TOKEN"`
	InUse    int    `json:"inUse" form:"inUse" gorm:"Column:IN_USE"`
}

// IsValid can do some very very simple "low-level" data validations.
func (u Biz_user) IsValid() bool {
	return u.ID > 0
}
