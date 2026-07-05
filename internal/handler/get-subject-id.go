package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) GetSubjectById(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid subject id",
		})
		return
	}

	if id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "subject id must be greater than zero",
		})
		return
	}

	subject, err := h.service.GetSubjectById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"subject": subject,
	})

}
