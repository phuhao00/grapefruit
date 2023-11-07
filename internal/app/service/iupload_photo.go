package service

type IUploadPhoto interface {
	S3UpLoadPhoto()
	OssUploadPhoto()
}
