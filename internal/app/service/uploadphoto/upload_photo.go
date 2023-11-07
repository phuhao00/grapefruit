package uploadphoto

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"grapefruit/internal/app/service"
	"mime/multipart"
)

var UploadPhotoService service.IUploadPhoto = &UploadPhoto{}

type UploadPhoto struct {
}

func (u *UploadPhoto) S3UpLoadPhoto(src multipart.File, fileName string) error {
	// 设置 AWS 配置
	awsConfig := &aws.Config{
		Region: aws.String("YOUR_AWS_REGION"), // 替换为你的 AWS 区域
		// 可选：如果需要使用 AWS 认证信息，例如访问私有 S3 存储桶，可以在这里配置
		// Credentials: credentials.NewStaticCredentials("YOUR_AWS_ACCESS_KEY", "YOUR_AWS_SECRET_KEY", ""),
	}

	// 创建 S3 会话
	sess, err := session.NewSession(awsConfig)
	if err != nil {
		return err
	}

	// 创建 S3 服务客户端
	svc := s3.New(sess)

	// 创建 S3 对象的输入参数
	params := &s3.PutObjectInput{
		Bucket: aws.String("YOUR_S3_BUCKET_NAME"), // 替换为你的 S3 存储桶名称
		Key:    aws.String(fileName),              // 使用上传文件的原始文件名作为 S3 对象的键
		Body:   src,
		ACL:    aws.String("public-read"), // 可选：设置 S3 对象的访问权限
	}

	// 将图片文件上传到 S3
	_, err = svc.PutObject(params)
	return err
}

func (u *UploadPhoto) OssUploadPhoto() {
	//TODO implement me
	panic("implement me")
}
