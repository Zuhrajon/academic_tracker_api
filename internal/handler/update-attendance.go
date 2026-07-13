package handler

import (
	"academic-tracker-api/internal/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) UpdateAttendance(c *gin.Context) {
	attendanceIdStr := c.Param("id")
	if attendanceIdStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id",
		})
		return
	}

	attendanceId, err := strconv.Atoi(attendanceIdStr)
	if err != nil || attendanceId <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid attendanceId; must be a positive number.",
		})
		return
	}

	var attendance model.Attendance
	err = c.ShouldBindJSON(&attendance)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	actorUserID, actorRole, ok := getAuthUser(c)
	if !ok {
		return
	}

	updateAttendance, err := h.service.UpdateAttendance(attendanceId, attendance, actorUserID, actorRole)
	if err != nil {
		writeAccessError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"attendance": updateAttendance,
	})

}
