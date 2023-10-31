package service

import "grapefruit/internal/domain/po"

type IUser interface {
	UpdateInfo(user *po.User)
}
