package service

import "mime/multipart"

type IUploadPhoto interface {
	S3UpLoadPhoto(file multipart.File, fileName string) error
	OssUploadPhoto()
}
