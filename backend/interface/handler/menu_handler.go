package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	usecase "practice-go/backend/usecase/menu"
)

type MenuHandler struct {
	Service *usecase.Service
}

func (h *MenuHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	menus, err := h.Service.GetAll()
	if err != nil {
		http.Error(w, "error", 500)
		return
	}
	json.NewEncoder(w).Encode(menus)
}

func (h *MenuHandler) Create(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Name string `json:"name"`
		Date string `json:"date"`
		Note string `json:"note"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "bad request", 400)
		return
	}
	menu, err := h.Service.Create(input.Name, input.Date, input.Note)
	if err != nil {
		http.Error(w, "fail to create", 500)
		return
	}
	json.NewEncoder(w).Encode(menu)
}

func (h *MenuHandler) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(idStr)
	if err := h.Service.Delete(id); err != nil {
		http.Error(w, "not found", 404)
		return
	}
	w.WriteHeader(http.StatusOK)
}