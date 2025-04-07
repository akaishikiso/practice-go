package main

import (
	"net/http"

	application "practice-go/backend/app"
	"practice-go/backend/infrastructure"
	"practice-go/backend/interfaces"
)

func main() {
    repo := infrastructure.NewInMemoryTaskRepository()
    service := application.NewTaskService(repo)
    handler := interfaces.NewTaskHandler(service)

    http.HandleFunc("/tasks", handler.GetTasks)

    http.ListenAndServe(":8080", nil)
}