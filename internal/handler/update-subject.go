package handler

import (
	"academic-tracker-api/internal/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) UpdateSubject(c *gin.Context) {
	subjectsIdStr := c.Param("id")
	if subjectsIdStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "id отсутствует в пути",
		})
		return
	}

	subjectId, err := strconv.Atoi(subjectsIdStr)
	if err != nil || subjectId <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "некорректный subjectId, должно быть положительное число",
		})
		return
	}

	var subject model.Subject
	err = c.ShouldBindJSON(&subject)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "некорректное тело запроса",
		})
		return
	}

	updatedSubject, err := h.service.UpdateSubject(subjectId, subject)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"subject": updatedSubject,
	})
}
