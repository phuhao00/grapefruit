package config

const (
	Version = "1.0.1"

	Getwxacodeunlimit                = "https://api.weixin.qq.com/wxa/getwxacodeunlimit?access_token=%s" //wechat GetMiniProgram
	S3BucketUrl                      = ""
	FetchAccessTokenBaseUrl          = "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s"
	AppIdPublicAccount               = "" //  public account AppId
	AppSecretPublicAccount           = "" //  public account AppSecret
	AppIdMiniProgram                 = "" // 小程序 AppId
	AppSecretMiniProgram             = "" // 小程序 AppSecret
	MiniProgramFetchAccessTokenTimes = 10
)
