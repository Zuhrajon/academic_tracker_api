package handler

import (
	"academic-tracker-api/internal/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) CreateAttendance(c *gin.Context) {
	var attendance model.Attendance

	err := c.ShouldBindJSON(&attendance)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	createAttendance, err := h.service.CreateAttendance(attendance)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"attendance": createAttendance,
	})
}
