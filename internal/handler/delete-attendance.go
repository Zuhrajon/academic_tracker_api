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
		return
	}

	attendanceId, err := strconv.Atoi(attendanceStr)
	if err != nil || attendanceId <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid attendanceId must be a positive number.",
		})
		return
	}

	actorUserID, actorRole, ok := getAuthUser(c)
	if !ok {
		return
	}

	err = h.service.DeleteAttendance(attendanceId, actorUserID, actorRole)
	if err != nil {
		writeAccessError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "successfully deleted",
	})

}
