package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) GetAttendanceByStudentID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id",
		})
		return
	}

	if id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "id must be greater than zero",
		})
		return
	}

	getAttendance, err := h.service.GetAttendanceByStudentID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "status not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"attendance": getAttendance,
	})

}
