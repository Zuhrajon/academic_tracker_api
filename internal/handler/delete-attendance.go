package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) DeleteAttendance(c *gin.Context) {
	attendanceStr := c.Param("id")
	if attendanceStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id",
		})
	}

	attendanceId, err := strconv.Atoi(attendanceStr)
	if err != nil || attendanceId <= 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Invalid studentID; must be a positive number.\n",
		})
		return
	}

	err = h.service.DeleteAttendance(attendanceId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "successfully deleted",
	})

}
