package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) GetSubject(c *gin.Context) {
	subjects := h.service.GetSubject()

	c.JSON(http.StatusOK, gin.H{
		"subjects": subjects,
	})
}
