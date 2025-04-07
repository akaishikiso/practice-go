package task

type TaskRepository interface {
    FindAll() ([]*Task, error)
}