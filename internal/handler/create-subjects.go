package handler

import (
	"academic-tracker-api/internal/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) CreateSubjects(c *gin.Context) {
	var subject model.Subject

	err := c.ShouldBindJSON(&subject)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	createdSubjects, err := h.service.CreateSubjects(subject)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"subject": createdSubjects,
	})
}
