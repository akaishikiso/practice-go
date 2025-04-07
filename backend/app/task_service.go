package application

import "practice-go/backend/domain/task"

type TaskService struct {
    repo task.TaskRepository
}

func NewTaskService(repo task.TaskRepository) *TaskService {
    return &TaskService{repo: repo}
}

func (s *TaskService) GetTasks() ([]*task.Task, error) {
    return s.repo.FindAll()
}