package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) GetStudents(c *gin.Context) {
	students := h.service.GetStudents()

	c.JSON(http.StatusOK, gin.H{
		"students": students,
	})
}
