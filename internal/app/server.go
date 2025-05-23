package app

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run() {
	if err := s.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("error occured on running server: %s", err.Error())
	}
}

func (s *Server) Shutdown() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := s.httpServer.Shutdown(ctx); err != nil {
		log.Fatalf("error occured on server shutting down: %s", err.Error())
	}
}

func NewServer(port string, handler *Handler) *Server {
	router := gin.Default()
	router.GET("/get/:id", handler.ProfileHandler.GetProfile)

	server := &Server{
		httpServer: &http.Server{
			Addr:         ":" + port,
			Handler:      router.Handler(),
			ReadTimeout:  10 * time.Second,
			WriteTimeout: 10 * time.Second,
		},
	}

	return server
}
