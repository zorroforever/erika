package services

import (
	"fmt"
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
	GetCharacterDataById(id int) viewmodels.ChDataModel
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
		atStr := commons.TimeCovert2Str(at)
		if atStr == "" {
			gcdbci.CurrentStatus = commons.CH_FREE
			s.repo.UpdMoveStatus(chId, commons.CH_FREE)
		} else {
			res.ArriveTime = atStr
		}
	}
	return res
}

func (s *userService) GetCharacterDataById(chId int) (res viewmodels.ChDataModel) {
	gcdbci := s.repo.GetChDataByChId(chId)
	gcpbbci := s.repo.GetChPropertyByChId(chId)
	gcmdbci, b := s.repo.GetChMoveDataByChId(chId)
	res.MapId, _ = strconv.Atoi(gcdbci.MapId)
	res.ChName = gcdbci.ChName
	res.Dex = gcpbbci.Dex
	res.UserId = gcdbci.UserId
	if b {
		res.CdStart = commons.GetNowStr()
		res.CdEnd = commons.TimeCovert2Str(gcmdbci.ArriveTime)
		fmt.Println("出来的时间是=" + res.CdEnd)
		res.CdNow = res.CdStart
		res.IsNeedCD = 1
		res.Sx = gcmdbci.SX
		res.Sy = gcmdbci.SY
	} else {
		res.IsNeedCD = 0
	}

	// 随意添加需要的角色相关字段
	return res
}
