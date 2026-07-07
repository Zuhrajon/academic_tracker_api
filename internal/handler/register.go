package handler

import (
	"academic-tracker-api/internal/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) Register(c *gin.Context) {
	var request model.RegisterRequest

	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	user, err := h.service.Register(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"user": user,
	})
}
