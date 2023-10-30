package adapter

import (
	"encoding/json"
	"fmt"
	"grapefruit/config"
	"grapefruit/internal/adapter/psql"
	"grapefruit/internal/adapter/redis"
	"grapefruit/internal/domain/po"
	"grapefruit/kit/log"
	"strconv"
	"sync"
	"time"
)

var (
	TokenCacheSeconds         = config.SyncFrequency
	UserId2GroupCacheSeconds  = config.SyncFrequency
	UserId2QuotaCacheSeconds  = config.SyncFrequency
	UserId2StatusCacheSeconds = config.SyncFrequency
)

func CacheGetTokenByKey(key string) (*po.Token, error) {
	keyCol := "`key`"
	if config.UsingPostgreSQL {
		keyCol = `"key"`
	}
	var token po.Token
	if !config.RedisEnabled {
		err := psql.GetGormDB().Where(keyCol+" = ?", key).First(&token).Error
		return &token, err
	}
	tokenObjectString, err := redis.RedisGet(fmt.Sprintf("token:%s", key))
	if err != nil {
		err := psql.GetGormDB().Where(keyCol+" = ?", key).First(&token).Error
		if err != nil {
			return nil, err
		}
		jsonBytes, err := json.Marshal(token)
		if err != nil {
			return nil, err
		}
		err = redis.RedisSet(fmt.Sprintf("token:%s", key), string(jsonBytes), time.Duration(TokenCacheSeconds)*time.Second)
		if err != nil {
			log.Error("Redis set token error: " + err.Error())
		}
		return &token, nil
	}
	err = json.Unmarshal([]byte(tokenObjectString), &token)
	return &token, err
}

func CacheGetUserGroup(id int) (group string, err error) {
	if !config.RedisEnabled {
		return GetUserGroup(id)
	}
	group, err = redis.RedisGet(fmt.Sprintf("user_group:%d", id))
	if err != nil {
		group, err = GetUserGroup(id)
		if err != nil {
			return "", err
		}
		err = redis.RedisSet(fmt.Sprintf("user_group:%d", id), group, time.Duration(UserId2GroupCacheSeconds)*time.Second)
		if err != nil {
			log.Error("Redis set user group error: " + err.Error())
		}
	}
	return group, err
}

func CacheGetUserQuota(id int32) (quota int32, err error) {
	if !config.RedisEnabled {
		return GetUserQuota(id)
	}
	quotaString, err := redis.RedisGet(fmt.Sprintf("user_quota:%d", id))
	if err != nil {
		quota, err = GetUserQuota(id)
		if err != nil {
			return 0, err
		}
		err = redis.RedisSet(fmt.Sprintf("user_quota:%d", id), fmt.Sprintf("%d", quota), time.Duration(UserId2QuotaCacheSeconds)*time.Second)
		if err != nil {
			log.Error("Redis set user quota error: " + err.Error())
		}
		return quota, err
	}
	quotaTmp, err := strconv.Atoi(quotaString)
	return int32(quotaTmp), err
}

func CacheUpdateUserQuota(id int32) error {
	if !config.RedisEnabled {
		return nil
	}
	quota, err := GetUserQuota(id)
	if err != nil {
		return err
	}
	err = redis.RedisSet(fmt.Sprintf("user_quota:%d", id), fmt.Sprintf("%d", quota), time.Duration(UserId2QuotaCacheSeconds)*time.Second)
	return err
}

func CacheDecreaseUserQuota(id int, quota int) error {
	if !config.RedisEnabled {
		return nil
	}
	err := redis.RedisDecrease(fmt.Sprintf("user_quota:%d", id), int64(quota))
	return err
}

func CacheIsUserEnabled(userId int) (bool, error) {
	if !config.RedisEnabled {
		return IsUserEnabled(userId)
	}
	enabled, err := redis.RedisGet(fmt.Sprintf("user_enabled:%d", userId))
	if err == nil {
		return enabled == "1", nil
	}

	userEnabled, err := IsUserEnabled(userId)
	if err != nil {
		return false, err
	}
	enabled = "0"
	if userEnabled {
		enabled = "1"
	}
	err = redis.RedisSet(fmt.Sprintf("user_enabled:%d", userId), enabled, time.Duration(UserId2StatusCacheSeconds)*time.Second)
	if err != nil {
		log.Error("Redis set user enabled error: " + err.Error())
	}
	return userEnabled, err
}

var channelSyncLock sync.RWMutex
