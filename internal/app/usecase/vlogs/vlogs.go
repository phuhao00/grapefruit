package vlogs

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"mime/multipart"
	"net/http"
)

func UpdateVlogs(ctx *gin.Context) {
	//todo 限制数量
	//todo update
	file, err := ctx.FormFile("video")
	if err != nil {
		// 处理错误
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 打开上传的视频文件
	src, err := file.Open()
	if err != nil {
		// 处理错误
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer src.Close()

	// 创建一个 buffer 用于存储文件内容
	buffer := &bytes.Buffer{}

	// 创建一个 multipart.Writer
	writer := multipart.NewWriter(buffer)

	// 创建一个表单字段，用于存储视频文件
	fileField, err := writer.CreateFormFile("video", "video.mp4")
	if err != nil {
		log.Fatal(err)
	}
	// 将视频文件内容复制到表单字段中
	_, err = io.Copy(fileField, src)
	if err != nil {
		log.Fatal(err)
	}

	// 关闭 multipart.Writer，以便写入结束标志
	err = writer.Close()
	if err != nil {
		log.Fatal(err)
	}

	// 创建 HTTP POST 请求
	req, err := http.NewRequest("POST", "http://example.com/upload", buffer)
	if err != nil {
		log.Fatal(err)
	}

	// 设置请求头
	req.Header.Set("Content-Type", writer.FormDataContentType())

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// 检查响应状态码
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("upload failed with status: %s", resp.Status)
	}

	// 处理响应结果
	fmt.Println("Upload successful!")
}

func UploadVlogs(ctx *gin.Context) {
	//todo 限制数量
	//todo 传到s3
	//todo update
}
