package config

import (
	"fmt"
	"grapefruit/kit/utils"
)

func MakeS3SaveQRCodeKey(userId int64, picTag string) string {
	return fmt.Sprintf("qrcode-%d"+picTag, userId)
}

var (
	FetchMiniProgramAccessTokenCompleteUrl   = fmt.Sprintf(FetchAccessTokenBaseUrl, AppIdMiniProgram, AppSecretMiniProgram)
	FetchPublicAccountAccessTokenCompleteUrl = fmt.Sprintf(FetchAccessTokenBaseUrl, AppIdPublicAccount, AppSecretPublicAccount)
)

var SystemName = "GrapeFruit"
var ServerAddress = "http://localhost:8080"

var SyncFrequency = utils.GetOrDefault("SYNC_FREQUENCY", 10*60) // unit is second

var UsingPostgreSQL = false
var RedisEnabled = true

var BatchUpdateEnabled = false
var BatchUpdateInterval = utils.GetOrDefault("BATCH_UPDATE_INTERVAL", 5)
var QuotaRemindThreshold = int32(1000)

var SMTPServer = ""
var SMTPPort = 587
var SMTPAccount = ""
var SMTPFrom = ""
var SMTPToken = ""
