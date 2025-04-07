package task

import "time"

type Task struct {
    ID          string
    Name        string
    Deadline    time.Time
    IsCompleted bool
}

func NewTask(id, name string, deadline time.Time) *Task {
    return &Task{
        ID:          id,
        Name:        name,
        Deadline:    deadline,
        IsCompleted: false,
    }
}