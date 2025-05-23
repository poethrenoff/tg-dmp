package app

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProfileHandler interface {
	GetProfile(c *gin.Context)
}

type profileHandler struct {
	service ProfileService
}

func (h *profileHandler) GetProfile(c *gin.Context) {
	id := c.Params.ByName("id")
	profile, err := h.service.GetProfileById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorMessage{fmt.Sprintf("User with id \"%s\" not found", id)})
	} else {
		c.JSON(http.StatusOK, profile)
	}
}

type Handler struct {
	ProfileHandler
}

func (h *Handler) GerRouter() *gin.Engine {
	router := gin.New()
	router.GET("/get/:id", h.ProfileHandler.GetProfile)
	return router
}

func NewHandler(s *Service) *Handler {
	return &Handler{
		ProfileHandler: &profileHandler{service: s.ProfileService},
	}
}
