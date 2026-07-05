package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) GetStudentById(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid student id",
		})
		return
	}

	if id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "student id must be greater than zero",
		})
		return
	}

	student, err := h.service.GetStudentById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "student not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"student": student,
	})
}
