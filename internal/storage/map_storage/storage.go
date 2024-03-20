package mapstorage

import (
	"context"

	"github.com/Richtermnd/todoApp/internal/domain"
	"github.com/Richtermnd/todoApp/internal/storage"
)

type Storage struct {
	i int
	m map[int]domain.Todo
}

func New() *Storage {
	return &Storage{
		m: make(map[int]domain.Todo),
	}
}

func (s *Storage) Todo(ctx context.Context, id int) (domain.Todo, error) {
	todo, ok := s.m[id]
	if !ok {
		return domain.Todo{}, storage.ErrNotFound
	}
	return todo, nil
}

func (s *Storage) CreateTodo(ctx context.Context, todo domain.Todo) (domain.Todo, error) {
	todo.ID = s.i
	s.m[s.i] = todo
	s.i++
	return todo, nil
}

func (s *Storage) UpdateTodo(ctx context.Context, todo domain.Todo) error {
	s.m[todo.ID] = todo
	return nil
}

func (s *Storage) DeleteTodo(ctx context.Context, id int) error {
	delete(s.m, id)
	return nil
}
