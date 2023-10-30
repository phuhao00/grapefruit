package middleware

import (
	"github.com/gin-gonic/gin"
	"grapefruit/config"
	"grapefruit/kit/log"
	"grapefruit/kit/utils"
)

func abortWithMessage(c *gin.Context, statusCode int, message string) {
	c.JSON(statusCode, gin.H{
		"error": gin.H{
			"message": utils.MessageWithRequestId(message, c.GetString(config.RequestIdKey)),
			"type":    "one_api_error",
		},
	})
	c.Abort()
	log.Error(message)
}
