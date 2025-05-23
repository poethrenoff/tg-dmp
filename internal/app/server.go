package app

import (
	"os"

	"github.com/gin-gonic/gin"
)

func RunServer(h *Handler) {
	if os.Getenv("APP_ENV") == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()
	router.GET("/get/:id", h.ProfileHandler.GetProfile)
	router.Run(":8080")
}
