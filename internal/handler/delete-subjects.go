package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

func (h *Handler) DeleteSubjects(c *gin.Context) {
	subjectIdStr := c.Param("id")
	if subjectIdStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "id отсутствует в пути",
		})
		return
	}

	subjectId, err := strconv.Atoi(subjectIdStr)
	if err != nil || subjectId <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "некорректный subjectId, должно быть положительное число",
		})
		return
	}

	err = h.service.DeleteSubjects(subjectId)
	if err != nil {
		if strings.Contains(err.Error(), "violates foreign key constraint") {
			c.JSON(http.StatusConflict, gin.H{
				"error": "subject cannot be deleted because it has grades or attendance records",
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
