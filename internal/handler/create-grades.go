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

	actorUserID, actorRole, ok := getAuthUser(c)
	if !ok {
		return
	}

	createGrades, err := h.service.CreateGradesForUser(grades, actorUserID, actorRole)
	if err != nil {
		writeAccessError(c, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"grades": createGrades,
	})

}
