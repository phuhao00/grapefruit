package service

import "grapefruit/internal/domain/po"

type IUser interface {
	//UpdateBaseInfo 更新基础信息
	UpdateBaseInfo(user *po.User)
	//UpdateVlogs 更新vlog
	UpdateVlogs(user *po.User)
	//UpdatePhotos 更新照片墙
	UpdatePhotos(user *po.User)
	//UpdateProject 更新项目
	UpdateProject(user *po.User)
	//ShareIntroduction 分享介绍
	ShareIntroduction(user *po.User)
}
