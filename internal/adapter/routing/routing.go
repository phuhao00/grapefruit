package routing

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "grapefruit/docs"
	"grapefruit/internal/adapter/routing/middleware"
	"grapefruit/internal/app/usecase/login"
	"grapefruit/internal/app/usecase/photos"
	"grapefruit/internal/app/usecase/vlogs"
	"net/http"
)

// @Title grapefruit
// @securityDefinitions.apiKey JWT
// @in header
// @name Authorization
func Register(engine *gin.Engine) {
	engine.GET("/api/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	engine.GET("/healthy/check", func(context *gin.Context) {
		context.JSON(http.StatusOK, "I'm healthy")
	})
	//
	engine.Use(middleware.WithCors())
	//
	engine.POST("api/login", login.Login)
	engine.POST("api/register", login.UserRegister)
	//
	engine.POST("api/upload_vlog", vlogs.UpdateVlogs)
	engine.POST("api/upload_photo", photos.UploadPhotos)
}
