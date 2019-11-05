package services

import (
	"iris/datamodels"
	"iris/repositories"
)

// UserService handles CRUID operations of a user datamodel,
// it depends on a user repository for its actions.
// It's here to decouple the data source from the higher level compoments.
// As a result a different repository type can be used with the same logic without any aditional changes.
// It's an interface and it's used as interface everywhere
// because we may need to change or try an experimental different domain logic at the future.
type UserService interface {
	GetAll() (users []datamodels.Biz_user,total int)
	//GetByID(id int64) (datamodels.User, bool)
	//GetByUsernameAndPassword(username, userPassword string) (datamodels.User, bool)
	//DeleteByID(id int64) bool
	//
	//Update(id int64, user datamodels.User) (datamodels.User, error)
	//UpdatePassword(id int64, newPassword string) (datamodels.User, error)
	//UpdateUsername(id int64, newUsername string) (datamodels.User, error)
	//
	//Create(userPassword string, user datamodels.User) (datamodels.User, error)
}

// NewUserService returns the default user service.
func NewUserService(repo repositories.UserRepository) UserService {
	return &userService{
		repo: repo,
	}
}

type userService struct {
	repo repositories.UserRepository
}

// GetAll returns all users.
func (s *userService) GetAll() (users []datamodels.Biz_user,total int) {
	return s.repo.SelectAll()
}



