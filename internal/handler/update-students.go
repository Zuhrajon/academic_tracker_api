package handler

import (
	"academic-tracker-api/internal/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) UpdateStudents(c *gin.Context) {
	studentsIdStr := c.Param("id")
	if studentsIdStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "id отсутствует в пути",
		})
		return
	}

	studentsId, err := strconv.Atoi(studentsIdStr)
	if err != nil || studentsId <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "некорректный studentsId, должно быть положительное число",
		})
		return
	}

	var student model.Student
	err = c.ShouldBindJSON(&student)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "некорректное тело запроса",
		})
		return
	}

	updatedStudent, err := h.service.UpdateStudents(studentsId, student)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"student": updatedStudent,
	})
}
