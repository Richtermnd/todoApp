package service

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	"github.com/Richtermnd/todoApp/internal/domain"
)

var (
	ErrBadRequest = errors.New("bad request")
)

type Storage interface {
	Todo(ctx context.Context, id int) (domain.Todo, error)
	CreateTodo(ctx context.Context, todo domain.Todo) (domain.Todo, error)
	UpdateTodo(ctx context.Context, todo domain.Todo) error
	DeleteTodo(ctx context.Context, id int) error
}

type Service struct {
	log     *slog.Logger
	storage Storage
}

func New(log *slog.Logger, storage Storage) *Service {
	return &Service{
		log:     log,
		storage: storage,
	}
}

func (s *Service) Todo(ctx context.Context, id int) (domain.Todo, error) {
	const op = "service.Todo"

	log := s.log.With("op", op, "id", id)
	log.Info("get todo")

	if id < 0 {
		log.Info("invalid id")
		return domain.Todo{}, fmt.Errorf("%s: %w", op, ErrBadRequest)
	}

	todo, err := s.storage.Todo(ctx, id)
	if err != nil {
		log.Error("failed to get todo", "err", err)
		return domain.Todo{}, fmt.Errorf("%s: %w", op, err)
	}
	log.Info("got todo")
	return todo, nil
}

func (s *Service) CreateTodo(ctx context.Context, todo domain.Todo) (domain.Todo, error) {
	return s.storage.CreateTodo(ctx, todo)
}

func (s *Service) UpdateTodo(ctx context.Context, todo domain.Todo) error {
	return s.storage.UpdateTodo(ctx, todo)
}

func (s *Service) DeleteTodo(ctx context.Context, id int) error {
	return s.storage.DeleteTodo(ctx, id)
}
