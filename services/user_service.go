package services

import (
	"iris/commons"
	"iris/datamodels"
	"iris/repositories"
	"iris/web/viewmodels"
	"strconv"
)

// UserService handles CRUID operations of a user datamodel,
// it depends on a user repository for its actions.
// It's here to decouple the data source from the higher level compoments.
// As a result a different repository type can be used with the same logic without any aditional changes.
// It's an interface and it's used as interface everywhere
// because we may need to change or try an experimental different domain logic at the future.
type UserService interface {
	GetAll() (users []datamodels.Biz_user, total int)
	GetByID(id int64) (user datamodels.Biz_user, found bool)
	//GetByID(id int64) (datamodels.User, bool)
	GetByUsernameAndPassword(username, userPassword string) (datamodels.Biz_user, bool)
	//DeleteByID(id int64) bool
	//
	//Update(id int64, user datamodels.User) (datamodels.User, error)
	//UpdatePassword(id int64, newPassword string) (datamodels.User, error)
	//UpdateUsername(id int64, newUsername string) (datamodels.User, error)
	//
	CreateUser(user datamodels.Biz_user) (datamodels.Biz_user, bool)
	GetCharacterposBy(id int) viewmodels.ChPositionModel
}

type userService struct {
	repo repositories.UserRepository
}

// NewUserService returns the default user service.
func NewUserService() UserService {
	return &userService{
		repo: repositories.NewUserDBRep(),
	}
}

// GetAll returns all users.
func (s *userService) GetAll() (users []datamodels.Biz_user, total int) {
	return s.repo.SelectAll()
}

func (s *userService) GetByID(id int64) (user datamodels.Biz_user, found bool) {
	return s.repo.GetID(id)
}

func (s *userService) GetByUsernameAndPassword(username, userPassword string) (user datamodels.Biz_user, found bool) {
	return s.repo.GetByUsernameAndPassword(username, userPassword)
}

func (s *userService) CreateUser(user datamodels.Biz_user) (datamodels.Biz_user, bool) {
	return s.repo.CreateUser(user)
}

func (s *userService) GetCharacterposBy(chId int) (res viewmodels.ChPositionModel) {
	gcdbci := s.repo.GetChDataByChId(chId)
	res.MapId, _ = strconv.Atoi(gcdbci.MapId)
	res.PosX = gcdbci.PointX
	res.PosY = gcdbci.PointY
	if gcdbci.CurrentStatus == commons.CH_MOVING {
		at := s.repo.DoCheckMoveStatus(chId)
		if at == "" {
			gcdbci.CurrentStatus = commons.CH_FREE
			s.repo.UpdMoveStatus(chId, commons.CH_FREE)
		} else {
			res.ArriveTime = at
		}
	}
	return res
}
