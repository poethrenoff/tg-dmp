package main

import (
	"os"
	"tg-dmp/internal/app"

	"github.com/gin-gonic/gin"
)

func main() {
	app.LoadEnvironment()

	database := app.InitDatabase()

	repository := app.NewProfileRepository(database)
	service := app.NewProfileService(repository)
	handler := app.NewProfileHandler(service)

	if os.Getenv("APP_ENV") == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()
	router.GET("/get/:id", handler.GetProfile)
	router.Run(":8080")
}
