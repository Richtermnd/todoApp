package application

import (
	"log/slog"

	"github.com/Richtermnd/todoApp/internal/server"
	"github.com/Richtermnd/todoApp/internal/service"
	mapstorage "github.com/Richtermnd/todoApp/internal/storage/map_storage"
)

type App struct {
	log *slog.Logger

	server *server.Server
}

func New(log *slog.Logger) *App {
	storage := mapstorage.New()
	service := service.New(log, storage)
	server := server.New(service)
	return &App{
		log:    log,
		server: server,
	}
}

func (a *App) Start() {
	go a.server.Start()
}

func (a *App) Shutdown() {
	a.server.Shutdown()
}
