package view

import (
	"github.com/gin-gonic/gin"
	"grapefruit/internal/app/service/other"
)

//View 游客浏览
func View(ctx *gin.Context) {
	other.ViewService.View()
}
