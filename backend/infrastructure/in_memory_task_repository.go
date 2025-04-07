package infrastructure

import (
	"practice-go/backend/domain/task"
	"time"
)

type InMemoryTaskRepository struct {
    tasks []*task.Task
}

func NewInMemoryTaskRepository() *InMemoryTaskRepository {
    return &InMemoryTaskRepository{
        tasks: []*task.Task{
            task.NewTask("1", "Task 1", time.Date(2025, 4, 10, 0, 0, 0, 0, time.UTC)),
            task.NewTask("2", "Task 2", time.Date(2025, 4, 15, 0, 0, 0, 0, time.UTC)),
        },
    }
}

func (r *InMemoryTaskRepository) FindAll() ([]*task.Task, error) {
    return r.tasks, nil
}