package services

import (
	"iris/datamodels"
	"iris/repositories"
)

type TaskService interface {
	GetAllTaskList() []datamodels.BizTask
	ScrambleTask(userId int64, taskId int) (bool, error)
}

func NewTaskService(repo repositories.TaskRepository) TaskService {
	return &taskService{
		repo: repo,
	}
}

type taskService struct {
	repo repositories.TaskRepository
}

func (s *taskService) GetAllTaskList() []datamodels.BizTask {
	return s.repo.GetAllTaskList()
}

func (s *taskService) ScrambleTask(userId int64, taskId int) (bool, error) {
	return s.repo.ScrambleTask(userId, taskId)
}
