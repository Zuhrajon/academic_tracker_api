package handler

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func getAuthUser(c *gin.Context) (int, string, bool) {
	userIDValue, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "user_id not found in token",
		})
		return 0, "", false
	}

	roleValue, exists := c.Get("role")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "role not found in token",
		})
		return 0, "", false
	}

	userID, ok := userIDValue.(int)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "invalid user_id in token",
		})
		return 0, "", false
	}

	role, ok := roleValue.(string)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "invalid role in token",
		})
		return 0, "", false
	}

	return userID, role, true
}

func writeAccessError(c *gin.Context, err error) {
	status := http.StatusBadRequest

	if err.Error() == "access denied" {
		status = http.StatusForbidden
	} else if strings.Contains(err.Error(), "not found") {
		status = http.StatusNotFound
	}

	c.JSON(status, gin.H{
		"error": err.Error(),
	})
}
