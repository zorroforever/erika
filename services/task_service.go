package services

import (
	"iris/commons"
	"iris/datamodels"
	"iris/repositories"
)

type TaskService interface {
	GetAllTaskList() commons.Page
	ScrambleTask(userId int64, taskId int) (bool, error)
	GetTaskListByMapId(mapId int) (list []datamodels.BizTask)
	DoPersonGetTask(userId int, chId int, taskId int) (bool, error)
}

type taskService struct {
	repo repositories.TaskRepository
}

func NewTaskService() TaskService {
	return &taskService{
		repo: repositories.NewTaskDBRep(),
	}
}

func (s *taskService) GetAllTaskList() commons.Page {
	return s.repo.GetAllTaskList()
}

func (s *taskService) ScrambleTask(userId int64, taskId int) (bool, error) {
	return s.repo.ScrambleTask(userId, taskId)
}

func (s *taskService) GetTaskListByMapId(mapId int) (list []datamodels.BizTask) {
	cnt, list := s.repo.GetTaskListByMapId(mapId)
	if cnt > 0 {
		return list
	} else {
		return nil
	}
}

func (s *taskService) DoPersonGetTask(userId int, chId int, taskId int) (bool, error) {
	return s.repo.DoPersonGetTask(userId, chId, taskId)
}
