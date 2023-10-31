package service

import "grapefruit/internal/domain/po"

//候选者

type ICandidate interface {
	PublishJob(job *po.Job)
	CloseJob(jobId int64)
	EditJob(job *po.Job)
}
