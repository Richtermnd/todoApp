package main

import (
	"log/slog"
	"os"
	"os/signal"

	"github.com/Richtermnd/todoApp/internal/application"
)

func main() {
	log := slog.Default()
	app := application.New(log)

	app.Start()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	// Не выполнится до тех пор пока сюда что-то не придёт.
	<-ch
	log.Info("Shutting down...")
	app.Shutdown()
}
