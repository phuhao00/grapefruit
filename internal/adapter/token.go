package adapter

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"grapefruit/config"
	"grapefruit/internal/adapter/psql"
	"grapefruit/internal/domain/po"
	"grapefruit/kit/log"
	"grapefruit/kit/utils"
)

func GetAllUserTokens(userId int, startIdx int, num int) ([]*po.Token, error) {
	var tokens []*po.Token
	var err error
	err = psql.GetGormDB().Where("user_id = ?", userId).Order("id desc").Limit(num).Offset(startIdx).Find(&tokens).Error
	return tokens, err
}

func SearchUserTokens(userId int, keyword string) (tokens []*po.Token, err error) {
	err = psql.GetGormDB().Where("user_id = ?", userId).Where("name LIKE ?", keyword+"%").Find(&tokens).Error
	return tokens, err
}

func ValidateUserToken(key string) (token *po.Token, err error) {
	if key == "" {
		return nil, errors.New("未提供令牌")
	}
	token, err = CacheGetTokenByKey(key)
	if err == nil {
		if token.Status == config.TokenStatusExhausted {
			return nil, errors.New("该令牌额度已用尽")
		} else if token.Status == config.TokenStatusExpired {
			return nil, errors.New("该令牌已过期")
		}
		if token.Status != config.TokenStatusEnabled {
			return nil, errors.New("该令牌状态不可用")
		}
		if token.ExpiredTime != -1 && token.ExpiredTime < utils.GetTimestamp() {
			if !config.RedisEnabled {
				token.Status = config.TokenStatusExpired
				err := SelectUpdate(token)
				if err != nil {
					log.Error("failed to update token status" + err.Error())
				}
			}
			return nil, errors.New("该令牌已过期")
		}
		if !token.UnlimitedQuota && token.RemianQuota <= 0 {
			if !config.RedisEnabled {
				// in this case, we can make sure the token is exhausted
				token.Status = config.TokenStatusExhausted
				err := SelectUpdate(token)
				if err != nil {
					log.Error("failed to update token status" + err.Error())
				}
			}
			return nil, errors.New("该令牌额度已用尽")
		}
		return token, nil
	}
	return nil, errors.New("无效的令牌")
}

func GetTokenByIds(id int32, userId int32) (*po.Token, error) {
	if id == 0 || userId == 0 {
		return nil, errors.New("id 或 userId 为空！")
	}
	token := po.Token{ID: id, UserID: userId}
	var err error = nil
	err = psql.GetGormDB().First(&token, "id = ? and user_id = ?", id, userId).Error
	return &token, err
}

func GetTokenById(id int32) (*po.Token, error) {
	if id == 0 {
		return nil, errors.New("id 为空！")
	}
	token := po.Token{ID: id}
	var err error = nil
	err = psql.GetGormDB().First(&token, "id = ?", id).Error
	return &token, err
}

func Insert(token *po.Token) error {
	var err error
	err = psql.GetGormDB().Create(token).Error
	return err
}

// Update Make sure your token's fields is completed, because this will update non-zero values
func Update(token *po.Token) error {
	var err error
	err = psql.GetGormDB().Model(token).Select("name", "status", "expired_time", "remain_quota", "unlimited_quota").Updates(token).Error
	return err
}

func SelectUpdate(token *po.Token) error {
	// This can update zero values
	return psql.GetGormDB().Model(token).Select("accessed_time", "status").Updates(token).Error
}

func Delete(token *po.Token) error {
	var err error
	err = psql.GetGormDB().Delete(token).Error
	return err
}

func DeleteTokenById(id int32, userId int32) (err error) {
	// Why we need userId here? In case user want to delete other's token.
	if id == 0 || userId == 0 {
		return errors.New("id 或 userId 为空！")
	}
	token := po.Token{ID: id, UserID: userId}
	err = psql.GetGormDB().Where(token).First(&token).Error
	if err != nil {
		return err
	}
	return Delete(&token)
}

func IncreaseTokenQuota(id int32, quota int32) (err error) {
	if quota < 0 {
		return errors.New("quota 不能为负数！")
	}
	if config.BatchUpdateEnabled {
		addNewRecord(BatchUpdateTypeTokenQuota, int(id), int(quota))
		return nil
	}
	return increaseTokenQuota(int(id), int(quota))
}

func increaseTokenQuota(id int, quota int) (err error) {
	err = psql.GetGormDB().Model(&po.Token{}).Where("id = ?", id).Updates(
		map[string]interface{}{
			"remain_quota":  gorm.Expr("remain_quota + ?", quota),
			"used_quota":    gorm.Expr("used_quota - ?", quota),
			"accessed_time": utils.GetTimestamp(),
		},
	).Error
	return err
}

func DecreaseTokenQuota(id int32, quota int32) (err error) {
	if quota < 0 {
		return errors.New("quota 不能为负数！")
	}
	if config.BatchUpdateEnabled {
		addNewRecord(BatchUpdateTypeTokenQuota, int(id), int(-quota))
		return nil
	}
	return decreaseTokenQuota(id, quota)
}

func decreaseTokenQuota(id int32, quota int32) (err error) {
	err = psql.GetGormDB().Model(&po.Token{}).Where("id = ?", id).Updates(
		map[string]interface{}{
			"remain_quota":  gorm.Expr("remain_quota - ?", quota),
			"used_quota":    gorm.Expr("used_quota + ?", quota),
			"accessed_time": utils.GetTimestamp(),
		},
	).Error
	return err
}

func PreConsumeTokenQuota(tokenId int32, quota int32) (err error) {
	if quota < 0 {
		return errors.New("quota 不能为负数！")
	}
	token, err := GetTokenById(tokenId)
	if err != nil {
		return err
	}
	if !token.UnlimitedQuota && token.RemianQuota < quota {
		return errors.New("令牌额度不足")
	}
	userQuota, err := GetUserQuota(token.UserID)
	if err != nil {
		return err
	}
	if userQuota < quota {
		return errors.New("用户额度不足")
	}
	quotaTooLow := userQuota >= config.QuotaRemindThreshold && userQuota-quota < config.QuotaRemindThreshold
	noMoreQuota := userQuota-quota <= 0
	if quotaTooLow || noMoreQuota {
		go func() {
			email, err := GetUserEmail(token.UserID)
			if err != nil {
				log.Error("failed to fetch user email: " + err.Error())
			}
			prompt := "您的额度即将用尽"
			if noMoreQuota {
				prompt = "您的额度已用尽"
			}
			if email != "" {
				topUpLink := fmt.Sprintf("%s/topup", config.ServerAddress)
				err = SendEmail(prompt, email,
					fmt.Sprintf("%s，当前剩余额度为 %d，为了不影响您的使用，请及时充值。<br/>充值链接：<a href='%s'>%s</a>", prompt, userQuota, topUpLink, topUpLink))
				if err != nil {
					log.Error("failed to send email" + err.Error())
				}
			}
		}()
	}
	if !token.UnlimitedQuota {
		err = DecreaseTokenQuota(tokenId, quota)
		if err != nil {
			return err
		}
	}
	err = DecreaseUserQuota(token.UserID, quota)
	return err
}

func PostConsumeTokenQuota(tokenId int32, quota int32) (err error) {
	token, err := GetTokenById(tokenId)
	if quota > 0 {
		err = DecreaseUserQuota(token.UserID, quota)
	} else {
		err = IncreaseUserQuota(token.UserID, -quota)
	}
	if err != nil {
		return err
	}
	if !token.UnlimitedQuota {
		if quota > 0 {
			err = DecreaseTokenQuota(tokenId, quota)
		} else {
			err = IncreaseTokenQuota(tokenId, -quota)
		}
		if err != nil {
			return err
		}
	}
	return nil
}
