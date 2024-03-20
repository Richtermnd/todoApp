package server

import (
	"net/http"
)

// POST /todo/
func (s *Server) createTodo(w http.ResponseWriter, r *http.Request) {
	todo, err := decodeTodo(r)

	if err != nil {
		handleErr(err, w)
		return
	}

	todo, err = s.service.CreateTodo(r.Context(), todo)
	if err != nil {
		handleErr(err, w)
		return
	}

	if err := encodeTodo(w, todo); err != nil {
		handleErr(err, w)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// PUT /todo/:id
func (s *Server) updateTodo(w http.ResponseWriter, r *http.Request) {
	todo, err := decodeTodo(r)
	if err != nil {
		handleErr(err, w)
		return
	}

	id, err := idFromPathValue(r)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	todo.ID = id

	if err := s.service.UpdateTodo(r.Context(), todo); err != nil {
		handleErr(err, w)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// GET /todo/:id
func (s *Server) getTodo(w http.ResponseWriter, r *http.Request) {
	id, err := idFromPathValue(r)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	todo, err := s.service.Todo(r.Context(), id)
	if err != nil {
		handleErr(err, w)
		return
	}
	if err := encodeTodo(w, todo); err != nil {
		handleErr(err, w)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// DELETE /todo/:id
func (s *Server) deleteTodo(w http.ResponseWriter, r *http.Request) {
	id, err := idFromPathValue(r)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	if err := s.service.DeleteTodo(r.Context(), id); err != nil {
		handleErr(err, w)
		return
	}
	w.WriteHeader(http.StatusOK)
}
