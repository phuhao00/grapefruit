package routing

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Register(engine *gin.Engine) {
	engine.GET("/healthy/check", func(context *gin.Context) {
		context.JSON(http.StatusOK, "I'm healthy")
	})

	//
}
