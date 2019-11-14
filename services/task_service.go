package services

import (
	"iris/datamodels"
	"iris/repositories"
)

type TaskService interface {
	GetAllTaskList() []datamodels.BizTask
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
