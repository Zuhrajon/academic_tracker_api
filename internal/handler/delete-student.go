package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

func (h *Handler) DeleteStudent(c *gin.Context) {
	studentIdStr := c.Param("id")
	if studentIdStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id",
		})
		return
	}

	studentId, err := strconv.Atoi(studentIdStr)
	if err != nil || studentId <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid studentID; must be a positive number.\n",
		})
		return
	}

	err = h.service.DeleteStudent(studentId)
	if err != nil {
		if strings.Contains(err.Error(), "violates foreign key constraint") {
			c.JSON(http.StatusConflict, gin.H{
				"error": "student cannot be deleted because it has grades or attendance records",
			})
			return
		}

		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "successfully deleted",
	})
}
