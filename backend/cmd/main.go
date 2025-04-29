package main

import (
	"log"
	"net/http"

	"practice-go/backend/interface/handler"
	"practice-go/backend/interface/persistence"
	"practice-go/backend/usecase/menu"
)

func main() {
	repo := persistence.NewInMemoryMenuRepo()
	service := &menu.Service{Repo: repo}
	h := &handler.MenuHandler{Service: service}

	http.HandleFunc("/menus", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			h.GetAll(w, r)
		case "POST":
			h.Create(w, r)
		case "DELETE":
			h.Delete(w, r)
		default:
			http.Error(w, "method not allowed", 405)
		}
	})

	log.Println("Server running at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}