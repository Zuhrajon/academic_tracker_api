package handler

import (
	"academic-tracker-api/internal/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) Login(c *gin.Context) {
	var request model.LoginRequest

	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	response, err := h.service.Login(request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response)
}
