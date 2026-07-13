package handler

import (
	"academic-tracker-api/internal/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) UpdateGrade(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id",
		})
		return
	}

	if id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "id must not be less or equal to zero",
		})
		return
	}

	var grade model.Grade
	err = c.ShouldBindJSON(&grade)
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

	updateGrade, err := h.service.UpdateGrade(id, grade, actorUserID, actorRole)
	if err != nil {
		writeAccessError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"grade": updateGrade,
	})

}
