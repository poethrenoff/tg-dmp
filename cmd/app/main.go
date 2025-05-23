package main

import (
	"os"
	"os/signal"
	"syscall"
	"tg-dmp/internal/app"
)

func main() {
	app.LoadEnvironment()

	database := app.InitDatabase()
	defer database.Close()

	repository := app.NewRepository(database)
	service := app.NewService(repository)
	handler := app.NewHandler(service)

	server := app.NewServer("8080", handler)
	defer server.Shutdown()

	go func() {
		server.Run()
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
}
