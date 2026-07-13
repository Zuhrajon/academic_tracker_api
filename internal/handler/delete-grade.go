package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) DeleteGrade(c *gin.Context) {
	idStr := c.Param("id")
	if idStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id",
		})
		return
	}

	gradeId, err := strconv.Atoi(idStr)
	if err != nil || gradeId <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid gradeId; must be a positive number.\n",
		})
		return
	}

	actorUserID, actorRole, ok := getAuthUser(c)
	if !ok {
		return
	}

	err = h.service.DeleteGrade(gradeId, actorUserID, actorRole)
	if err != nil {
		writeAccessError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "successfully deleted",
	})

}
