package employer

import (
	"context"
	"gorm.io/gorm/clause"
	"grapefruit/internal/adapter/psql"
	"grapefruit/internal/domain/po"
)

type Employer struct {
}

func (e *Employer) UpdateInfo(user *po.User) {
	psql.GetGormDB().WithContext(context.Background()).Updates(user)
}

func (e *Employer) AddResume(resume *po.Resume) {
	psql.GetGormDB().WithContext(context.Background()).Create(resume)
}

func (e *Employer) UpdateResume(resume *po.Resume) {
	psql.GetGormDB().WithContext(context.Background()).Clauses(clause.OnConflict{Columns: []clause.Column{{Name: "id"}}}).
		Updates(resume)

}

func (e *Employer) DeleteResume(resume *po.Resume) {
	psql.GetGormDB().WithContext(context.Background()).Where("id=?", resume.ID).Delete(resume)

}

func (e *Employer) ShareResume(id int64) {
	//todo
}
