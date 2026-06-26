package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) DeleteStudent(c *gin.Context) {
	studentIdStr := c.Param("id")
	if studentIdStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "id отсутствует в пути",
		})
		return
	}

	studentId, err := strconv.Atoi(studentIdStr)
	if err != nil || studentId <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "некорректный studentID, должно быть положительное число",
		})
		return
	}

	err = h.service.DeleteStudent(studentId)
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
