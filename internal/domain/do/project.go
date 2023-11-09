package do

import "time"

type Project struct {
	Name        string    //名称
	BeginTime   time.Time //开始时间
	EndTime     time.Time //结束时间
	Description string    //描述
	Skills      []string  //技能，技术栈
	Works       []string  //职能，
}
