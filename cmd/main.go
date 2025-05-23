package main

import (
	"tg-dmp/internal/app"
)

func main() {
	app.LoadEnvironment()

	database := app.InitDatabase()

	repository := app.NewRepository(database)
	service := app.NewService(repository)
	handler := app.NewHandler(service)

	app.RunServer(handler)
}
