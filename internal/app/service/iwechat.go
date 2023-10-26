package service

import "grapefruit/internal/domain/do"

type IWechat interface {
	GetMiniProgramQRCode(req *do.ClientCreateQRCodeReq) (url string, err error)
	GetAccessToken() (string, error)
}
