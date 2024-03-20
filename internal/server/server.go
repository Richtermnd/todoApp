package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Richtermnd/todoApp/internal/config"
	"github.com/Richtermnd/todoApp/internal/domain"
)

type Service interface {
	Todo(ctx context.Context, id int) (domain.Todo, error)
	CreateTodo(ctx context.Context, todo domain.Todo) (domain.Todo, error)
	UpdateTodo(ctx context.Context, todo domain.Todo) error
	DeleteTodo(ctx context.Context, id int) error
}

type Server struct {
	server  *http.Server
	service Service
	handler *http.ServeMux
}

func New(service Service) *Server {
	handler := http.NewServeMux()
	httpServer := &http.Server{
		Addr:    getAddr(),
		Handler: handler,
	}

	s := &Server{
		server:  httpServer,
		service: service,
		handler: handler,
	}
	s.register()
	return s
}

func (s *Server) Start() {
	if err := s.server.ListenAndServe(); err != nil {
		fmt.Println("Server stopped")
	}
}

func (s *Server) Shutdown() {
	s.server.Shutdown(context.Background())
}

func (s *Server) register() {

	s.handler.HandleFunc("GET /todo/{id}", checkValidId(s.getTodo))
	s.handler.HandleFunc("POST /todo", s.createTodo)
	s.handler.HandleFunc("PUT /todo/{id}", checkValidId(s.updateTodo))
	s.handler.HandleFunc("DELETE /todo/{id}", checkValidId(s.deleteTodo))
}

func getAddr() string {
	port := config.Config().Server.Port

	return fmt.Sprintf(":%d", port)
}
