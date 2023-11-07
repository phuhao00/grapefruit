package service

type IUploadVlog interface {
	S3UpLoadVlog()
	OssUploadVlog()
}
