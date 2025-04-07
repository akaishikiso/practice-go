package interfaces

import (
	"encoding/json"
	"net/http"
	application "practice-go/backend/app"
)

type TaskHandler struct {
    service *application.TaskService
}

func NewTaskHandler(service *application.TaskService) *TaskHandler {
    return &TaskHandler{service: service}
}

func (h *TaskHandler) GetTasks(w http.ResponseWriter, r *http.Request) {
    tasks, err := h.service.GetTasks()
    if err != nil {
        http.Error(w, "Failed to fetch tasks", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(tasks)
}