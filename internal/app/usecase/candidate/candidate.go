package candidate

import "grapefruit/internal/domain/po"

//求职者

type Candidate struct {
}

func (c *Candidate) PublishJob(job *po.Job) {
	//TODO implement me
	panic("implement me")
}

func (c *Candidate) CloseJob(jobId int64) {
	//TODO implement me
	panic("implement me")
}

func (c *Candidate) EditJob(job *po.Job) {
	//TODO implement me
	panic("implement me")
}
