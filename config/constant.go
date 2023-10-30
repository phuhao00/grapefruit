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

const (
	RequestIdKey = "X-Oneapi-Request-Id"
)

const (
	TokenStatusEnabled   = 1 // don't use 0, 0 is the default value!
	TokenStatusDisabled  = 2 // also don't use 0
	TokenStatusExpired   = 3
	TokenStatusExhausted = 4
)

const (
	UserStatusEnabled  = 1 // don't use 0, 0 is the default value!
	UserStatusDisabled = 2 // also don't use 0
)

const (
	RoleGuestUser  = 0
	RoleCommonUser = 1
	RoleAdminUser  = 10
	RoleRootUser   = 100
)
