package app

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

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

func NewProfileHandler(s ProfileService) *profileHandler {
	return &profileHandler{service: s}
}
