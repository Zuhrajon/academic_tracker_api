package handler

import (
	"academic-tracker-api/internal/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) CreateGrades(c *gin.Context) {
	var grades model.Grade

	err := c.ShouldBindJSON(&grades)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "bad request",
		})
		return
	}

	createGrades, err := h.service.CreateGrades(grades)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"grades": createGrades,
	})

}
