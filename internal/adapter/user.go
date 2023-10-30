package adapter

import (
	"errors"
	"gorm.io/gorm"
	"grapefruit/config"
	"grapefruit/internal/adapter/psql"
	"grapefruit/internal/domain/po"
	"grapefruit/kit/log"
	"strings"
)

func GetUserQuota(id int32) (quota int32, err error) {
	err = psql.GetGormDB().Model(&po.User{}).Where("id = ?", id).Select("quota").Find(&quota).Error
	return quota, err
}

func GetUserGroup(id int) (group string, err error) {
	groupCol := "`group`"
	if config.UsingPostgreSQL {
		groupCol = `"group"`
	}

	err = psql.GetGormDB().Model(&po.User{}).Where("id = ?", id).Select(groupCol).Find(&group).Error
	return group, err
}

func GetUserEmail(id int32) (email string, err error) {
	err = psql.GetGormDB().Model(&po.User{}).Where("id = ?", id).Select("email").Find(&email).Error
	return email, err
}

func IsUserEnabled(userId int) (bool, error) {
	if userId == 0 {
		return false, errors.New("user id is empty")
	}
	var user po.User
	err := psql.GetGormDB().Where("id = ?", userId).Select("status").Find(&user).Error
	if err != nil {
		return false, err
	}
	return user.Status == config.UserStatusEnabled, nil
}

func IncreaseUserQuota(id int32, quota int32) (err error) {
	if quota < 0 {
		return errors.New("quota 不能为负数！")
	}
	if config.BatchUpdateEnabled {
		addNewRecord(BatchUpdateTypeUserQuota, int(id), int(quota))
		return nil
	}
	return increaseUserQuota(id, quota)
}

func increaseUserQuota(id int32, quota int32) (err error) {
	err = psql.GetGormDB().Model(&po.User{}).Where("id = ?", id).Update("quota", gorm.Expr("quota + ?", quota)).Error
	return err
}

func DecreaseUserQuota(id int32, quota int32) (err error) {
	if quota < 0 {
		return errors.New("quota 不能为负数！")
	}
	if config.BatchUpdateEnabled {
		addNewRecord(BatchUpdateTypeUserQuota, int(id), int(-quota))
		return nil
	}
	return decreaseUserQuota(id, quota)
}

func updateUserUsedQuota(id int, quota int) {
	err := psql.GetGormDB().Model(&po.User{}).Where("id = ?", id).Updates(
		map[string]interface{}{
			"used_quota": gorm.Expr("used_quota + ?", quota),
		},
	).Error
	if err != nil {
		log.Error("failed to update user used quota: " + err.Error())
	}
}

func updateUserRequestCount(id int, count int) {
	err := psql.GetGormDB().Model(&po.User{}).Where("id = ?", id).Update("request_count", gorm.Expr("request_count + ?", count)).Error
	if err != nil {
		log.Error("failed to update user request count: " + err.Error())
	}
}

func GetUsernameById(id int) (username string) {
	psql.GetGormDB().Model(&po.User{}).Where("id = ?", id).Select("username").Find(&username)
	return username
}

func ValidateAccessToken(token string) (user *po.User) {
	if token == "" {
		return nil
	}
	token = strings.Replace(token, "Bearer ", "", 1)
	user = &po.User{}
	if psql.GetGormDB().Where("access_token = ?", token).First(user).RowsAffected == 1 {
		return user
	}
	return nil
}

func IsAdmin(userId int32) bool {
	if userId == 0 {
		return false
	}
	var user po.User
	err := psql.GetGormDB().Where("id = ?", userId).Select("role").Find(&user).Error
	if err != nil {
		log.Error("no such user " + err.Error())
		return false
	}
	return user.Role >= config.RoleAdminUser
}
