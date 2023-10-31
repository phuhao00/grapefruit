package service

import "grapefruit/internal/domain/po"

//IEmployer 雇主

type IEmployer interface {
	AddResume(resume *po.Resume)
	UpdateResume(resume *po.Resume)
	DeleteResume(resume *po.Resume)
	ShareResume(id int64)
}
