package photos

import (
	"github.com/gin-gonic/gin"
	"grapefruit/internal/app/service/uploadphoto"
	"net/http"
)

func UpdatePhotos(ctx *gin.Context) {
	//todo 限制数量
	//todo update
}

func UploadPhotos(c *gin.Context) {
	//todo 限制数量
	//todo 传到s3
	//todo update
	file, err := c.FormFile("image")
	if err != nil {
		// 处理错误
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 打开上传的图片文件
	src, err := file.Open()
	if err != nil {
		// 处理错误
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer src.Close()
	uploadphoto.UploadPhotoService.S3UpLoadPhoto(src)
	// 在这里可以对图片文件进行处理，例如保存到本地或进行其他操作

	// 返回成功的响应
	c.JSON(http.StatusOK, gin.H{"message": "图片上传成功"})
}
