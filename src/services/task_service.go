package services

import (
	"fmt"
	"redis-task-queue/src/repositories"

	"github.com/google/uuid"
)

type TaskService struct {
	TaskRepo *repositories.TaskRepository
}

func NewTaskService(taskRepo *repositories.TaskRepository) *TaskService {
	return &TaskService{TaskRepo: taskRepo}
}

func (s *TaskService) CreateTask(queueName string) (string, error) {
	taskId := uuid.New().String()
	statusKey := fmt.Sprintf("task:%s:status", taskId)

	err := s.TaskRepo.AddTask(queueName, taskId, statusKey)
	if err != nil {
		return "", nil
	}

	return taskId, nil
}

func (s *TaskService) GetTaskStatus(taskId string) (string, error) {
	statusKey := fmt.Sprintf("task:%s:status", taskId)
	return s.TaskRepo.GetTaskStatus(statusKey)
}
