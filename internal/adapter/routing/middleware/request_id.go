package middleware

import (
	"context"
	"github.com/gin-gonic/gin"
	"grapefruit/config"
	"grapefruit/kit/utils"
)

func RequestId() func(c *gin.Context) {
	return func(c *gin.Context) {
		id := utils.GetTimeString() + utils.GetRandomString(8)
		c.Set(config.RequestIdKey, id)
		ctx := context.WithValue(c.Request.Context(), config.RequestIdKey, id)
		c.Request = c.Request.WithContext(ctx)
		c.Header(config.RequestIdKey, id)
		c.Next()
	}
}
