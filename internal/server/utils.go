package server

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Richtermnd/todoApp/internal/domain"
	"github.com/Richtermnd/todoApp/internal/service"
	"github.com/Richtermnd/todoApp/internal/storage"
)

func decodeTodo(r *http.Request) (domain.Todo, error) {
	var todo domain.Todo
	err := json.NewDecoder(r.Body).Decode(&todo)
	return todo, err
}

func encodeTodo(w http.ResponseWriter, todo domain.Todo) error {
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(todo)
}

func idFromPathValue(r *http.Request) (int, error) {
	return strconv.Atoi(r.PathValue("id"))
}

func handleErr(err error, w http.ResponseWriter) {
	switch err {
	case storage.ErrNotFound:
		http.Error(w, err.Error(), http.StatusNotFound)
	case service.ErrBadRequest:
		http.Error(w, err.Error(), http.StatusBadRequest)
	case storage.ErrAlreadyExist:
		http.Error(w, err.Error(), http.StatusConflict)
	default:
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
