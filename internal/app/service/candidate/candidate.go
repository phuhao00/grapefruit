package candidate

import (
	"context"
	"gorm.io/gorm/clause"
	"grapefruit/internal/adapter/psql"
	"grapefruit/internal/domain/po"
)

//求职者

type Candidate struct {
}

func (c *Candidate) UpdateInfo(user *po.User) {
	psql.GetGormDB().WithContext(context.Background()).Updates(user)
}

func (c *Candidate) PublishJob(job *po.Job) {
	psql.GetGormDB().WithContext(context.Background()).Create(job)
}

func (c *Candidate) CloseJob(jobId int64) {
	psql.GetGormDB().WithContext(context.Background()).Where("id=?", jobId).Set("status", 2)
}

func (c *Candidate) EditJob(job *po.Job) {
	psql.GetGormDB().WithContext(context.Background()).Clauses(clause.OnConflict{Columns: []clause.Column{{Name: "id"}}}).Updates(job)

}
