package config

import "fmt"

func MakeS3SaveQRCodeKey(userId int64, picTag string) string {
	return fmt.Sprintf("qrcode-%d"+picTag, userId)
}

var (
	FetchMiniProgramAccessTokenCompleteUrl   = fmt.Sprintf(FetchAccessTokenBaseUrl, AppIdMiniProgram, AppSecretMiniProgram)
	FetchPublicAccountAccessTokenCompleteUrl = fmt.Sprintf(FetchAccessTokenBaseUrl, AppIdPublicAccount, AppSecretPublicAccount)
)
